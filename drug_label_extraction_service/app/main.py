from fastapi import FastAPI, UploadFile
from starlette.middleware.cors import CORSMiddleware
import os
import cv2
import numpy as np
from starlette.responses import JSONResponse
from services.image_preprocessing.service import preprocess
from services.ocr.service import ocr
from services.extract_data.service import extract_data

app = FastAPI()

origins = [
    "http://localhost:3000",
]

app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

usage_config = '--oem 3 --psm 6 -l tha'
drug_name_config = '--oem 3 --psm 6 -l tha+eng'


@app.post("/drug_label/extract/")
async def extract_drugs(file: UploadFile):
    try:
        # with open(file.filename, "wb") as f:
        #     f.write(file.file.read())
        # image = cv2.imread(file.filename)
        content = await file.read()
        nparr = np.frombuffer(content, np.uint8)
        image = cv2.imdecode(nparr, cv2.IMREAD_COLOR)

        app_dir = os.path.dirname(os.path.realpath(__file__))
        template_image_path = os.path.join(app_dir, 'images', 'logo.jpg')
        template_image = cv2.imread(template_image_path)
        #
        usage_image, drug_name_image = preprocess(image, template_image)
        usage_text = ocr(usage_image, usage_config)
        drug_name_text = ocr(drug_name_image, drug_name_config)
        res = extract_data(drug_name_text, usage_text)

        return JSONResponse(status_code=200, content=res)

    except Exception as e:
        return JSONResponse(status_code=500, content={"error": str(e)})
