# ChGo

> Simple to do list API with Gin and Gorm (with Postgres)

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

| Method | Route               | Header                                      | Body                                         |
| ------ | ------------------- | ------------------------------------------- | -------------------------------------------- |
| GET    | /v1/tasks           | `{Authorization: "Bearer ${Token}"}`        |                                              |
| POST   | /v1/tasks           | `{Authorization: "Bearer ${Token}"}`        | `{"title": "task title"}`                    |
| DELETE | /v1/tasks/:id       | `{Authorization: "Bearer ${Token}"}`        |                                              |
| PUT    | /v1/tasks/:id       | `{Authorization: "Bearer ${Token}"}`        | `{"title": "task title", "completed": true}` |
| POST   | /v1/noauth/register |                                             | `{"Usename": "admin", "Password": "admin"}`  |
| POST   | /v1/noauth/login    |                                             | `{"Usename": "admin", "Password": "admin"}`  |
| GET    | /v1/noauth/refresh  | `{Authorization: "Bearer ${RefreshToken}"}` | `{"Usename": "admin", "Password": "admin"}`  |

## Development

### Running locally

```
docker-compose down
./run.sh
```

Using _adminer_ access: `localhost:8080`
Username: `postgres`
Password: `c9BqhGZM5v7EPTs7`

## Build docker image for ChGo

```
./build.sh
```
