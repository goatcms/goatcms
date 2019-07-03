# Goatcms
[![Go Report Card](https://goreportcard.com/badge/github.com/goatcms/goatcms)](https://goreportcard.com/report/github.com/goatcms/goatcms)
[![GoDoc](https://godoc.org/github.com/goatcms/goatcms?status.svg)](https://godoc.org/github.com/goatcms/goatcms)

Golang CMS system. Project is still in development.

## Build
To build project use [goatcli](https://github.com/goatcms/goatcli)

Install goatcli
```
cd $GOPATH/src/github.com/goatcli
git colone https://github.com/goatcms/goatcli.git
cd goatcli
go install
```
Remember to set path environment to "$GOPATH/bin".

Run project build:
```
goatcli build
```

## Index
* [Intership guide - A quick start for trainees is here](docs/intership_guide.md)
* [Configuration](docs/configuration.md)

## Internship guide
[A quick start for trainees is here](docs/intership_guide.md)

## TODO
Create documentation

## Environment varibale
 - GOATCMS_TEST_MYSQL_USERNAME - is mysql database username
 - GOATCMS_TEST_MYSQL_PASSWORD - is mysql database password
 - GOATCMS_TEST_MYSQL_HOST - is mysql hostname (or ip)

## Run tests
1) Create MySQL / MariaDB database and set system environments variable
```
docker run --name test-mariadb  -p 3306:3306 -e MYSQL_ROOT_PASSWORD=pass123 -d mariadb
export GOATCMS_TEST_MYSQL_USERNAME=root
export GOATCMS_TEST_MYSQL_PASSWORD=pass123
export GOATCMS_TEST_MYSQL_HOST=localhost
```
2) Run tests
```
cd YOUR_PROJECT_DIRECTORY
go test ./...
```

## Authors
* Sebastian Pozoga <sebastian@pozoga.eu> - Founder
* s3c0nDD (https://github.com/s3c0nDD) - Internship
