# simple-book-rest-api
This is example of rest api using clean architecture with 2 entities: 
* book
* author

This service had a purpose to serve a CMS used by operational team to manage master data of **book** and **author**. Beside that this service also used by end user service client to to show author's profile.

**Usecase solved in this example**
* Aggregate 2 or more database's tables
* Multiple usecases:
    * Simple CRUD/Create Read Update Delete (will be used by CMS)
    * Profile Usecase (will be used by end user client service)
