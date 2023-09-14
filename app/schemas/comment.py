from pydantic import BaseModel


class Comment(BaseModel):
    title: str
    rating: int
    author_id: int

    class Config:
        orm_mode = True
