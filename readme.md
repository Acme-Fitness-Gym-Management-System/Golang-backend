# Gym-backend

<a href="https://github.com/JaswanthKarangula/Gym-backend/actions">
      <img alt="Tests Passing" src="https://github.com/JaswanthKarangula/Go-Banking/actions/workflows/test.yml/badge.svg?event=push" />
    </a>

## 1 Creating a Database


```docker pull postgres:14-alpine```
docker images

```docker run --name postgres  -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine```
docker ps


docker exec -it postgres psql -U root





- Use dbdiagram.io to create a database schema and import the .sql file
- use the .sql file to create the database
  <img src="/dbschema.png"/>

[//]: # ()
[//]: # (``` psql postgres   ``` <br>)

[//]: # (``` create role Jaswanth with login password 'password'  ```)

[//]: # (<br>)

[//]: # ()
[//]: # ()
[//]: # (``` psql postgres -U jaswanth ```)

[//]: # (<br>)

[//]: # (``` create database simple_bank ```)

## 2 Database Migration

- Can install go-migrate using brew install golang-migrate  or use it as a library
```
import (
      "github.com/golang-migrate/migrate/v4"
      _ "github.com/golang-migrate/migrate/v4/database/postgres"
      _ "github.com/golang-migrate/migrate/v4/source/github"
      )

        func main() {
        m, err := migrate.New(
        "github://mattes:personal-access-token@mattes/migrate_test",
        "postgres://localhost:5432/database?sslmode=enable")
        m.Steps(2)
        } 

```
- ``` migrate -help ```
- Create a Schema up and down files  ```  migrate create -ext sql -dir db/migration gym-backendsql  ```
- Create db using make createdb
- Update the migrate up and down files
- Update Makefile for migrateup and down

## 3  Using SQLC


brew install kyleconroy/sqlc/sqlc
<br>
or
<br>
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

sqlc init 
```sqlc generate``` to generate the code or add it to make file and use make generate


Get Gin using ``` go get -u github.com/gin-gonic/gin```

GOMOCK for mocking db in tests

```go install github.com/golang/mock/mockgen@v1.6.0``` install mockgen and ``` which mockgen ``` If not visible add Go to Path



Mocking database to generate tests

mockgen -package mockdb -destination db/mock/store.go  github.com/JaswanthKarangula/Go-Banking/db/sqlc Store

## Postgres Driver 

go get github.com/lib/pq


github.com/spf13/viper


## Deploying

Multi stage dockerfile 


docker build -t gymbackend:latest .    
docker run --name gymbackend -p 8080:8080 jaswanthk1/gym-backend:latest

docker container inspect postgres
docker container inspect gymbackend
docker network inspect bridge

docker network create gym-network
docker network connect gym-network postgres
docker network inspect gym-network
docker container inspect postgres

docker run --name gymbackend --network bank-network -p 8080:8080 -e DB_SOURCE="postgresql://root:secret@postgres:5432/simple_bank?sslmode=disable" gymbackend:latest




## Gin -- 

github.com/gin-gonic/gin
viper for config

go run main.go


## Swagger 

install swagger cmd
go get -u github.com/swaggo/files
go get -u github.com/swaggo/gin-swagger


## AWS Github Actions

AWS IAM












