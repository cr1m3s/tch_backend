import os
import uvicorn
import firebase_admin
import pyrebase
import json
from dotenv import load_dotenv
from firebase_admin import credentials, auth

from fastapi import FastAPI, Request
from fastapi.middleware.cors import CORSMiddleware
from fastapi.responses import JSONResponse
from fastapi.exceptions import HTTPException

from fastapi_sqlalchemy import DBSessionMiddleware

from app.routes.authors import router as authors_router
from app.routes.comments import router as comments_router
from app.routes.auth import router as auth_router

load_dotenv()

fire_config = {
  "apiKey": "AIzaSyAAU01dFLJrAhruwRZ95-n3_rfBZ9A-KtQ",
  "authDomain": "tch-auth.firebaseapp.com",
  "projectId": "tch-auth",
  "storageBucket": "tch-auth.appspot.com",
  "messagingSenderId": "298275933015",
  "appId": "1:298275933015:web:199b1cc8193c430456cc12",
  "databaseURL":"postgresql://postgres:postgres@tch_postgres:5432/store"
}


keys = os.path.join(os.path.dirname(os.path.abspath(__file__)), './', 'tch_firefbase_account_keys.json')
config = os.path.join(os.path.dirname(os.path.abspath(__file__)), './', 'firebase_config.json')
cred = credentials.Certificate(keys)
firebase = firebase_admin.initialize_app(cred)
pb = pyrebase.initialize_app(fire_config)
app = FastAPI()

allow_all = ['*']

app.add_middleware(
    CORSMiddleware,
    allow_origins=allow_all,
    allow_credentials=True,
    allow_methods=allow_all,
    allow_headers=allow_all,
)

# to avoid csrftokenError
app.add_middleware(DBSessionMiddleware, db_url=os.environ['DATABASE_URL'])


@app.get("/")
async def root():
    return {"message": "Hello, world!"}


@app.post("/ping")
async def validate(request: Request):
    headers = request.headers
    jwt = headers.get('authorization')
    print(f"jwt:{jwt}")
    user = auth.verify_id_token(jwt)
    return user["uid"]


app.include_router(authors_router)
app.include_router(comments_router)
app.include_router(auth_router)
