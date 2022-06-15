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
	go run main.go

dockerBuild:
	docker build -t pokedex-app:v0.1 .

dockerDexRun:
	docker run --name pokedex -p 8080:8080 -tid pokedex-app:v0.1