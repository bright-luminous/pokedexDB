FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/bright-luminous/pokedexDB/postgresDB
RUN cd /build && git clone --branch postgresDB https://github.com/bright-luminous/pokedexDB.git

RUN cd /build && go build

EXPOSE 8080

ENTRYPOINT [ "/build" ]