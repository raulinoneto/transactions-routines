# Transaction Routines
This project provides an API like a Real World payment processor, it was proposed for a Pismo Challenge and I felt very 
challenged and excited to do it.
In this challenge I forced myself to apply the greater quantity of my knowledge and best pratices that I know, sometimes
I couldn't apply these things, but I think that what have at the test sounds great.


## Instalation
Requirements:
```shell script
docker 19 or later
docker-compose 1.25 or later
```

Instalation:
```shell script
make run-docker-clean
make migrate
```

## Project Organization 

### Third Part Packages

* [github.com/go-chi/chi](github.com/go-chi/chi) - Mux for prepare http requests.
* [github.com/go-sql-driver/mysql](github.com/go-sql-driver/mysql) - MySQL driver needed to run native database connector.
* [github.com/reactivex/rxgo/v2](github.com/reactivex/rxgo/v2) - Was used to store blocked accounts that are in processing. 

### Transactions Flow
![Transaction Flow Diagram](https://github.com/raulinoneto/transactions-routines/blob/master/images/transaction-diagram.png "Transaction Flow Diagram")

### Hexagonal Architecture
Hexagonal Architecture consists in to divide an application in layers with your resposabilities and to focus in the 
logical business layer, the domain
Hexagonal Architecture are called "Ports and Adapters Pattern" too.

![Hexagonal Architecture](https://github.com/raulinoneto/transactions-routines/blob/master/images/hexagonal-architecture.png "Hexagonal Architecture")

### Source tree explained
```text
.
├── api     # Postman API to import
│   └── Transactions.postman_collection.json
├── bin # Stores the binaries of compiled code
├── build # Docker configuration
│   └── docker
│       ├── Dockerfile 
│       └── start.server.sh
├── cmd # Entrypoints for the application
│   └── http # Http entrypoint
│       └── main.go
├── configs # Configuration needed for the application
│   ├── container # Load all application injections
│   │   ├── container.go # Container Injection structure
│   │   ├── primaryadapterscontainer.go # Injections about primary adapters
│   │   ├── secondaryadapterscontainer.go # Injections about secondary adapters
│   │   └── servicescontainer.go # Injections about services (domains)
│   └── routes # Applications routes
│       ├── accountsroutes.go 
│       ├── routes.go
│       └── transactionsroutes.go
├── docker-compose.yml
├── go.mod
├── go.sum
├── internal # Dependencies for application, here are no logical business
│   ├── apierror # Package to prettify error responses
│   │   └── apierror.go
│   ├── primary # Primary adapters, who calls the applications
│   │   ├── httpadapters # Adapters for the http
│   │   │   ├── accountbody.go # The entry data model for account
│   │   │   ├── accountshttpadapter.go 
│   │   │   ├── helpers.go 
│   │   │   ├── httpadapter.go
│   │   │   ├── responses.go # Define behavior for adapters
│   │   │   ├── transactionsbody.go # The entry data model for transactions
│   │   │   └── transactionshttpadapter.go
│   │   └── observer # ReactiveX Observer
│   │       └── observer.go
│   └── secondary # Secondary Adapters
│       ├── persistence # Database storage
│       │   ├── accountsmodel.go # Define de Database Model to be stored
│       │   ├── accountsrepositoryadapter.go
│       │   ├── mysqladapter.go 
│       │   └── transactionsrepositoryadapter.go
│       └── rx # ReactiveX Store
│           └── transactionsobserver.go
├── LICENSE
├── Makefile
├── pkg # Where are the all logical business from the application
│   ├── core # Package that can be shared by domains
│   │   └── enum.go
│   └── domains 
│       ├── accounts # Account logical business
│       │   ├── accounts.go # Interfaces to comunicate with accounts Domain
│       │   ├── service.go
│       └── transactions
│           ├── service.go
│           └──  transactions.go # Interfaces to comunicate with accounts Domain
├── README.md
└── scripts
    └── database.sql # Database to migrate

```

## Biography

1. [My Talk About Hexagonal Architecture](https://docs.google.com/presentation/d/1nEpfDEfnwGB3Xy-7CMfccW7L2qZeo4UVR738434bVVY/edit?usp=sharing)
2. [My Talk About Postman](https://docs.google.com/presentation/d/1SHUSATWs-vOkScWXm6ae4vgjomEKeDrMkm0JQ1fXpRw/edit?usp=sharing)
3. [Clean Architecture - Uncle Bob](https://www.amazon.com.br/Clean-Architecture-Craftsmans-Software-Structure-ebook/dp/B075LRM681/)
4. [The Proposed Project Structure](https://github.com/golang-standards/project-layout)