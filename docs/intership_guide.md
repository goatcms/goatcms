
# Internship guide

## github
If you don't have a github, you must [sing up](https://github.com/join?source=header-home)
Next send me your username.

## Install software
* Install golang
https://golang.org/doc/install
* Install git
  * For wirndows: https://desktop.github.com/
  * For ubuntu: `sudo apt-get install git`
* Install docker (https://docs.docker.com/engine/installation/)
* Install [atom](https://atom.io/) with [go-plus plugin](https://atom.io/packages/go-plus)
* Installing Missing Tools
```
go get -u golang.org/x/tools/cmd/goimports
go get -u golang.org/x/tools/cmd/gorename
go get -u github.com/sqs/goreturns
go get -u github.com/nsf/gocode
go get -u github.com/alecthomas/gometalinter
go get -u github.com/zmb3/gogetdoc
go get -u github.com/rogpeppe/godef
go get -u golang.org/x/tools/cmd/guru
```
* Get project
```
go get github.com/goatcms/goatcore
go get github.com/goatcms/goatcms
```
* Install dependencies
* Run
```
cd github.com/goatcms/goatcms
# to show help
go run ./main.go
# to run server with developer log level
go run ./main.go run --loglvl=dev
```
