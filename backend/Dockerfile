FROM golang:1.18.1-bullseye

RUN apt-get update
RUN apt-get install -y --no-install-recommends \
    g++ \
    gcc \
    libc6-dev \
    make \
    pkg-config

RUN mkdir /go/src/app
WORKDIR /go/src/app
RUN go get -u github.com/cosmtrek/air
RUN go get -u github.com/go-delve/delve/cmd/dlv
RUN go mod tidy

CMD ["air", "-c", ".air.toml"]**