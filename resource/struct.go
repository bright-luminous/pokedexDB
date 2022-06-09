package resource

import (
	"github.com/bright-luminous/pokedexDB/graph/model"
	"github.com/uptrace/bun"
)

type PokemonSQLop struct {
	db         *bun.DB
	modelToUse *model.Pokemon
}
