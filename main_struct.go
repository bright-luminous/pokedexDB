package main

import "github.com/uptrace/bun"

type pokemonSQLop struct {
	db         *bun.DB
	modelToUse *pokemon
}

type pokemon struct {
	ID          int64 `bun:",pk,autoincrement"`
	Name        string
	Description string
	Category    string
	Type        pokemonType
	Abilities   []string
}

type pokemonType string

// type pokemonUpdateInput struct {
// 	Name        string
// 	Description string
// 	Category    string
// 	Type        pokemonType
// 	Abilities   []string
// }

const (
	bug      pokemonType = "bug"
	dark     pokemonType = "dark"
	dragon   pokemonType = "dragon"
	electric pokemonType = "lighting"
	fairy    pokemonType = "fairy"
	fighting pokemonType = "fighting"
	fire     pokemonType = "fire"
	flying   pokemonType = "flying"
	ghost    pokemonType = "ghost"
	grass    pokemonType = "grass"
	ground   pokemonType = "ground"
	ice      pokemonType = "ice"
	normal   pokemonType = "normal"
	poison   pokemonType = "poison"
	psychic  pokemonType = "psychic"
	rock     pokemonType = "rock"
	steel    pokemonType = "steel"
	water    pokemonType = "water"
)
