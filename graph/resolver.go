package graph

import "github.com/bright-luminous/pokedexDB/resource"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB *resource.PokemonSQLop
}
