from datetime import timedelta
from typing import Union

from dotenv import load_dotenv
from fastapi import FastAPI, Depends, HTTPException
from fastapi.security import OAuth2PasswordRequestForm
from sqlalchemy.orm import Session

from .models import User
from . import schemas
from .auth import create_access_token, get_current_user, hash_password, verify_password
from .db import Base, engine, get_db
from .repositories.user_repo import UserRepositoryDB

load_dotenv()
app = FastAPI()

Base.metadata.create_all(bind=engine)


def issue_token(username: str):
    token = create_access_token(
        data={"sub": username},
        expires_delta=timedelta(minutes=30)
    )
    return {"access_token": token, "token_type": "bearer"}


@app.post("/analyze")
async def read_root():
    return {}


@app.get("/status/{task_id}")
async def read_item(task_id: int, q: Union[str, None] = None):
    return {"item_id": task_id, "q": q}


@app.get("/report/{username}")
async def read_item(username: str):
    return {"username": username}


@app.post("/register")
async def register(user: schemas.UserCreate, db: Session = Depends(get_db)):
    repo = UserRepositoryDB(db)
    if repo.get_by_username(user.username):
        raise HTTPException(status_code=400, detail="Username already exists")
    hashed_password = hash_password(user.password)
    user_with_hash = user.copy(update={"password": hashed_password})
    new_user = repo.create(user_with_hash)
    return issue_token(new_user.username)


@app.post("/token")
async def login(form_data: OAuth2PasswordRequestForm = Depends(), db: Session = Depends(get_db)):
    repo = UserRepositoryDB(db)
    user = repo.get_by_username(form_data.username)
    if not user or not verify_password(form_data.password, user.password):
        raise HTTPException(status_code=401, detail="Incorrect username or password")

    return issue_token(username=user.username)


@app.get("/protected")
async def read_protected(current_user: User = Depends(get_current_user)):
    return {"message": f"Hello, {current_user.username}"}
