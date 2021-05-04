# book-rest-api
This is example of rest api using clean architecture with single entity: `book`. In this example, more complex than video-rest-api case since the entity in this service is not represent database schema.

In this microservice, we're gonna use both mysql and redis which will be aggregated in repository to serve what the microservice needs based on the usecase.

**How to Run?
Prerequisite:
- docker
- golang

```
make build
go run main.go
```

**Usecase solved in this example**
* Aggregate 2 or more database's tables
* Handle Join Query
* Using redis + mysql in repository

**List of endpoints**
```bash
GET    /book/:id
POST   /book/search
```
