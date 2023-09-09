import uvicorn
import os
from fastapi import FastAPI, HTTPException
from fastapi_sqlalchemy import DBSessionMiddleware, db 

from app.schema import Comment as SchemaComment
from app.schema import Author as SchemaAuthor

from app.schema import Comment
from app.schema import Author

from app.models import Comment as ModelComment
from app.models import Author as ModelAuthor

from fastapi.middleware.cors import CORSMiddleware

app = FastAPI()

#origins = [
#    "http://localhost:3000",
#    "http://localhost:3080",
#    "https://react-x1x9.onrender.com/",
#]

app.add_middleware(
    CORSMiddleware,
    allow_origins=['*'],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# to avoid csrftokenError
app.add_middleware(DBSessionMiddleware, db_url=os.environ['DATABASE_URL'])


@app.get("/")
async def root():
    return {"message": "hello world"}


@app.post('/comment/', response_model=SchemaComment)
async def comment(comment: SchemaComment):
    db_comment = ModelComment(title=comment.title, rating=comment.rating, author_id = comment.author_id)
    db.session.add(db_comment)
    db.session.commit()
    return db_comment


@app.get('/comment/')
async def comment():
    comment = db.session.query(ModelComment).all()
    return comment

@app.get('/comment/{id}')
async def comment(id: int):
    comment = db.session.query(ModelComment).get({"id": id})
    return comment


@app.delete('/comment/{id}')
async def comment(id: int):
    comment = db.session.query(ModelComment).filter(ModelComment.id == id).first()
    result = {"status": "not found"}
    if comment:
        db.session.delete(comment)
        db.session.commit()
        result["status"] = "deleted"
    return result

 
@app.post('/author/', response_model=SchemaAuthor)
async def author(author: SchemaAuthor):
    db_author = ModelAuthor(name=author.name, age=author.age)
    db.session.add(db_author)
    db.session.commit()
    return db_author

@app.get('/author/')
async def author():
    author = db.session.query(ModelAuthor).all()
    return author
