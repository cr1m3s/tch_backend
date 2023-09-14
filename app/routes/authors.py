from fastapi import APIRouter

from app.schemas.author import Author as SchemaAuthor
from app.models.models import Author as ModelAuthor

from fastapi_sqlalchemy import db

router = APIRouter()


@router.post('/author/', response_model=SchemaAuthor)
async def author(author: SchemaAuthor):
    db_author = ModelAuthor(name=author.name, age=author.age)
    db.session.add(db_author)
    db.session.commit()
    return db_author


@router.get('/author/')
async def author():
    author = db.session.query(ModelAuthor).all()
    return author
