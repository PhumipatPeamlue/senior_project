import os
import cv2
import numpy as np
from fastapi import FastAPI, UploadFile, HTTPException, Form
from fastapi.middleware.cors import CORSMiddleware
from typing import Union, Annotated

from core.image_preprocessing_service import preprocess
from core.ocr_service import ocr
from core.extract_data_service import extract_data


usage_config = '--oem 3 --psm 6 -l tha'
drug_name_config = '--oem 3 --psm 6 -l tha+eng'

app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=["http://localhost:3000"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)


@app.get("/")
async def root():
    return {"message": "Hello World"}


@app.get("/hello/{name}")
async def say_hello(name: str):
    return {"message": f"Hello {name}"}


@app.post("/drug_label_extraction")
async def extract_drug_label(user_id: Annotated[str, Form()], file: Union[UploadFile, None] = None):
    if not file:
        raise HTTPException(status_code=404, detail="image file not found")

    content = await file.read()
    nparr = np.frombuffer(content, np.uint8)
    image = cv2.imdecode(nparr, cv2.IMREAD_COLOR)

    app_dir = os.path.dirname(os.path.realpath(__file__))
    template_image_path = os.path.join(app_dir, 'images', 'logo.jpg')
    template_image = cv2.imread(template_image_path)

    usage_image, drug_name_image = preprocess(image, template_image)
    usage_text = ocr(usage_image, usage_config)
    drug_name_text = ocr(drug_name_image, drug_name_config)
    res = extract_data(drug_name_text, usage_text, user_id)

    return res