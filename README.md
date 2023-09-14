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

## Локальний запуск

Для того аби запустити проект локально потрібно спочатку встановити собі docker та docker-compose. Актуальні інструкції по встановленню можна знайти по наведеним посиланням, для кожної операційної системи як то linux, windows, mac.

- https://docs.docker.com/engine/install/
- https://docs.docker.com/compose/install/

### Запуск docker-compose

Після встановлення вище наведених залежностей, потрібно викачати собі на локальну машину репозиторій з кодом і перейти в гілку dev.

Cтворити docker мережу tch_network за допомогою команди:

```bash
docker network create tch_network
```

Перевірити чи була створена docker мережа tch_network за допомогою команди:

```bash
docker network ls
```

Створити постійний docker volume за допомогою команди:

```bash
docker volume create tch_postgres_pg_data
```

Перевірити чи був створений docker volume з назвою tch_postgres_pg_data за допомогою команди:

```bash
docker volume ls
```

Використовуючи консоль зайти в директорію dev_local яка містить файл docker-compose.yml і скориставшись наведеною комадною запустити контейнери для локальної розробки фронтенд застосунку.

```bash
docker-compose up -d
```

Дочекатися поки будуть завантажені усі необхідні образи з інтернету і перевірити стан роботи за допомогою наведеної команди в середині директорії dev_local

```bash
docker-compose ps
```

Якщо контейнери запущені тоді потрібно зайти в середину виконавши таку команду:

```bash
docker exec -it tch_backend bash
```

Перебуваючи в середині контейнера запустити наступні команди:

```bash
pip install --no-cache-dir --upgrade -r ./requirements.txt
apt-get update
apt-get install postgresql-client
```

Запустити базу данних ввівши пароль `postgres`, перевірити вміст:
```bash
psql "postgresql://postgres:postgres@tch_postgres:5432/store"
Password for user postgres:
...
postgres-# \dt
              List of relations
 Schema |      Name       | Type  |  Owner
--------+-----------------+-------+----------
 public | alembic_version | table | postgres
 public | author          | table | postgres
 public | comments        | table | postgres
postgres-# \q
```

Перебуваючи в середині контейнера провести міграцію бази даних виконавши наступні команди:

```bash
cd app
alembic upgrade head -- для початкової міграції, якщо бд пуста
alembic revision --autogenerate -m "New Migration"   -- для подальших
alembic upgrade head                                 --
```
- за замовчування посилання на базу данних визначено в `app/main.py` як:

Для регістрації і входу використовується firebase, скопіюйьте tch_firefbase_account_keys.json в app/.

```python3
db_url="postgresql://postgres:postgres@tch_postgres:5432/store"
```

Для запуску застосунку, виконати команду в середині контейнера:

```bash
uvicorn "app.main:app" --host "0.0.0.0" --port "8000"
```

В випадку успішного запуску ви зможете мати доступ до застосунку через ваш браузре за адресою http://localhost:8000
