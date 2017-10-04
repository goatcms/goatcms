
# Internship guide

## github
If you don't have a github, you have to [sing up](https://github.com/join?source=header-home)
Then please send me your username.

## Docker
Try [start with docker image](start/docker.md)

## For backend developers

### Install software
* Install gcc
  * Windows
    * Install [mingw64](http://mingw-w64.org/doku.php/download)
    * [Set system environments](https://superuser.com/questions/949560/how-do-i-set-system-environment-variables-in-windows-10) example:
      * Set C_INCLUDE_PATH
      ```
      C:\Program Files\mingw-w64\x86_64-6.3.0-posix-seh-rt_v5-rev1\mingw64\include
      C:\Program Files\mingw-w64\x86_64-6.3.0-posix-seh-rt_v5-rev1\mingw64\lib\gcc\mingw32\4.5.1\include
      ```
      * Set CPLUS_INCLUDE_PATH
      ```
      C:\Program Files\mingw-w64\x86_64-6.3.0-posix-seh-rt_v5-rev1\mingw64\include
      C:\Program Files\mingw-w64\x86_64-6.3.0-posix-seh-rt_v5-rev1\mingw64\lib\gcc\mingw32\4.5.1\include
      C:\Program Files\mingw-w64\x86_64-6.3.0-posix-seh-rt_v5-rev1\mingw64\lib\gcc\mingw32\4.5.1\include\c++
      C:\Program Files\mingw-w64\x86_64-6.3.0-posix-seh-rt_v5-rev1\mingw64\x86_64-w64-mingw32\include
      ```
      * Add to Path
      ```
      C:\Program Files\mingw-w64\x86_64-6.3.0-posix-seh-rt_v5-rev1\mingw64\bin
      C:\Program Files\mingw-w64\x86_64-6.3.0-posix-seh-rt_v5-rev1\mingw64\libexec\gcc\x86_64-w64-mingw32\6.2.0
      ```
  * Ubuntu
    ```
    sudo apt-get install build-essential
    ```
* Install golang
https://golang.org/doc/install
* Install git
  * Windows: https://desktop.github.com/
  * Ubuntu: `sudo apt-get install git`
* Install docker (https://docs.docker.com/engine/installation/)
* Install [atom](https://atom.io/) with [go-plus plugin](https://atom.io/packages/go-plus) or another IDE of choice.
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

### Get project
* Get project
```
go get github.com/goatcms/goatcore
go get github.com/goatcms/goatcms
```
* Install dependencies

### Run
```
cd github.com/goatcms/goatcms
# to show help
go run ./main.go
# to run server with developer log level
go run ./main.go run --loglvl=dev
```

## For frontend developers
Learn [react-redux-starter-kit](https://github.com/davezuko/react-redux-starter-kit)

### Install software for frontend developers
* Install [node and npm](https://nodejs.org/)
* Install [Yarn](https://yarnpkg.com/lang/en/docs/install/)

### Build frontend
```
npm run deploy
```
