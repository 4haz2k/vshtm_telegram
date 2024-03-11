FROM golang:latest AS build-app

RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -mod vendor -o ./main ./main.go

VOLUME /app/

FROM scratch

COPY --from=build-app /app/main /opt/main
COPY --from=build-app /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/opt/main"]

