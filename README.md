# API Practice

A practice project with the same object API implemented in .NET, Go, Java, and
Python, plus a small frontend. Docker Compose builds and runs the services.

## Requirements

- Docker Desktop (or Docker Engine) with the Compose plugin

## Start the services

From the project root, run:

```sh
docker compose up --build
```

Open the frontend at <http://localhost:8888>. Its API requests are proxied by
Nginx to the Python backend, so the browser does not need to call the API
directly across origins.

To stop the services, press `Ctrl+C`, then run:

```sh
docker compose down
```

## Service endpoints

| Service | Host URL | Container port |
| --- | --- | ---: |
| Frontend | <http://localhost:8888> | 8888 |
| Python API | <http://localhost:8000> | 8000 |
| Go API | <http://localhost:8081> | 8081 |
| .NET API | <http://localhost:8080> | 8080 |
| Java API | <http://localhost:8082> | 8080 |

Each backend implements object operations under `/objects`:

| Method | Path | Operation |
| --- | --- | --- |
| GET | `/objects` | List objects |
| GET | `/objects/{id}` | Get one object |
| POST | `/objects` | Create an object |
| PUT | `/objects/{id}` | Replace an object |
| PATCH | `/objects/{id}` | Partially update an object |
| DELETE | `/objects/{id}` | Delete an object |

For example, query the Python API with:

```sh
curl http://localhost:8000/objects
```

The frontend buttons make requests through the Python API and log responses to
the browser developer console.

## Useful Compose commands

```sh
# View service logs
docker compose logs -f

# Rebuild and start after changing source
docker compose up --build

# Stop and remove the containers and Compose network
docker compose down
```

The APIs use in-memory sample data. Changes made through create, update, or
delete operations are not persisted when a backend container restarts.
