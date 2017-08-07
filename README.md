service-pattern-go
-------

Hey! welcome, this is and example of clean architecture written in Go Lang with complete Dependency Injection along with Mocking example.

Inspired by [Manuel Kiessling go-cleanarchitecture](http://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/) and [Joshua Partogi TDD training session](https://github.com/jpartogi/tennis-kata-laravel/)

It has simple dependencies:

 - [Gorilla Mux (Router)](https://github.com/gorilla/mux)
 - [Jinzhu GORM (ORM)](https://github.com/jinzhu/gorm)
 - [Testify (Test & Mock framework)](https://github.com/stretchr/testify)

Get Started:

 - [Install](https://irahardianto.github.io/service-pattern-go#install)
 - Introduction
 - Folder Structure
 - Naming Convention
 - Depency Injection
 - Mocking


----------


[Install](https://irahardianto.github.io/service-pattern-go#install)
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


 - Naming Convention
	 - Namespace
	 - Struct
	 - Interface
	 - Service
	 - Repository
	 - Controller



Cheers,
M. Ichsan Rahardianto.
