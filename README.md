# Backend for teamchallange project

## Used FastAPI

- To run install all packages from requirements.txt:
    `pip install requirements.txt`
- From root:
    `uvicorn app.main:app` -- by default localhost:8000
- Hosted at [render](https://hello-world-sije.onrender.com/docs#/).
- Docker image [cr1m3s/fastapi](https://hub.docker.com/repository/registry-1.docker.io/cr1m3s/fastapi/general).
- Requires DATABASE_URL set as sys env to deploy.
- Docs can befound at [{URL}/docs](https://hello-world-sije.onrender.com/docs#/)

## For DB used PostgreSQL

- For migration use:
    1. `alembic revision --autogenerate -m "New Migration"`
    2. `alembic upgrade head`
