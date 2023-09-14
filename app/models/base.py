import os
import sqlalchemy
from sqlalchemy_utils import create_database, database_exists
from dotenv import load_dotenv

load_dotenv()
engine = sqlalchemy.create_engine(os.environ['DATABASE_URL'])


def validate_database():
    if not database_exists(engine.url):
        create_database(engine.url)
