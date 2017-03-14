FROM ubuntu

EXPOSE 5555

MAINTAINER Sebastian Pożoga <sebastian@pozoga.eu>
RUN apt-get update
RUN apt-get install -y wget git build-essential
RUN wget https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.8.linux-amd64.tar.gz

ENV GOROOT="/usr/local/go"
ENV GOPATH="/root/go"
ENV PATH="${PATH}:$GOROOT/bin:$GOPATH/bin"

RUN git clone https://github.com/goatcms/goatcore.git /root/go/src/github.com/goatcms/goatcore
RUN go get github.com/gorilla/mux
RUN go get github.com/jmoiron/sqlx
RUN go get github.com/mattn/go-sqlite3
RUN go install github.com/mattn/go-sqlite3
RUN go get golang.org/x/crypto/bcrypt
RUN go get github.com/buger/jsonparser

COPY devops/developers/standalone/config_dev.json /root/go/src/github.com/goatcms/goatcms/config/config_dev.json
COPY devops/developers/standalone/config_test.json /root/go/src/github.com/goatcms/goatcms/config/config_test.json
COPY devops/developers/standalone/config_prod.json /root/go/src/github.com/goatcms/goatcms/config/config_prod.json

ADD . /root/go/src/github.com/goatcms/goatcms
RUN mkdir /root/go/src/github.com/goatcms/goatcms/tmp
