package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/bright-luminous/pokedexDB/graph"
	"github.com/bright-luminous/pokedexDB/graph/generated"
	"github.com/bright-luminous/pokedexDB/resource"
	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
)

func main() {
	r := chi.NewRouter()

	var host string = "postgreSQL_1"
	var port string = "5432"
	var user string = "postgres"
	var password string = os.Getenv("POSTGRES_PASSWORD")
	var dbname string = "postgres"
	var goChiPort string = "8080"

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	operator, err := resource.NewPokemonPostgresOperation(psqlInfo)
	resource.PrintIfErrorExist(err)

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					DB: operator,
				},
			},
		),
	)

	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	log.Printf("connect to http://%s:%s/ for GraphQL playground", host, goChiPort)
	log.Fatal(http.ListenAndServe(":"+goChiPort, r))
}
