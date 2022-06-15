FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
COPY . /build
RUN go mod download && go mod verify

RUN cd /build && go build

EXPOSE 8080


ENTRYPOINT [ "/build/pokedexDB" ]