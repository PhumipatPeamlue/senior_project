import pytesseract as tess


def ocr(image, config):
    return tess.image_to_string(image, config=config)
