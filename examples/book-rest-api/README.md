# book-rest-api
This is example of rest api using clean architecture with 2 entities: 
* book
* author

This service had a purpose to serve a CMS used by operational team to manage master data of **book** and **author**. Beside that this service also used by end user service client to to show author's profile.

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
* Multiple usecases:
    * Simple CRUD/Create Read Update Delete (will be used by CMS)
    * Profile Usecase (will be used by end user client service)

**List of endpoints**
```bash
GET    /cms/author
GET    /cms/author/:id
POST   /cms/author
GET    /cms/book
GET    /cms/book/:id
POST   /cms/book
GET    /author/:id/profile
POST   /search/author
```
