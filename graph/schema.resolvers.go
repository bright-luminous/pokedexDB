package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/bright-luminous/pokedexDB/graph/generated"
	"github.com/bright-luminous/pokedexDB/graph/model"
	"github.com/bright-luminous/pokedexDB/resource"
)

func (r *mutationResolver) CreatePokemon(ctx context.Context, input model.PokemonCreateInput) (*model.Pokemon, error) {
	if input.ID != nil {
		return nil, fmt.Errorf("id must be null")
	}
	result, err := r.DB.PokeCreate(ctx, input.Name, input.Description, input.Category, resource.PokemonType(input.Type), input.Abilities)
	pokeToReturn := model.Pokemon{
		ID:          result.ID,
		Name:        result.Name,
		Description: result.Description,
		Category:    result.Category,
		Type:        string(result.Type),
		Abilities:   result.Abilities,
	}
	return &pokeToReturn, err
}

func (r *mutationResolver) UpdatePokemon(ctx context.Context, input model.PokemonUpdateInput) ([]*model.Pokemon, error) {
	result, err := r.DB.PokeUpdate(ctx, input.ID, input.UpdateKey, input.UpdateVal)
	var pokeToReturn []*model.Pokemon
	for _, v := range result {
		pokeToAppend := model.Pokemon{
			Name:        v.Name,
			Description: v.Description,
			Category:    v.Category,
			Type:        string(v.Type),
			Abilities:   v.Abilities,
		}
		pokeToReturn = append(pokeToReturn, &pokeToAppend)
	}
	return pokeToReturn, err
}

func (r *mutationResolver) DeletePokemon(ctx context.Context, input model.DeleteIDInput) ([]*model.Pokemon, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteAllPokemon(ctx context.Context) ([]*model.Pokemon, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListAllPokemon(ctx context.Context) ([]*model.Pokemon, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) QueryPokemonID(ctx context.Context, input string) ([]*model.Pokemon, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) QueryPokemonName(ctx context.Context, input string) ([]*model.Pokemon, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
