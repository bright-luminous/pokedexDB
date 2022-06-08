run:
	go run *.go

main:
	go run main.go

test:
	go run main_test.go

gen:
	go get github.com/99designs/gqlgen
	go run github.com/99designs/gqlgen generate

server:
	go run server.go