service-pattern-go
-------

An example Go application demonstrating The Clean Architecture, Dependency Injection & Mocking.

Inspired by Manuel Kiessling go-cleanarchitecture and Joshua Partogi TDD training session.

It is written in Go using following library
- Jinzhu GORM (ORM)
- Testify (Test & Mock framework)

[http://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/](http://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/)

[https://github.com/jpartogi/tennis-kata-laravel/](https://github.com/jpartogi/tennis-kata-laravel/)


Install
-------

Clone the source

    git clone https://github.com/irahardianto/service-pattern-go

Setup dependencies

    go get -u github.com/jinzhu/gorm
    go get github.com/stretchr/testify

Setup sqlite data structure

    sqlite3 /var/tmp/gorm.db < setup.sql

Run the app, and visit

    http://localhost:8080/getPlayer?playerId=101

Cheers,

M. Ichsan Rahardianto.
