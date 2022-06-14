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
	docker build -t pokedex-app .

dockerDexRun:
	docker run -p 8080:8080 -tid pokedex-app