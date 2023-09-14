import requests
import json


def login(email: str, password: str):
    body = {
       "email": email,
       "password": password
    }
    response = requests.post(url="http://0.0.0.0:8000/login", json=body)
    return json.loads(response.text)["token"]

token = login("abcd@abcd.com", "password")


def ping(token: str):
    headers = {
       'authorization': token
    }
    response = requests.post(url="http://0.0.0.0:8000/ping", headers=headers)
    return(response.text)


print(ping(token))
