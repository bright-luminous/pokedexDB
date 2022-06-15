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
	"github.com/spf13/viper"
)

func main() {
	r := chi.NewRouter()

	viper.SetConfigName("pokemonConfig")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var host string = viper.GetString("connectionDetail.host")
	var port string = viper.GetString("connectionDetail.port")
	var user string = viper.GetString("connectionDetail.user")
	var password string = viper.GetString("connectionDetail.password")
	var dbname string = viper.GetString("connectionDetail.dbname")
	var goChiPort string = viper.GetString("connectionDetail.goChiPort")

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
