# ChGo

> Simple to do list API with Gin and Gorm (with Postgres)

## Docker

Clone this repository and run:

```
./build.sh
docker-compose up
```

You can then hit the following endpoints:

| Method | Route             | Body                                         |
| ------ | ----------------- | -------------------------------------------- |
| GET    | /api/v1/tasks     |                                              |
| POST   | /api/v1/tasks     | `{"title": "task title"}`                    |
| DELETE | /api/v1/tasks/:id |                                              |
| PUT    | /api/v1/tasks/:id | `{"title": "task title", "completed": true}` |

## Development

### Running locally

```
docker-compose down
./run.sh
```

Using _adminer_ access: `localhost:8080`
