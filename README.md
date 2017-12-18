service-pattern-go
-------

Hey! Welcome, this is an example of simple REST API implementation with clean architecture written in Go with complete Dependency Injection along with Mocking example, following SOLID principles.

Inspired by [Manuel Kiessling go-cleanarchitecture](http://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/) and [Joshua Partogi TDD training session](https://github.com/jpartogi/tennis-kata-laravel/)

It has simple dependencies:

 - [Chi (Router)](https://github.com/go-chi/chi)
 - [Testify (Test & Mock framework)](https://github.com/stretchr/testify)
 - [Mockery (Mock generator)](https://github.com/vektra/mockery)
 - [Hystrix-Go (Circuit Breaker)](https://github.com/afex/hystrix-go)

Get Started:

 - [Install](https://irahardianto.github.io/service-pattern-go/#install)
 - [Introduction](https://irahardianto.github.io/service-pattern-go/#introduction)
 - [Folder Structure](https://irahardianto.github.io/service-pattern-go/#folder-structure)
 - [Depency Injection](https://irahardianto.github.io/service-pattern-go/#dependency-injection)
 - [Mocking](https://irahardianto.github.io/service-pattern-go/#mocking)
 - [Testing](https://irahardianto.github.io/service-pattern-go/#testing)
 - [Circuit Breaker](https://irahardianto.github.io/service-pattern-go/#circuit-breaker)


----------

[Install](https://irahardianto.github.io/service-pattern-go/#install)
-------

Clone the source

    git clone https://github.com/irahardianto/service-pattern-go

Setup dependencies

    go get -u github.com/go-chi/chi
    go get -u github.com/jinzhu/gorm
    go get github.com/stretchr/testify
    go get github.com/vektra/mockery/.../
    go get github.com/afex/hystrix-go/hystrix

Setup sqlite data structure

    sqlite3 /var/tmp/tennis.db < setup.sql

Test first for your liking

    go test ./... -v

Run the app

    go build && ./service-pattern-go

And visit

    http://localhost:8080/getScore/Rafael/vs/Serena


----------

[Introduction](https://irahardianto.github.io/service-pattern-go/#introduction)
-------
This is an example of Go clean architecture implementing Dependency Injection and Mocking for unit testing purposes to achieve safe, reliable and secure source code.

The idea of the pattern itself is to create decoupled systems that the implementation of lower level domain is not a concern of the implementor, and can be replaced without having concern of breaking implementor function.

The aim of the architecture is to produce a system that are:

 - Independent of frameworks. The system should be able to become an independent system, not bound into any framework implementation that cause the system to be bloated, instead those framework should be used as a tools to support the system implementation rather than limiting the system capabilities.
 - Highly testable. All codes are guilty and tests is the only way we can prove it otherwise, this means that our test coverage has to be able to cover as much layers as we can so we can be sure of our code reliability.
 - Independent of database. Business logic should not be bound to the database, the system should be able to swap MySQL, Maria DB, PosgreSQL, Mongo DB, Dynamo DB without breaking the logic.
 - Independent of 3rd party library. No 3rd party library should be implemented directly to the system logic, we should abstract in away that our system can replace the library anytime we want.

Every implementation should only be by using interface, there should be no direct access from the implementor to implementation, that way we can inject its dependency and replace it with mock object during unit tests. For example:

PlayerService -> implement IPlayerRepository, instead of direct PlayerRepository


    type PlayerService struct {
      interfaces.IPlayerRepository
    }

    func (service *PlayerService) GetScores(player1Name string, player2Name string) (string, error) {

      baseScore := [4]string{"Love", "Fifteen", "Thirty", "Forty"}
      var result string

      player1, err := service.GetPlayerByName(player1Name)
      if err != nil {
        //Handle error
      }

      player2, err := service.GetPlayerByName(player2Name)
      if err != nil {
        //Handle error
      }

      if player1.Score < 4 && player2.Score < 4 && !(player1.Score+player2.Score == 6) {

        s := baseScore[player1.Score]

        if player1.Score == player2.Score {
          result = s + "-All"
        } else {
           result = s + "-" + baseScore[player2.Score]
        }
      }

      if player1.Score == player2.Score {
        result = "Deuce"
      }

      return result, nil
    }
    
If you look into the implementation of these lines

    player1, err := service.GetPlayerByName(player1Name)
    player2, err := service.GetPlayerByName(player2Name)

Both are actually abstract implementation of the interface, not the real implementation itself.
So later on the Dependency Injection section, we will learn those interface will be injected with the implementation during the compile time. This way, we can switch the implementation of IPlayerService & IPlayerRepository during the injection with whatever implementation without changing the implementation logic.

Throughout this repo you will find implementation of design patterns such as **Strategy Pattern** when we inject our dependencies with the real implementations. We create **Singleton** and use it to wired up our router and services. We use **Composite** for all our abstract interface implementations so that the implementor can abstractly implement the methods it has, just as the example above where **PlayerService** implements **interfaces.IPlayerRepository** and allows it to directly invoke **GetPlayerByName** which is **IPlayerRepository's** method.  We also use **Decorator Pattern** to hook up our circuit breaker without needing to change / modify the original implementation.

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
    servicecontainer.go

The folder structure is created to accomodate seperation of concern principle, where every struct should have single responsibility to achieve decoupled system.

Every folder is a namespace of their own, and every file / struct under the same folder should only use the same namepace as their root folder.

### controllers

controllers folder hosts all the structs under controllers namespace, controllers are the handler of all requests coming in, to the router, its doing just that, business logic and data access layer should be done separately.

controller struct implement services through their interface, no direct services implementation should be done in controller, this is done to maintain decoupled systems. The implementation will be injected during the compiled time.


### infrasctructures

infrasctructures folder host all structs under infrasctructures namespace, infrasctructures consists of setup for the system to connect to external data source, it is used to host things like database connection configurations, MySQL, MariaDB, MongoDB, DynamoDB.

### interfaces

interfaces folder hosts all the structs under interfaces namespace, interfaces as the name suggest are the bridge between different domain so they can interact with each other, in our case, this should be the only way for them to interact.

interface in Go is a bit different then you might already find in other language like Java or C#, while the later implements interface explicitly, Go implements interface implicitly. You just need to implement all method the interface has, and you're good to "Go".

In our system, our PlayerController implements IPlayerService to be able to interact with the implementation that will be injected. In our case, IPlayerService will be injected with PlayerService.

The same thing applies on PlayerService which implements IPlayerRepository to be able interact with the injected implementation. In our case, IPlayerRepository will be injected with PlayerRepository during the compile time.

PlayerRepository on the other hand, will be injected with infrasctructure configuration that has been setup earlier, this ensure that you can change the implementation of PlayerRepository, without changing the implementor which in this case PlayerService let alone break it. The same thing goes to PlayerService and PlayerController relationship, we can refactor PlayerService, we can change it however we want, without touching the implementor which is PlayerController.

### models

models folder hosts all structs under models namespace, model is a struct reflecting our data object from / to database. models should only define data structs, no other functionalities should be included here.

### repositories

repositories folder hosts all structs under repositories namespace, repositories is where the implementation of data access layer. All queries and data operation from / to database should happen here, and the implementor should be agnostic of what is the database engine is used, how the queries is done, all they care is they can pull the data according to the interface they are implementing.

### services

services folder hosts all structs under services namespace, services is where the business logic lies on, it handles controller request and fetch data from data layer it needs and run their logic to satisfy what controller expect the service to return.

controller might implement many services interface to satisfy the request needs, and controller should be agnostic of how services implements their logic, all they care is that they should be able to pull the result they need according to the interface they implements.

### viewmodels

viewmodels folder hosts all the structs under viewmodels namespace, viewmodels are model to be use as a response return of REST API call

### main.go

main.go is the entry point of our system, here lies the router bindings it triggers ChiRouter singleton and call InitRouter to bind the router.

### router.go

router.go is where we binds controllers to appropriate route to handle desired http request. By default we are using Chi router as it is a light weight router and not bloated with unnecessary unwanted features.

### servicecontainer.go

servicecontainer.go is where the magic begins, this is the place where we injected all implementations of interfaces. Lets cover throughly in the dependency injection section.

----------

[Dependecy Injection](https://irahardianto.github.io/service-pattern-go/#dependency-injection)
-------

Dependecy injection is the heart of TDD, without it we wont be able to do proper TDD because there will be no mocking and we cannot decoupled our code properly. This is one of the misconception when people thinks that they are doing unit testing instead actually they are doing integration test which connects the logic to database. Unit test should be done independently and database should not come in to play when we are doing unit test. One thing to not though, in Go dependency has to be injected during compile time instead of runtime which cause it a bit different than Java / C# implementation, but anyway, its just plain old dependency injection.

In essence unit test is created to test our logic not our data integrity, and by taking database during unit testing it will add huge complexity to the tests itself, and this creates barrier for programmers new to unit testing as they are struggling to create proper testing for their functions.

Now why dependency injection is a crucial part in doing proper TDD? the answer lies in the usage of interface. Back when I have never encountered mocking, I always wondering, what is the use of interface, why we should create abstraction for our functions instead of just write it all already, why the hell should we create a duplicate, abstraction that we will be implementing shortly anyway, some says that, because in doing so, your code will be much cleaner and we have proper pattern, I called that bullshit because in essence we dont have to do it if it only for that reason, and I'm still wondering until I learned about mocking.

Some other people says that interface is used so your program is decoupled, and when needed you can replace the implementations without needing to adjust the implementor. That make sense right? much better than the bullshit. Yea that make sense, we can replace whatever implement whatever interface with whatever. Yea, but how many times would you replace you database connection calls? chances are rare if not never especially if you working on software house that deliver projects after projects after projects, you will never see you component got replaced.

The when I learned about mocking, all that I have been asking coming to conclusions as if I was like having epiphany, we will discuss more about mocking in the mocking section, but for now lets discuss it in regards of dependency injection usage. So as you see in our project structure, instead of having all component directly talks to each other, we are using interface, take PlayerController for example

    type PlayerController struct {
      interfaces.IPlayerService
    }
    
    func (controller *PlayerController) GetPlayerScore(res http.ResponseWriter, req *http.Request) {
    
      player1Name := chi.URLParam(req, "player1")
      player2Name := chi.URLParam(req, "player2")
    
      scores, err := controller.GetScores(player1Name, player2Name)
      if err != nil {
        //Handle error
      }
    
	  json.NewEncoder(res).Encode(viewmodels.ScoresVM{scores})
    }

You see that PlayerController uses IPlayerService interface, and since IPlayerService has GetScores method, PlayerController can invoke it and get the result right away. Wait a minute, isn't that the interface is just merely abstraction? so how do it get executed, where is the implementation?

    type IPlayerService interface {
      GetScores(player1Name string, player2Name string) (string, error)
    }

You see, instead of calling directly to PlayerService, PlayerController uses the interface of PlayerService which is IPlayerService, there could be many implementation of IPlayerService not just limited to PlayerService it could be BrotherService etc, but how do we determined that PlayerService will be used instead?

    func (k *kernel) InjectPlayerController() controllers.PlayerController {

      sqlConn, _ := sql.Open("sqlite3", "/var/tmp/tennis.db")
      sqliteHandler := &infrastructures.SQLiteHandler{}
      sqliteHandler.Conn = sqlConn

	  playerRepository := &repositories.PlayerRepository{sqliteHandler}
	  playerService := &services.PlayerService{&repositories.PlayerRepositoryWithCircuitBreaker{playerRepository}}
	  playerController := controllers.PlayerController{playerService}

      return playerController
    }

This is where dependency injection come in to play, as you see here in servicecontainer.go we are creating **playerController** and inject it with **playerService** as simple as that, this is what dependency injection all about no more. So **playerController's IPlayerService** will be injected by **playerService** along with all implementation that it implements, so for example **GetPlayerByName** now returns whatever **GetPlayerByName** implemented by **playerService** as you can see it in **PlayerService.go**

Now, how does this relates to TDD & mocking?

	playerService := new(mocks.IPlayerService)

You see, in PlayerController_test.go we are using mock object to inject the implementation of our service, lets discuss more detail about mocking and testing in each section.

----------

[Mocking](https://irahardianto.github.io/service-pattern-go/#mocking)
-------

Mocking is a concept many times people struggle to understand, let alone implement it, at least I was the one among the one who struggles to understand this concept. But understanding this concept is essential to do TDD. The key point is, we mock dependencies that we need to run our tests, this is why dependency injection is essential to proceed. We are using testfy as our mock library

Basically what mock object do is replacing injection instead of real implementation with mock as point out at the end of dependency injection session

    playerService := new(mocks.IPlayerService)

We then create mock GetScores functionalities along with its request and response.

    playerService.On("GetScores", "Rafael", "Serena").Return("Forty-Fifteen", nil)

As you see, then the mock object is injected to **playerService** of PlayerController, this is why dependency injection is essential to this proses as it is the only way we can inject interface with mock object instead of real implementation.

	playerController := PlayerController{playerService}

We generate mock our by using vektra mockery for IPlayerService, go to the interfaces folder and then just type.

    mockery -name=IPlayerService

The output will be inside ```mocks/IPlayerService.go``` and we can use it right away for our testing.

----------

[Testing](https://irahardianto.github.io/service-pattern-go/#testing)
-------

We have cover pretty much everything there is I hope that you already get the idea of proper unit testing and why we should implement interfaces, dependency injection and mocking. The last piece is the unit test itself.

    func TestPlayerScore(t *testing.T) {

      // create an instance of our test object
      playerService := new(mocks.IPlayerService)

      // setup expectations
      playerService.On("GetScores", "Rafael", "Serena").Return("Forty-Fifteen", nil)

	  playerController := PlayerController{playerService}

      // call the code we are testing
      req := httptest.NewRequest("GET", "http://localhost:8080/getScore/Rafael/vs/Serena", nil)
      w := httptest.NewRecorder()

      r := chi.NewRouter()
      r.HandleFunc("/getScore/{player1}/vs/{player2}", playerController.GetPlayerScore)

      r.ServeHTTP(w, req)

      expectedResult := viewmodels.ScoresVM{}
      expectedResult.Score = "Forty-Fifteen"

      actualResult := viewmodels.ScoresVM{}

      json.NewDecoder(w.Body).Decode(&actualResult)

      // assert that the expectations were met
      assert.Equal(t, expectedResult, actualResult)
    }

 As you see here after injecting playerService of playerController with mock object, we are calling the playerController.GetPlayer and simulate request all the way from the router.

     req := httptest.NewRequest("GET", "http://localhost:8080/getScore/Rafael/vs/Serena", nil)
     w := httptest.NewRecorder()

     r := chi.NewRouter()
     r.HandleFunc("/getScore/{player1}/vs/{player2}", playerController.GetPlayerScore)

     r.ServeHTTP(w, req)

And assert the result by using testify assertion library

    assert.Equal(t, expectedResult, actualResult)

----------

[Circuit Breaker](https://irahardianto.github.io/service-pattern-go/#circuit-breaker)
-------

Building a distributed system we should really think that everything is not reliable, networks could breaks, servers could suddenly crash, even your 100% unit-tested app could be the root cause of the problems.

With that in said, when designing distributed system we should keep that in mind, so when some of our system is down, it won't take the whole system. Circuit breaker is a pattern with which we could design our system to be fault-tolerant and can withstand one or more service failure. It should be wrapping all call outside application ex: db call, redis call, api call.

Essentially circuit breaker works just like electrical circuit breakers, nothing fancy here, the only different is when the breaker is tripped it can be automatically closed when the downstream service is responding properly as described in the picture below.

![circuit breaker](https://cdn.pbrd.co/images/GKpFVb1.png)

In our case, we will be using hystrix-go, it is a go port from Netflix's hystrix library, how it works is essentially the same, even hystrix-go supports turbine along with its hystrix dashboard, but in my case, I rather use the datadog plugins, since we are using datadog to monitor our system.

For the sake of SOLID principles implementation in our codebase, we will add hystrix-go to our PlayerRepository leveraging decorator pattern, this will maintain our base repository implementation, the one that calls database, clean from modification and we will create its extension which is named PlayerRepositoryWithCircuitBreaker. This is the O part of SOLID which stands for Open for extension, Close for modification.


If you recall we inject our PlayerService with PlayerRepositoryWithCircuitBreaker and the original PlayerRepository wrapped inside.

	playerService.PlayerRepository = &repositories.PlayerRepositoryWithCircuitBreaker{playerRepository}


Base PlayerRepository implementation :

	type PlayerRepository struct {
      interfaces.IDbHandler
	}

    func (repository *PlayerRepository) GetPlayerByName(name string) (models.PlayerModel, error) {

      row, err :=repository.Query(fmt.Sprintf("SELECT * FROM player_models WHERE name = '%s'", name))
      if err != nil {
        return models.PlayerModel{}, err
      }

      var player models.PlayerModel

      row.Next()
      row.Scan(&player.Id, &player.Name, &player.Score)

      return player, nil
	}

PlayerRepository extension implementation :

    type PlayerRepositoryWithCircuitBreaker struct {
      PlayerRepository interfaces.IPlayerRepository
    }

    func (repository *PlayerRepositoryWithCircuitBreaker) GetPlayerByName(name string) (models.PlayerModel, error) {

      output := make(chan models.PlayerModel, 1)
      hystrix.ConfigureCommand("get_player_by_name", hystrix.CommandConfig{Timeout: 1000})
      errors := hystrix.Go("get_player_by_name", func() error {

        player, _ := repository.PlayerRepository.GetPlayerByName(name)

        output <- player
        return nil
      }, nil)

      select {
      case out := <-output:
        return out, nil
      case err := <-errors:
        println(err)
        return models.PlayerModel{}, err
      }
    }

Basically PlayerRepositoryWithCircuitBreaker implement the same interface as PlayerRepository, IPlayerRepository

    type IPlayerRepository interface {
      GetPlayerByName(name string) (models.PlayerModel, error)
    }


As you see here, it is very easy to implement hystrix-go circuit breaker, you just need to wrap your db call inside hystrix if the timeout reached, the circuit breaker will be tripped and all calls to database will be halt, error will be returned instead for future call until db service is up and healthy.


Cheers,
M. Ichsan Rahardianto.