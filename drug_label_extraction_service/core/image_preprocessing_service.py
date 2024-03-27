import cv2
import numpy as np


def _apply_gray(img):
    return cv2.cvtColor(img, cv2.COLOR_RGB2GRAY)

def _apply_filter(img):
    gray = _apply_gray(img)
    kernel = np.ones((5, 5), np.float32) / 15
    filtered = cv2.filter2D(gray, -1, kernel)
    return filtered

def _apply_threshold(filtered):
    thresh = cv2.threshold(filtered, 250, 255, cv2.THRESH_OTSU)[1]
    return thresh

def _detect_contour(img, img_shape):
    canvas = np.zeros(img_shape, np.uint8)
    contours, hierarchy = cv2.findContours(img, cv2.RETR_TREE, cv2.CHAIN_APPROX_NONE)
    cnt = sorted(contours, key=cv2.contourArea, reverse=True)[0]
    cv2.drawContours(canvas, cnt, -1, (0, 255, 255), 3)
    return canvas, cnt

def _detect_corners(canvas, cnt):
    epsilon = 0.02 * cv2.arcLength(cnt, True)
    approx_corners = cv2.approxPolyDP(cnt, epsilon, True)
    cv2.drawContours(canvas, approx_corners, -1, (255, 255, 0), 10)
    approx_corners = sorted(np.concatenate(approx_corners).tolist())
    approx_corners = [approx_corners[i] for i in [0, 2, 1, 3]]
    approx_corners = np.array(approx_corners)
    approx_corners = approx_corners.reshape((4, 2))
    new_approx_corners = np.zeros((4, 1, 2), dtype=np.int32)
    add = approx_corners.sum(1)

    new_approx_corners[0] = approx_corners[np.argmin(add)]
    new_approx_corners[3] = approx_corners[np.argmax(add)]
    diff = np.diff(approx_corners, axis=1)
    new_approx_corners[1] = approx_corners[np.argmin(diff)]
    new_approx_corners[2] = approx_corners[np.argmax(diff)]

    result_list = [sublist[0].tolist() if sublist.size > 0 else [] for sublist in new_approx_corners]
    return result_list

def _get_destination_points(corners):
    w1 = np.sqrt((corners[0][0] - corners[1][0]) ** 2 + (corners[0][1] - corners[1][1]) ** 2)
    w2 = np.sqrt((corners[2][0] - corners[3][0]) ** 2 + (corners[2][1] - corners[3][1]) ** 2)
    w = max(int(w1), int(w2))

    h1 = np.sqrt((corners[0][0] - corners[2][0]) ** 2 + (corners[0][1] - corners[2][1]) ** 2)
    h2 = np.sqrt((corners[1][0] - corners[3][0]) ** 2 + (corners[1][1] - corners[3][1]) ** 2)
    h = max(int(h1), int(h2))

    destination_corners = np.float32([(0, 0), (w - 1, 0), (0, h - 1), (w - 1, h - 1)])

    return destination_corners, h, w

def _unwarp(img, src, dst):
    h, w = img.shape[:2]
    H, _ = cv2.findHomography(src, dst, method=cv2.RANSAC, ransacReprojThreshold=3.0)
    un_warped = cv2.warpPerspective(img, H, (w, h), flags=cv2.INTER_LINEAR)
    return un_warped

def _resize_image(image, target_size=(2500, 2000)):
    return cv2.resize(image, target_size)

def _match_template(image, template):
    result = cv2.matchTemplate(image, template, cv2.TM_CCOEFF)

    min_val, max_val, min_loc, max_loc = cv2.minMaxLoc(result)

    h, w = template.shape
    top_left = max_loc
    bottom_right = (top_left[0] + w, top_left[1] + h)
    cv2.rectangle(image, top_left, bottom_right, (0, 255, 0), 2)
    return top_left[0], top_left[1]

def _crop_image(image, coord):
    x, y, w, h = coord[0], coord[1], coord[2], coord[3]
    return image[y:y + h, x:x + w]

def _rotate_image(image, h, w, y):
    if h > w:
        if y > 1000:
            return cv2.rotate(image, cv2.ROTATE_90_CLOCKWISE)
        else:
            return cv2.rotate(image, cv2.ROTATE_90_COUNTERCLOCKWISE)
    else:
        if y > 1000:
            return cv2.rotate(image, cv2.ROTATE_180)
    return image

def preprocess(image, template):
    filtered_image = _apply_filter(image)
    threshold_image = _apply_threshold(filtered_image)

    cnv, largest_contour = _detect_contour(threshold_image, image.shape)
    corners = _detect_corners(cnv, largest_contour)

    destination_points, h, w = _get_destination_points(corners)
    un_warped = _unwarp(image, np.float32(corners), destination_points)
    cropped = un_warped[0:h, 0:w]
    gray = _apply_gray(cropped)
    thresh = _apply_threshold(gray)

    gray_template_image = _apply_gray(template)
    thresh_template_image = _apply_threshold(gray_template_image)
    x, y = _match_template(thresh, thresh_template_image)
    rotated = _rotate_image(thresh, h, w, y)
    resize = _resize_image(rotated)
    x, y = _match_template(resize, thresh_template_image)

    usage_text_image = _crop_image(resize, (x - 100, y + 500, 2200, 350))
    drug_name_text_image = _crop_image(resize, (x - 100, y + 1300, 2200, 200))

    return usage_text_image, drug_name_text_image