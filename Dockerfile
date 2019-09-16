FROM golang:1.11.10

WORKDIR /app

ADD . /app

ENTRYPOINT ["go", "run", "server/*.go"]