FROM golang:latest

RUN mkdir /build
WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download && go mod verify

RUN export GO111MODULE=on
# RUN go get github.com/bright-luminous/pokedexDB
RUN cd /build && git clone --branch postgresDB https://github.com/bright-luminous/pokedexDB.git

RUN cd /build/pokedexDB && go build

EXPOSE 8080

RUN cd /build/pokedexDB && ls

ENTRYPOINT [ "/build/pokedexDB/pokedexDB" ]