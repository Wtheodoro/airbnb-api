# todo for production use: run with less privilege, add build stage
FROM golang:alpine3.14

ARG PORT

WORKDIR /opt/app

COPY main.go go.mod go.sum ./

RUN go mod download

CMD ["go", "run", "main.go"]

