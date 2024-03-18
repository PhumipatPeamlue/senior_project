FROM python:3.9
LABEL authors="besto"

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

ENV USER_SERVICE_HOST="localhost"
ENV USER_SERVICE_PORT="8080"

CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]