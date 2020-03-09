# {{.AppName}}

## Golang Spell
The project was initialized using [Golang Spell](https://github.com/danilovalente/golangspell).

## Architectural Model
The Architectural Model adopted to structure the application is based on The Clean Architecture.
Further details can be found here: [The Clean Architecture](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html) and in the Clean Architecture Book.

## Environment variables
export PORT=8080

export APP_NAME={{.AppName}}

export LOG_LEVEL=INFO

## Dependency Management
The project is using [Go Modules](https://blog.golang.org/using-go-modules) for dependency management
Module: {{.ModuleName}}

## Test and coverage
TESTRUN=true go test -race ./... -coverprofile=cover.out

go tool cover -html=cover.out

## Docker Build

docker build .