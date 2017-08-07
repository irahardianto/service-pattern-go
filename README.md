service-pattern-go
-------

Hey! welcome, this is and example of clean architecture written in Go Lang with complete Dependency Injection along with Mocking example.

Inspired by [Manuel Kiessling go-cleanarchitecture](http://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/) and [Joshua Partogi TDD training session](https://github.com/jpartogi/tennis-kata-laravel/)

It has simple dependencies:

 - [Gorilla Mux (Router)](https://github.com/gorilla/mux)
 - [Jinzhu GORM (ORM)](https://github.com/jinzhu/gorm)
 - [Testify (Test & Mock framework)](https://github.com/stretchr/testify)

Get Started:

 - [Install](https://irahardianto.github.io/service-pattern-go/#install)
 - [Introduction](https://irahardianto.github.io/service-pattern-go/#introduction)
 - [Folder Structure](https://irahardianto.github.io/service-pattern-go/#folder-structure)
 - [Naming Convention](https://irahardianto.github.io/service-pattern-go/#naming-convention)
 - [Depency Injection](https://irahardianto.github.io/service-pattern-go/#dependency-injection)
 - [Mocking](https://irahardianto.github.io/service-pattern-go/#mocking)


----------

[Install](https://irahardianto.github.io/service-pattern-go/#install)
-------

Clone the source

    git clone https://github.com/irahardianto/service-pattern-go

Setup dependencies

    go get -u github.com/gorilla/mux
    go get -u github.com/jinzhu/gorm
    go get github.com/stretchr/testify

Setup sqlite data structure

    sqlite3 /var/tmp/gorm.db < setup.sql

Run the app, and visit

    http://localhost:8080/getPlayer?playerId=101


----------

[Introduction](https://irahardianto.github.io/service-pattern-go/#introduction)
-------
This is an example of Go clean architecture implementing Dependency Injection and Mocking for unit testing purposes to achieve safe, reliable and secure source code.

The idea of the pattern itself is to create decoupled systems that the implementation of lower level domain is not a concern of the implementor, and can be replaced without having concern of breaking implementor function.

The aim of the architecture is to produce a system that are:

 - Independent of frameworks. The system should be able to become an independent system, not bound into any framework implementation that cause the system to be bloated, instead those framework should be used as a tools to support the system implementation rather than limiting the system capabilities.
 - Highly testable. All codes are guilty and tests is the only way we can prove it otherwise, this means that our test coverage has to be able to cover as much layers as we can so we can be sure of our code reliability.
 - Independent of database. Business logic should not be bound to the database, the system should be able to swap Maria DB, Mongo DB, Dynamo DB without breaking the logic.
 - Independent of 3rd party library. No 3rd party library should be implemented directly to the system logic, we should abstract in away that our system can replace the library anytime we want.

Every implementation should only be by using interface, there should be no direct access from the implementor to implementation, that way we can inject its dependency and replace it with mock object during unit tests. For example:

 - PlayerController -> implement IPlayerService,  instead of direct PlayerService

        type PlayerController struct {
          PlayerService interface.IPlayerService
        }

        func (controller *PlayerController) GetPlayer(res http.ResponseWriter, req *http.Request) {
        	playerId, _ := strconv.Atoi(req.FormValue("playerId"))
        	player := controller.PlayerService.FindById(playerId)
        	playerVM := controller.PlayerHelper.BuildPlayerVM(player)
          json.NewEncoder(res).Encode(playerVM)
        }

 - PlayerService -> implement IPlayerRepository, instead of direct PlayerRepository

       type PlayerService struct {
         PlayerRepository interfaces.IPlayerRepository
       }

       func (repository *PlayerService) FindById(playerId int) models.Player {

         player := repository.PlayerRepository.GetPlayerById(playerId)

         return player
       }

This way, we can switch the implementation of IPlayerService & IPlayerRepository during the injection with whatever implementation without changing the implementation logic.

Router that is used should only the one that **net/http** compatible, that way we can use **net/http/httptest** to unit test our controllers and be sure that we have proper implementation, there are other routers that offers more performance, but if we have to test them with ServeHTTP function, that means we are doing integration tests instead of unit tests.

----------

[Folder Structure](https://irahardianto.github.io/service-pattern-go/#folder-structure)
-------
    /
    |- controllers
    |- infrastructures
    |- interfaces
    |- models
    |- repositories
    |- services
    |- viewmodels
    main.go
    router.go

The folder structure is created to accomodate seperation of concern principle, where every struct should have single responsibility to achieve decoupled system.

Every folder is a namespace of their own, and every file / struct under the same folder should only use the same namepace as their root folder.

**infrasctructures**


**controllers**

**interfaces**

**models**

**repositories**

**services**

**viewmodels**

----------

[Naming Convention](https://irahardianto.github.io/service-pattern-go/#naming-convention)
-------
- Namespace
- Struct
- Interface
- Service
- Repository
- Controller

----------

[Dependecy Injection](https://irahardianto.github.io/service-pattern-go/#dependency-injection)
-------


----------

[Mocking](https://irahardianto.github.io/service-pattern-go/#mocking)
-------

Cheers,
M. Ichsan Rahardianto.
