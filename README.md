# Backend for teamchallange project

## Used FastAPI

- To run install all packages from requirements.txt:
    `pip install requirements.txt` -- used pip 22.0.2, python 3.10
- From root:
    `uvicorn app.main:app` -- by default localhost:8000
- Hosted at [render](https://hello-backend-7125.onrender.com/docs#/).
- Docker image [cr1m3s/fastapi](https://hub.docker.com/repository/registry-1.docker.io/cr1m3s/fastapi/general) (outdated).
- Requires DATABASE_URL set as sys env to deploy.
- Docs can befound at [{URL}/docs](https://hello-backend-7125.onrender.com/docs)

## For DB used PostgreSQL

- For migration use:
    1. `alembic revision --autogenerate -m "New Migration"`
    2. `alembic upgrade head`
