# go-clean-architecture
The right way to implement clean architecture on golang (I think) :yum:. Promotion stuff: to help you easier writing code you can check https://kodingless.com . This platform was made inspired from this go-clean-architecture

## Description
The Clean Architecture is a software architecture proposed by Robert C. Martin (better known as Uncle Bob). In this repository, the contents are list of examples of implementation of Clean Architecture in Golang. The examples will using real world scenario. 

New in clean architecture? I suggest you to look at examples/video-rest-api: https://github.com/herryg91/go-clean-architecture/tree/main/examples/video-rest-api

To understand more about Clean Architecture: https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html

## Dependency Rule
Based on Uncle Bob, there are 4 layers:
* Entity
* Use Cases
* Interface Adapters
* Frameworks and Drivers.

In this repository, we also using 4 layers (with modification) like this:
* Entity. 
* Use Cases (Implementation).
* Interface Adapters. Will be splitted into two:
    * Repository Interface. Bridging repository implementation and use cases layer.
    * Use Case Interface. Bridging handler and use cases layer.
* Driver Layer
    * Handler
    * Repository implementation

## Design Structure
### Diagram
<img src="https://raw.githubusercontent.com/herryg91/go-clean-architecture/main/diagram.png" alt="drawing" width="400"/>

### Folder Structure
    .
    ├── app                             # inner layer of the pattern. we focus on writting the logic here
    │   ├── repository                  # repository interface. The software in this layer is a set of adapters that convert data from/to the format most convenient for the use cases and entities, from/to the format most convenient for some external agency such as the Database or the Web
    │       ├── video_repo.go
    │       ├── user_repo.go
    │   ├── usecase                     # usecase layer
    │       ├── usecase1
    │           ├── errors.go
    │           ├── implement.go
    │           ├── interface.go
    │       ├── usecase2
    │           ├── errors.go
    │           ├── implement.go
    │           ├── interface.go
    ├── client                          # third party, api client, grpc client/wrapper
    │   ├── test1-api
    │       ├── dto.go
    │       ├── client.go
    │   ├── test2-api
    │   ├── test3-grpc
    │   ├── test4-grpc
    ├── config
    ├── entity                          # encapsulate Enterprise wide business rules (struct with methods)
    │   ├── entity1.go
    │   ├── entity2.go
    │   ├── ...
    ├── handler                         # outermost layer / drivers which driving the business logic. implementation http/grpc, worker handler, etc.
    │   ├── api.go
    │   ├── subscribe.go
    │   ├── worker.go
    ├── pkg                             # Supporting library / script
    │   ├── helpers                     
    │   ├── password                    
    │   ├── ...                    
    ├── repository                      # implementation of repository. in this repository we can aggregate from db, client, external agency, etc. into convenient format for this service
    │   ├── video_repository_v1
    │       ├── model.go
    │       ├── repository.go
    │   ├── video_repository_v2
    │       ├── model.go
    │       ├── repository.go
    │   ├── user_repository_v1
    │       ├── model.go
    │       ├── repository.go
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── service.yaml
    └── README.md

## Guidelines
### Step-by-step writing code using this pattern
- Setup skeleton of the microservices (including: main.go, migrations, config, pkg and handler folder)
- Defining the `entities`
- Defining `usecase` (interface & implementation) in folder `app`. We're gonna focus in this folder since the business logic will be written here.
- When the usecase need to communicate to the external agency (Database, other apis, etc) then write it to the `repository interface`
- After the usecase layer was done, now time to write `repository implementation`
- Put it up together + register to the main.go and handler folder

### How to Define Entity
Entities encapsulate Enterprise wide business rules. An entity can be an object with methods, or it can be a set of data structures and functions. Therefore, we can got a clue to defining the entity by look into the business rule. My suggestion is analyze the outermost layer of the service (handler folder), then look into the param and the output struct. 

```
example:

GET http://{host}/profile/{Id}

return 200
{
    "id": 1,
    "name: "test",
    "age: 20,
}

Look, the output. Actually it's the indication that we need `Profile` entity. In golang we can write it like this:
type Profile struct{
    Id int
    Name string
    Age int
}

```

### How to Define Repository
Repository in this design pattern had a task to convert data from/to the format most convenient entities and usecases, from/to the format most convenient for external agencies such as Database, Api/Grpc/ etc. By that definition, my suggestion is to define the repository by `domain/aggregate root`. 1 domain/aggregate root = 1 repository

To define the domain itself, we can list down the entities and group it together by looking to the common requirement/terminology/functionality (aggregate root). https://docs.microsoft.com/en-us/dotnet/architecture/microservices/microservice-ddd-cqrs-patterns/infrastructure-persistence-layer-design

Example, we have entity like this:
- user
- user_contact
- user_parent
- login_token
Then, we can put `user`, `user_contact`, `user_parent` into domain `user` and create `user_repository`. We can put `login_token` into `token_repository`


### How to Define UseCase
Like it's name, usecase strongly tied to user story/journey. Therefore the point of view is user centric. To get a clue about how to define the usecase, I suggest you to look up into the Product Requirement or UseCase diagram or List of endpoint you need to serve in your microservice.

Example, we have users-api which need to serve endpoint api like this:
- Login (POST http://{host}/auth/login)
- Logout (POST http://{host}/auth/logout)
- Register (POST http://{host}/auth/register)
- ShowProfile (POST http://{host}/profile)
- EditProfile (PUT http://{host}/profile)

then, based on the product requirement we can split it into 2 usecase: `AuthenticationUsecase` (Login, Logout, Register) and `ProfileUsecase` (ShowProfile, EditProfile)


## Future examples
* Rest API service with external dependency
* Rest API service with background worker
* Rest API service with event driven system (pubsub)
* Clean architecture on Command line interface (CLI)
* Clean architecture in Contract Driven Development
* etc.
