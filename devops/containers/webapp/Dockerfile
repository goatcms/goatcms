FROM ubuntu
MAINTAINER Sebastian Pożoga <sebastian@pozoga.eu>

EXPOSE 5555

RUN apt-get update
RUN apt-get install -y wget curl git build-essential nodejs npm postgresql-client
RUN npm install --global n
RUN n stable
RUN npm install --global yarn
RUN npm --version
RUN yarn --version
RUN wget https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.8.linux-amd64.tar.gz

ENV GOROOT="/usr/local/go"
ENV GOPATH="/root/go"
ENV PATH="${PATH}:$GOROOT/bin:$GOPATH/bin"

RUN go get github.com/gorilla/mux
RUN go get github.com/jmoiron/sqlx
RUN go get github.com/goatcms/goatcore
RUN go get github.com/mattn/go-sqlite3
RUN go install github.com/mattn/go-sqlite3
RUN go get github.com/lib/pq
RUN go install github.com/lib/pq
RUN go get golang.org/x/crypto/bcrypt
RUN go get github.com/buger/jsonparser

ADD . /root/go/src/github.com/goatcms/goatcms
COPY devops/containers/webapp/config_dev.json /root/go/src/github.com/goatcms/goatcms/config/config_dev.json
COPY devops/containers/webapp/config_test.json /root/go/src/github.com/goatcms/goatcms/config/config_test.json
COPY devops/containers/webapp/config_prod.json /root/go/src/github.com/goatcms/goatcms/config/config_prod.json

RUN cd /root/go/src/github.com/goatcms/goatcms && go build /root/go/src/github.com/goatcms/goatcms/main.go
WORKDIR /root/go/src/github.com/goatcms/goatcms
