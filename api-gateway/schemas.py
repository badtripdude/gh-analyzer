from pydantic import EmailStr, BaseModel


class UserCreate(BaseModel):
    username: str
    email: EmailStr
    password: str
