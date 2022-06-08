package main

// import (
// 	"context"
// 	"fmt"
// 	"time"

// 	"github.com/bright-luminous/pokedexDB/resource"
// )

// var myPokemon []resource.Pokemon = []resource.Pokemon{
// 	{
// 		ID:          "1",
// 		Name:        "phum",
// 		Description: "look like pola bare",
// 		Category:    "infar",
// 		Type:        resource.Fairy,
// 		Abilities:   []string{"drink coffee"},
// 	},
// 	{
// 		ID:          "2",
// 		Name:        "fone",
// 		Description: "have her own keyboard",
// 		Category:    "frontend",
// 		Type:        resource.Ghost,
// 		Abilities:   []string{"red screen"},
// 	},
// 	{
// 		ID:          "3",
// 		Name:        "chic",
// 		Description: "have Ipad never use it",
// 		Category:    "backend",
// 		Type:        resource.Flying,
// 		Abilities:   []string{"driving"},
// 	},
// }

// func main() {
// 	operator := new(resource.PokemonSQLop)
// 	operator.Init("sql.DB")
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	if cancel != nil {
// 		fmt.Printf("Context cancel msg : %v\n\n", cancel)
// 	}
// 	// result, err := operator.PokeCreate(ctx, myPokemon[0].Name, myPokemon[0].Description, myPokemon[0].Category, myPokemon[0].Type, myPokemon[0].Abilities)
// 	// if cancel != nil {
// 	// 	fmt.Printf("Context cancel msg : %v\n\n", err)
// 	// }
// 	// fmt.Printf("Context cancel msg : %v\n\n", result)
// 	operator.PokeDeleteAll(ctx)
// }

// test push
