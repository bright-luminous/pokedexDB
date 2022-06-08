package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/bright-luminous/pokedexDB/graph"
	"github.com/bright-luminous/pokedexDB/graph/generated"
	"github.com/bright-luminous/pokedexDB/resource"
	"github.com/go-chi/chi"
)

const defaultPort = "8080"

func main() {
	r := chi.NewRouter()

	operator := new(resource.PokemonSQLop)
	operator.Init("sql.DB")

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

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, r))
}
