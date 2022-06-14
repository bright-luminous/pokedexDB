package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/bright-luminous/pokedexDB/graph"
	"github.com/bright-luminous/pokedexDB/graph/generated"
	"github.com/bright-luminous/pokedexDB/resource"
	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
)

const (
	host      string = "localhost"
	port      string = "5432"
	user      string = "postgres"
	password  string = "Eauu0244"
	dbname    string = "postgres"
	goChiPort string = "8080"
)

func main() {
	r := chi.NewRouter()
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	operator, err := resource.NewPokemonPostgresOperation(psqlInfo)
	// operator.CreateTable(context.Background())
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

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", goChiPort)
	log.Fatal(http.ListenAndServe(":"+goChiPort, r))
}

// func main() {
// 	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
// 		"password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	err = db.Ping()
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Successfully connected!")
// }
