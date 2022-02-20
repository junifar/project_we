#Project We

## Pre-requisite
1. Postgresql installed
2. golang v 1.17.0++

## configuration
- all configuration there is in file ```/conf/app.conf```
- you can modify this especially for database config
- create database based on database name in ```/conf/app.conf```

## how to run
- clone all code from repository
- modify config file
- running go mod tidy to install all dependencies
- run command ```go run main.go``` from command line
- all table will be created default for first run