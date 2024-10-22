# Pismo Challenge

Implementing Pismo challenge

> Go version 1.22 \
> Postgres:14.2

## Setup

You will need Make, migrate (`brew install golang-migrate`) and Docker to better run this environment.
HTTP Requests covering the API are available in: `/test/development/pismo`

To run the application:
```
cp .env.example .env
make up
make migrate
```

Documentation is available in http://localhost:9090/docs/
