from fastapi import FastAPI
from app.routes.test_router import test_router

app = FastAPI()

app.include_router(test_router)
