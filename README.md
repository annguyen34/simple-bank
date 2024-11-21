# SIMPLE BANK

This repository contains a project I have learn about design, develop and deploy a complete backend system using Go, PostgreSQL and Docker.

## Setup local development

#### Install tools

- [Docker desktop](https://www.docker.com/products/docker-desktop)
- [TablePlus](https://tableplus.com/)
- [Golang](https://golang.org/)
- [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
- [DB Docs](https://dbdocs.io/docs)
- [DBML CLI](https://www.dbml.org/cli/#installation)
- [Sqlc](https://github.com/kyleconroy/sqlc#installation)
- [Gomock](https://github.com/golang/mock)
- [Postman](https://www.postman.com/)

## Quick start

Clone this repository:

```sh
git clone https://github.com/annguyen34/simple-bank.git
cd simple_bank
```

#### Using with docker compose

Run this command:

```sh
docker compose up
```

Now server is started on [http://localhost:8000/](http://localhost:8000/) for HTTP server and [http://localhost:9000/](http://localhost:9000/) for gRPC server. You can use Postman for testing purpose.

## API endpoints

#### With gateway server

- `POST /v1/create_user`: Create a new user
- `POST /v1/login_user`: Login user and get access token & refresh token

#### With Gin server

You can change to Gin server by using function `runGinServer` in `main.go`

- `POST /signup`: Create a new user
- `POST /login`: Login user and get access token & refresh token
- `POST /tokens/renew_access`: Return new access token
- `POST /accounts`: Create a new account
- `GET /accounts`: Get account list belong to users
- `GET /accounts/:id`: Get an account with account's ID
- `POST /transfers`: Transfer money between 2 accounts

## Testing

Run testing for all packages with random data

```sh
make test
```

## Setup

- Start postgres container:

  ```bash
  make postgres
  ```

- Create simple_bank database:

  ```bash
  make createdb
  ```

- Run db migration up all versions:

  ```bash
  make migrateup
  ```

- Run db migration up 1 version:

  ```bash
  make migrateup1
  ```

- Run db migration down all versions:

  ```bash
  make migratedown
  ```

- Run db migration down 1 version:
  ```bash
  make migratedown1
  ```

## Documentation

#### Database

- Generate DB documentation:

  ```bash
  make dbdocs
  ```

#### Swagger (API Documentation)

- Go to [http://localhost:8000/swagger](http://localhost:8000/swagger) to view API documentation

## How to generate code

- Generate schema SQL file with DBML:

  ```bash
  make db_schema
  ```

- Generate SQL CRUD with sqlc:

  ```bash
  make sqlc
  ```

- Generate DB mock with gomock:

  ```bash
  make mock
  ```

- Create a new db migration:
  ```bash
  migrate create -ext sql -dir db/migration -seq <migration_name>
  ```
