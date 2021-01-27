FROM golang:1.15-alpine

LABEL maintainer="Kittitach Piyamapornchai <bankgg@gmail.com>"

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

CMD ["./main"]
