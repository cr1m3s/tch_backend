from fastapi import APIRouter

from app.schemas.comment import Comment as SchemaComment
from app.models.models import Comment as ModelComment

from fastapi_sqlalchemy import db

router = APIRouter()


@router.post('/comment/', response_model=SchemaComment)
async def comment(comment: SchemaComment):
    db_comment = ModelComment(title=comment.title, rating=comment.rating, author_id=comment.author_id)
    db.session.add(db_comment)
    db.session.commit()
    return db_comment


@router.get('/comment/')
async def comment():
    comment = db.session.query(ModelComment).all()
    return comment

@router.get('/comment/{id}')
async def comment(id: int):
    comment = db.session.query(ModelComment).get({"id": id})
    return comment


@router.delete('/comment/{id}')
async def comment(id: int):
    comment = db.session.query(ModelComment).filter(ModelComment.id == id).first()
    result = {"status": "not found"}
    if comment:
        db.session.delete(comment)
        db.session.commit()
        result["status"] = "deleted"
    return result
