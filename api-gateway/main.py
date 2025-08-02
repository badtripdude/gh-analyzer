from datetime import timedelta
from typing import Union

from fastapi import FastAPI, Depends, HTTPException
from fastapi.security import OAuth2PasswordRequestForm

from auth import authenticate_user, create_access_token, get_current_user

app = FastAPI()


@app.post("/analyze")
async def read_root():
    return {}


@app.get("/status/{task_id}")
async def read_item(task_id: int, q: Union[str, None] = None):
    print(task_id)
    return {"item_id": task_id, "q": q}


@app.get("/report/{username}")
async def read_item(username: str):
    return {"username": username}


@app.post("/token")
async def login(form_data: OAuth2PasswordRequestForm = Depends()):
    user = authenticate_user(form_data.username, form_data.password)
    if not user:
        raise HTTPException(status_code=401, detail="Incorrect username or password")

    access_token = create_access_token(
        data={"sub": user["username"]},
        expires_delta=timedelta(minutes=30)
    )
    return {"access_token": access_token, "token_type": "bearer"}


@app.get("/protected")
async def read_protected(current_user: dict = Depends(get_current_user)):
    return {"message": f"Hello, {current_user['username']}"}
