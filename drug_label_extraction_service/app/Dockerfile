FROM python:3.9

RUN apt-get update && apt-get install -y \
    libgl1-mesa-glx \
    libsm6 libxext6 libxrender1 \
    tesseract-ocr \
    tesseract-ocr-eng \
    tesseract-ocr-tha

WORKDIR /app

COPY . .

RUN pip install -r requirements.txt

EXPOSE 8000

CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]