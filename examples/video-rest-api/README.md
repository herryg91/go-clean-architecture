# video-rest-api
Just a simple example of golang clean architecture implementation using 1 entities. This is just a simple CRUD service (Create Read Update Delete) using mysql as the database.

**How to Run?
Prerequisite:
- docker
- golang

```
make build
go run main.go

// for integration test
make test
```

**Usecase in this example**
* Simple CRUD service

**List of endpoints**
```bash
GET    /video
GET    /video/:id
POST   /video
POST   /video/:id
DELETE /video/:id
```