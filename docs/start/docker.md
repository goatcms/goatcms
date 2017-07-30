# Start with docker

## Important commands
* build image
```
cd $GOPATH/src/github.com/goatcms/goatcms
docker build -f "devops/containers/webapp/Dockerfile" -t goatwebapp .
```

* open bash
```
docker run -i -t -p 5555:5555 goatwebapp /bin/bash
```

* run goatcms webapp (in docker bash)
```
mkdir tmp
cd /root/go/src/github.com/goatcms/goatcms/
go test github.com/goatcms/goatcore/...
go test github.com/goatcms/goatcms/...
go run main.go run --loglvl=dev
```

* Open in your favorite web browser  http://[local_docker_ip]:5555

## Tips & Triks
- if your virtual machine disk is full, you can create a larger disk
```
docker-machine rm default
docker-machine create -d virtualbox --virtualbox-disk-size "100000" default
```
- if you haven't setup local developer environment, you can try [a virtual developer desktop](https://github.com/goatcms/developer-desktop) 
