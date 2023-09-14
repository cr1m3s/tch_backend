FROM python:3.11.5-bullseye
WORKDIR /var/www/app

RUN pip install --no-cache-dir --upgrade -r ./requirements.txt
RUN apt-get update
RUN apt-get install postgresql-client
