from typing import Optional

from sqlalchemy.orm import Session

from ..models import User
from ..schemas import UserCreate
from .base import UserRepository


class UserRepositoryDB(UserRepository):
    def __init__(self, db: Session):
        self.db = db

    def get_by_username(self, username: str) -> Optional[User]:
        return self.db.query(User).filter(User.username == username).first()

    def get_by_id(self, user_id: int):
        return self.db.query(User).filter(User.id == user_id).first()

    def get_all(self):
        return self.db.query(User).all()

    def create(self, user_data: UserCreate) -> User:
        new_user = User(**user_data.dict())
        self.db.add(new_user)
        self.db.commit()
        self.db.refresh(new_user)
        return new_user

    def delete(self, user_id: int):
        user = self.get_by_id(user_id)
        if user:
            self.db.delete(user)
            self.db.commit()
        return user
