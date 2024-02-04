import cv2
import numpy as np


def grayscale_conversion(image):
    return cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)


def thresholding_conversion(gray_image):
    return cv2.threshold(gray_image, 130, 255, cv2.THRESH_BINARY)[1]


def find_edges(thresh_image):
    t_lower = 100
    t_upper = 200
    l2_gradient = True
    return cv2.Canny(thresh_image, t_lower, t_upper, L2gradient=l2_gradient)


def find_angle(edges_image):
    lines = cv2.HoughLines(edges_image, 1, np.pi / 180, threshold=200)

    angles = []
    for line in lines:
        rho, theta = line[0]
        angle = np.degrees(theta)
        angles.append(angle)

    return np.median(angles)


def rotate_image(image, angle):
    center = tuple(np.array(image.shape[1::-1]) / 2)
    rotation_matrix = cv2.getRotationMatrix2D(center, angle - 90, 1.0)
    return cv2.warpAffine(image, rotation_matrix, image.shape[1::-1], flags=cv2.INTER_LINEAR)


def crop_drug_label(rotated_image):
    gray = grayscale_conversion(rotated_image)
    thresh = thresholding_conversion(gray)
    contours, _ = cv2.findContours(thresh, cv2.RETR_EXTERNAL, cv2.CHAIN_APPROX_SIMPLE)
    max_contour = max(contours, key=cv2.contourArea)
    mask = np.zeros_like(gray)
    cv2.drawContours(mask, [max_contour], -1, (255,), thickness=cv2.FILLED)
    # display(mask)
    res = cv2.bitwise_and(thresh, thresh, mask=mask)
    x, y, w, h = cv2.boundingRect(max_contour)
    return res[y:y + h, x:x + w]


def resize_image(image, target_size=(2500, 2000)):
    return cv2.resize(image, target_size)


def match_template(image, template):
    result = cv2.matchTemplate(image, template, cv2.TM_CCOEFF)

    min_val, max_val, min_loc, max_loc = cv2.minMaxLoc(result)

    h, w = template.shape
    top_left = max_loc
    bottom_right = (top_left[0] + w, top_left[1] + h)
    cv2.rectangle(image, top_left, bottom_right, (0, 255, 0), 2)
    return top_left[0], top_left[1]


def crop_image(image, coord):
    x, y, w, h = coord[0], coord[1], coord[2], coord[3]
    return image[y:y + h, x:x + w]


def preprocess(image, template):
    gray_image = grayscale_conversion(image)
    thresh_image = thresholding_conversion(gray_image)
    edges = find_edges(thresh_image)
    angle = find_angle(edges)
    rotated = rotate_image(image, angle)
    drug_label = crop_drug_label(rotated)
    resize = resize_image(drug_label)

    gray_template_image = grayscale_conversion(template)
    thresh_template_image = thresholding_conversion(gray_template_image)
    x, y = match_template(resize, thresh_template_image)

    usage_text = crop_image(resize, (x - 100, y + 500, 2200, 350))
    drug_name_text = crop_image(resize, (x - 100, y + 1300, 2200, 200))

    return usage_text, drug_name_text
