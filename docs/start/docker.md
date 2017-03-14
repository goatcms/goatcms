# Start with docker

* build image
```
cd $GOPATH/src/github.com/goatcms/goatcms
docker build -f "devops/developers/standalone/Dockerfile" -t goatcmsdev .
```

* open bash
```
docker run -i -t -p 5555:5555 goatcmsdev /bin/bash
```

* run goatcms
```
mkdir tmp
cd /root/go/src/github.com/goatcms/goatcms/
go test github.com/goatcms/goatcore/...
go test github.com/goatcms/goatcms/...
go run main.go run --loglvl=dev
```

* connect via http://[local_docker_ip]:5555
