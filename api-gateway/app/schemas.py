from pydantic import BaseModel, EmailStr


class UserCreate(BaseModel):
    username: str
    email: EmailStr
    password: str


class AnalyzeRequest(BaseModel):
    username: str


class AnalyzeResponse(BaseModel):
    task_id: str
    status: str
    preview: dict
