# Backend for teamchallange project

## Used Golang Gin Requires Go version 1.21+
- Before build project requires presence of .env file (example can be found in .env.example).
- To run install all packages from go.mod:
`go mod download`
- From root:
  `go run main.go` -- by default localhost:8000
- Hosted at [render](https://hello-backend-7125.onrender.com/).
- Requires DATABASE_URL set as sys env to deploy.

## Documentation
- Docs can befound at [{URL}/docs/index.html](https://hello-backend-7125.onrender.com/swagger/index.html)
- Swagger requires env variable 'SERVER_HOSTNAME' defined in .env file.
- To updated swagger after changing controllers run from the repo root:
1. Install swaggo:
  `go install github.com/swaggo/swag/cmd/swag@latest`
2. Generate docs:
  `swag init --parseDependency --parseInternal --parseDepth 1 -md ./documentation -o ./docs`
!! path to docs used in main.go: `_ "github.com/cr1m3s/tch_backend/app/docs"` !!

## DB communication
- To update router interfaces for queries:
  1. Install with:
    `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`
  2. Create rules in  `./query/{model_name}.sql` 
  3. In root folder with present `sqlc.yaml` file run: 
    `sqlc generate`
  4. Results will be in `./queries/`

## For migrations from docker install migrate:
```
curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh |  bash
apt-get update
apt-get install migrate
```
- run migrations:
```
migrate -path db/migrations -database "postgresql://postgres:postgres@tch_postgres:5432/store?sslmode=disable" -verbose up
```
- migration using files from:
```
/app/db/migrations/000001_init_schema.{up/down}.sql -- up used for upstream migration
                                                    -- down for downstream migration
```
- in case of failure and 'dirty database' connect to db with psql and run:
```
  update schema_migrations set dirty=false;
```

## Database examination

1. From project folder:
    `docker exec -it tch_postgres bash`
2. `psql -U postgres golang_postgres`
3. `\dt`
4.  If table 'users' not present run migrations, else:
    `SELECT * FROM users;`

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

- за замовчування посилання на базу данних визначено в .env:
`
DB_HOST=tch_postgres
DB_PORT=5432
DB_USER=postgres
DB_NAME=golang_postgres
DB_PASSWORD=postgres
DB_DRIVER=postgres
`

Для запуску застосунку, виконати команду в середині контейнера:

```bash
go run main.go
```

В випадку успішного запуску ви зможете мати доступ до застосунку через ваш браузр за адресою http://localhost:8000

Для тесту роботи серверу і підключення бд можна виконати запит наведений в req.txt

Документація знаходиться за адресою http://localhost:8000/api/docs/index.html

# Request examples
- User creation request example:
`curl -v  POST http://localhost:8000/api/auth/register -H "Content-type: application/json" -d '{"email": "hello@world", "name":"world", "password":"hello"}'`
Response body:
```
{
  "data": {
    "user": {
      "id": "143a5c7a-8e4b-4081-9da8-8005ee7b2e6f",
      "name": "world",
      "email": "hello@world",
      "photo": "default.jpeg",
      "verified": false,
      "password": "$2a$10$5bMojbONIM9KR3IsqvGQEO/eMOuKpRaJ8/qrSkphnaSKNkJeXbkw2",
      "role": "user",
      "created_at": "2023-09-25T07:42:26.808Z",
      "updated_at": "2023-09-25T07:42:26.804Z"
    }
  },
  "status": "success"
}```

- User login request example:
`curl -v  POST http://localhost:8000/api/auth/login -H "Content-type: application/json" -d '{"email": "hello@world", "password":"hello"}'`

Response body:
{
  "token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNWZkNTU1MTUtZjYxMS00MTgyLTgwOGUtZjgwY2E1MjNkM2MzIiwidXNlcm5hbWUiOiJ3b3JsZCIsImV4cCI6MTY5NTcxNjk0MCwiaWF0IjoxNjk1NjMwNTQwfQ.JEJkT1vQs_WWFZ_fAPe2i1ScZavD0LgQOGzVJH-coXo"
}

Token usage for protected endpoint:
`curl -X GET "http://localhost:8000/protected/userinfo" -H "Authorization: Bearer <token generated by login endpoint>"`

Response body:

{
  "user": {
    "id": "5fd55515-f611-4182-808e-f80ca523d3c3",
    "name": "world",
    "email": "hello@world",
    "photo": "default.jpeg",
    "verified": false,
    "password": "$2a$10$CqAghRSTwM3vNtGby4aDoOIf4ezGuJA4oUzNEV4oqgic3ORN9.RM2",
    "role": "user",
    "created_at": "2023-09-25T08:28:42.302Z",
    "updated_at": "2023-09-25T08:28:42.3Z"
  }
}
![API Diagram](./api.drawio.svg)
