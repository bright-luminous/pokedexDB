package resource

import (
	"context"

	"github.com/bright-luminous/pokedexDB/graph/model"
	"github.com/uptrace/bun"
)

type PokemonSQLop struct {
	db         *bun.DB
	modelToUse *model.Pokemon
}

type DatabaseOp interface {
	PokeCreate(ctx context.Context, pokeInput *model.PokemonCreateInput) (*model.Pokemon, error)
	PokeUpdate(ctx context.Context, ID string, updateField model.FieldAvailable, updateVal string) ([]*model.Pokemon, error)
	PokeUpdateMulti(ctx context.Context, updateInput model.Pokemon) ([]*model.Pokemon, error)
	PokeDelete(ctx context.Context, ID string) ([]*model.Pokemon, error)
	PokeDeleteAll(ctx context.Context) ([]*model.Pokemon, error)

	PokeList(ctx context.Context) ([]*model.Pokemon, error)
	PokeFindByID(ctx context.Context, ID string) ([]*model.Pokemon, error)
	PokeFindByName(ctx context.Context, Name string) ([]*model.Pokemon, error)
}
