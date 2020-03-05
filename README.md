# ChGo

> Backend list API with Gin and Gorm (with Postgres)

## Migrate database

```
./migrate.sh
```

## Seed data

```
./seed.sh
```

## Docker

Stop & remove older docker container

```
docker stop $(docker ps -aq)
docker system prune -aq
```

Clone this repository and run:

```
docker-compose up
```

You can then hit the following endpoints:

| Method | Route               | Header                                      | Body                                            |
| ------ | ------------------- | ------------------------------------------- | ----------------------------------------------- |
| GET    | /v1/products        | `{Authorization: "Bearer ${Token}"}`        |                                                 |
| POST   | /v1/products        | `{Authorization: "Bearer ${Token}"}`        | `{"title": "product title"}`                    |
| DELETE | /v1/products/:id    | `{Authorization: "Bearer ${Token}"}`        |                                                 |
| PUT    | /v1/products/:id    | `{Authorization: "Bearer ${Token}"}`        | `{"title": "product title", "completed": true}` |
| POST   | /v1/noauth/register |                                             | `{"Usename": "admin", "Password": "admin"}`     |
| POST   | /v1/noauth/login    |                                             | `{"Usename": "admin", "Password": "admin"}`     |
| GET    | /v1/noauth/refresh  | `{Authorization: "Bearer ${RefreshToken}"}` | `{"Usename": "admin", "Password": "admin"}`     |

## Development

### Running locally

```
docker-compose down
./run.sh
```

Using _adminer_ access: `localhost:8080`

- Username: `postgres`
- Password: `c9BqhGZM5v7EPTs7`

## Build docker image for ChGo

```
./build.sh
```
