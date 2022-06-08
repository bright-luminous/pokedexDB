package resource

import "github.com/uptrace/bun"

type PokemonSQLop struct {
	db         *bun.DB
	modelToUse *Pokemon
}

type Pokemon struct {
	ID          string
	Name        string
	Description string
	Category    string
	Type        PokemonType
	Abilities   []string
}

type PokemonType string

const (
	Bug      PokemonType = "bug"
	Dark     PokemonType = "dark"
	Dragon   PokemonType = "dragon"
	Electric PokemonType = "lighting"
	Fairy    PokemonType = "fairy"
	Fighting PokemonType = "fighting"
	Fire     PokemonType = "fire"
	Flying   PokemonType = "flying"
	Ghost    PokemonType = "ghost"
	Grass    PokemonType = "grass"
	Ground   PokemonType = "ground"
	Ice      PokemonType = "ice"
	Normal   PokemonType = "normal"
	Poison   PokemonType = "poison"
	Psychic  PokemonType = "psychic"
	Rock     PokemonType = "rock"
	Steel    PokemonType = "steel"
	Water    PokemonType = "water"
)
