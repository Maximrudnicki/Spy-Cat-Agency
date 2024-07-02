# Spy Cat Agency by Maxim Rudnitskiy

## How to start

In order to start the app you need to create file with name of test.env with following variables:
POSTGRES_HOST=
POSTGRES_PORT=
POSTGRES_DB=
POSTGRES_USER=
POSTGRES_PASSWORD=
PORT=8000



docker build -t test_rudnytskyi:latest .
docker run --env-file test.env -p 8000:8000 test_rudnytskyi:latest