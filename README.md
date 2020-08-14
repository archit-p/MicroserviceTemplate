[![Build Status](https://img.shields.io/travis/archit-p/MicroserviceTemplate)](https://travis-ci.org/archit-p/MicroserviceTemplate)
[![Go Report Card](https://goreportcard.com/badge/github.com/archit-p/MicroserviceTemplate)](https://goreportcard.com/report/github.com/archit-p/MicroserviceTemplate)
![Codacy grade](https://img.shields.io/codacy/grade/6ab701e403dd4aa39544b6b72be52506)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=MicroserviceTemplate&metric=alert_status)](https://sonarcloud.io/dashboard?id=MicroserviceTemplate)  
Go Microservice built using best practices, ideal for use as a starter template. It features an extensible model - Sample - with support for CRUD and Search operations for MongoDB.

## References
This project is inspired by multiple blogs and guides on best practices while writing Go code. I'm sharing links to these below.
### Project Structure
1.  [How to Write Go Code](https://golang.org/doc/code.html)
2.  [golang-standards/project-layout](https://github.com/golang-standards/project-layout)
3.  [Organizing Database Access](https://www.alexedwards.net/blog/organising-database-access)
### Documentation
1.  [Documenting APIs with Swagger](https://swagger.io/resources/articles/documenting-apis-with-swagger/)
2.  [Documenting Go Code](https://blog.golang.org/godoc)
### Containerization
1.  [Docker: Multi-Stage Builds](https://docs.docker.com/develop/develop-images/multistage-build/)
2.  [Why You Should Use Microservices and Containers](https://developer.ibm.com/technologies/microservices/articles/why-should-we-use-microservices-and-containers/)
### Testing
1.  [Structuring Tests in Go](https://medium.com/@benbjohnson/structuring-tests-in-go-46ddee7a25c)
2.  [Testing Web-Apps in Go](https://markjberger.com/testing-web-apps-in-golang/)
## Getting Started
### Clone the Repo
```sh
$ git clone https://github.com/archit-p/MicroserviceTemplate
```
### Build and Run
```sh
$ make help
Run make <target> where target is
	help: print out this message
	build: build the executables
	run: start a clean build, and run executable
	test: run go tests
	docs: build documentation
	clean: clean executables and docs
$ make run
```
## Project Layout
```text
cmd             (contains code for our apps)
|-+ web         (server router and controllers)
pkg             (contains reusable code)
|-+ dto         (data-transfer objects)
|-+ models      (database models)
    |-+ mongodb (models for mongo-db)
|-+ docs        (swagger documentation)
```
## Docs
Once the service is running, accompanying Swagger docs can be found at `http://localhost:8080/swagger/index.html`.

## Contributing
Feel free to fork the project for use as a base template for your next microservice or REST API!

## Built Using
[![Golang](https://www.vectorlogo.zone/logos/golang/golang-ar21.svg)](https://golang.org)
[![MongoDB](https://www.vectorlogo.zone/logos/mongodb/mongodb-ar21.svg)](https://www.mongodb.com/)
[![Docker](https://www.vectorlogo.zone/logos/docker/docker-icon.svg)](https://www.docker.com/)
