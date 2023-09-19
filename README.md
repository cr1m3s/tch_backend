# Backend for teamchallange project

## Used FastAPI

- To run install all packages from requirements.txt:
`go mod download`
- From root:
  `go run app/main.go` -- by default localhost:8000
- Hosted at [render](https://hello-backend-7125.onrender.com/).
- Requires DATABASE_URL set as sys env to deploy.
- Docs can befound at [{URL}/docs/index.html](https://hello-backend-7125.onrender.com/swagger/index.html)
- To updated swagger after changing controllers run from the repo root:
`swag init -g app/main.go --output docs/ginsimple`

# For migrations from docker install migrate:
```
curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh |  bash
apt-get update
apt-get install migrate
```
- run migrations:
```
migrate -path db/migrations -database "postgresql://postgres:postgres@tch_postgres:5432/golang_postgres?sslmode=disable" -verbose up
```
- migration using files from:
```
/app/db/migrations/000001_init_schema.{up/down}.sql -- up used for upstream migration
                                                    -- down for downstream migration
```

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
go mod download
apt-get update
apt-get install postgresql-client
```

Запустити базу данних ввівши пароль `postgres`, перевірити вміст:
```bash
psql "postgresql://postgres:postgres@tch_postgres:5432/golang_postgres"
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
[TODO]
```
- за замовчування посилання на базу данних визначено в 
``` 
db_url="postgresql://postgres:postgres@tch_postgres:5432/store"
```

Для запуску застосунку, виконати команду в середині контейнера:

```bash
go run app/main.go
```

В випадку успішного запуску ви зможете мати доступ до застосунку через ваш браузре за адресою http://localhost:8000

Документація знаходиться за адресою http://localhost:8000/swagger/index.html
