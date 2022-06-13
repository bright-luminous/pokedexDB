package main

import (
	"database/sql"
	"fmt"

	"github.com/bright-luminous/pokedexDB/resource"
	_ "github.com/lib/pq"
)

const (
	host     string = "localhost"
	port     string = "5432"
	user     string = "postgres"
	password string = "Eauu0244"
	dbName   string = "postgres"
)

// func main() {
// 	r := chi.NewRouter()

// 	operator, err := resource.NewPokemonSQLOperation("sql.DB")
// 	resource.PrintIfErrorExist(err)

// 	srv := handler.NewDefaultServer(
// 		generated.NewExecutableSchema(
// 			generated.Config{
// 				Resolvers: &graph.Resolver{
// 					DB: operator,
// 				},
// 			},
// 		),
// 	)

// 	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
// 	r.Handle("/query", srv)

// 	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
// 	log.Fatal(http.ListenAndServe(":"+defaultPort, r))
// }

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	resource.PrintIfErrorExist(err)
	defer db.Close()

	err = db.Ping()
	resource.PrintIfErrorExist(err)

	fmt.Println("Successfully connected")
}
