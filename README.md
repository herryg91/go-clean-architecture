# go-clean-architecture
The right way to implement clean architecture on golang (I think) :yum:. Just kidding, let's discuss together.

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
    * Repository Interface. Bridging driver out and use cases layer.
    * Use Case Interface. Bridging driver in and use cases layer.
* Drivers Layers. Will be splitted into two:
    * Driver In (handler, grpc server handler, etc)
    * Driver Out (external apis/interfaces, datasource, etc.)

## Design Structure
### Diagram
![picture alt](https://raw.githubusercontent.com/herryg91/go-clean-architecture/main/diagram.png "go-clean-architecture diagram")

### Folder Structure

    .
    ├── drivers                         # outermost layer to communication with external such as db, api, etc.
    │   ├── datasource                  # driver-out. for communication with database
    │       ├── mysql                    
    │       ├── elastic
    │   ├── external                    # driver-out. for communication with other services (api/grpc/etc.)
    │       ├── api
    │       ├── grpc
    │   ├── handler                     # driver-in. implementation http/grpc handler
    ├── entity                          # encapsulate Enterprise wide business rules (struct with methods)
    │   ├── entity1.go
    │   ├── entity2.go
    │   ├── ...
    ├── pkg                             # Supporting library / script will be written here
    │   ├── helpers                     
    │   ├── password                    
    ├── usecase                         # usecase can be separate
    │   ├── usecase1          
    │       ├── interface.go            # interface adaptors: repository & usecase interface will be written here 
    │       ├── repository-impl.go      # implement repository interface
    │       ├── usecase-impl.go          # implement usecase interface
    │   ├── usecase2                     
    │       ├── interface.go            
    │       ├── repository-impl.go      
    │       ├── usecase-impl.go        
    │   ├── ...          
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── service.yaml
    └── README.md

## Postman Collection
https://www.getpostman.com/collections/616edf710f86853574d2

## Future examples
* Rest API service with background worker
* Rest API service with event driven system (pubsub)
* Clean architecture on Command line interface (CLI)
* Rest API service with external dependency
* Clean architecture in Contract Driven Development
* etc.