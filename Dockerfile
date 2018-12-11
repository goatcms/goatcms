FROM golang:1.10
EXPOSE 80
USER 0

# make image
WORKDIR /go/src/github.com/goatcms/goatcms
COPY . .
RUN rm -rf config/*
RUN rm -rf data/*
RUN rm main.db ||:
RUN rm docker-compose.yaml ||:
RUN chmod +x /go/src/github.com/goatcms/goatcms/docker/secrets.sh
RUN sh ./docker/secrets.sh

RUN go get -u github.com/golang/dep/cmd/dep
RUN go get -u github.com/goatcms/goatcli

RUN goatcli build
RUN go build -o goatcms ./main.go
RUN chmod +x docker/entrypoint.sh
RUN mv /go/src/github.com/goatcms/goatcms /app

WORKDIR /app
RUN rm -rf .goat && \
  rm -rf cmsapp && \
  rm -rf vendor && \
  rm -rf modules && \
  rm -rf /app/data

RUN ln -s /data /app/data

ENTRYPOINT ["/app/docker/entrypoint.sh"]
CMD []
