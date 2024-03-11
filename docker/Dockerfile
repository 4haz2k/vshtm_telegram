FROM golang:1.21

LABEL authors="github.com:4haz2k"

RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -mod vendor -o ./main.go

FROM scratch

COPY --from=build-app /app/main /opt/main

ENTRYPOINT ["/opt/main"]
