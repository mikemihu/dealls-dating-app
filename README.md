# Dealls Dating App

A prototype of Dating app.


## Pre-Requisite

- Go v1.22.3 or higher
- Docker
- [wire](https://github.com/google/wire) – dependency injection
- [goose](https://github.com/pressly/goose) – database migration tool

## How to Run

1. Copy config file
   ```shell
   cp config/config.sample.json config/config.development.json
   ``` 
2. Spin up database & redis with docker
   ```shell
   docker-compose up -d
   ```
3. Run app
   ```shell
   make run
   ```

## Directory Structure

In an attempt to achieve Clean Architecture, here is the list for directory purposes.
```shell
cmd/                entry point for the app
config/             configuration file in .json format
docs/               all documentation service related
internal/           all module for the app
 ├─ app/            handle app startup & shutdown, registers background cron and web route
 ├─ config/         define and load config schema
 ├─ constant/       all constant used in the app
 ├─ delivery/       request, response and middleware handler
 ├─ entity/         app wide entities such models, request & response, filters, etc 
 ├─ pkg/            utility package coupled with the app
 ├─ provider/       dependency provider and injection
 ├─ repository/     interface to data source, eg: postgres, redis.
 └─ usecase/        business logic
migrations/         database migration in .sql file format
pkg/                generic utility package, decoupled from the app
```


## Postman
API request & response document

https://documenter.getpostman.com/view/893849/2sA3XJkQG8

## Commands

Found all existing commands by run `make help` command.
```
gen-mock                       Generate mocks
help                           Show this help
migrate                        Run db migration up
migrate-down                   Run db migration down
run                            Run the app
test                           Run unit test with coverage info
wire                           Generate wire
```

Create new migration command
```shell
cd migrations
goose -s create <file_name> sql
```
