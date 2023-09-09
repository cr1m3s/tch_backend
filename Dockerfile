FROM python:3.11
WORKDIR /code
COPY ./requirements.txt /code/requirements.txt
RUN pip install --no-cache-dir --upgrade -r /code/requirements.txt
COPY ./app /code/app
ENV DATABASE_URL="postgresql://hello_postgre_user:N6ZYqyzSzfagCu29r9p4HcgcVkopVW2k@dpg-cjto551gdgss738nkak0-a.frankfurt-postgres.render.com/hello_postgre"
CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "8000"]

