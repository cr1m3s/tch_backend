from fastapi import APIRouter, Request
from firebase_admin import auth
from fastapi.responses import JSONResponse
import pyrebase
import json
import os

path_to_json = os.path.join(os.path.dirname(os.path.abspath(__file__)), '..', 'firebase_config.json')
router = APIRouter()
pb = pyrebase.initialize_app(json.load(open(path_to_json)))


@router.post("/signup")
async def signup(request: Request):
    req = await request.json()
    email = req['email']
    password = req['password']
    if email is None or password is None:
        return HTTPException(detail={'message': 'Error! Missing Email or Password'}, status_code=400)
    try:
        user = auth.create_user(
           email=email,
           password=password
       )
        return JSONResponse(content={'message': f'Successfully created user {user.uid}'}, status_code=200)    
    except:
        return HTTPException(detail={'message': 'Error Creating User'}, status_code=400)


@router.post("/login")
async def login(request: Request):
    req_json = await request.json()
    email = req_json['email']
    password = req_json['password']
    try:
        user = pb.auth().sign_in_with_email_and_password(email, password)
        jwt = user['idToken']
        return JSONResponse(content={'token': jwt}, status_code=200)
    except:
        return HTTPException(detail={'message': 'There was an error logging in'}, status_code=400)
