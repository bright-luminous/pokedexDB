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
	pokeToReturn := resource.ResourceToModel([]resource.Pokemon{result})
	return pokeToReturn[0], err
}

func (r *mutationResolver) UpdatePokemon(ctx context.Context, input model.PokemonUpdateInput) ([]*model.Pokemon, error) {
	result, err := r.DB.PokeUpdate(ctx, input.ID, input.UpdateKey, input.UpdateVal)
	pokeToReturn := resource.ResourceToModel(result)
	return pokeToReturn, err
}

func (r *mutationResolver) UpdatePokemonMap(ctx context.Context, input model.PokemonMapUpdateInput) ([]*model.Pokemon, error) {
	result, err := r.DB.PokeUpdateMulti(ctx, input.ID, input.Name, input.Description, input.Category, resource.PokemonType(input.Type), input.Abilities)
	pokeToReturn := resource.ResourceToModel(result)
	return pokeToReturn, err
}

func (r *mutationResolver) DeletePokemon(ctx context.Context, input model.DeleteIDInput) ([]*model.Pokemon, error) {
	result, err := r.DB.PokeDelete(ctx, input.ID)
	pokeToReturn := resource.ResourceToModel(result)
	return pokeToReturn, err
}

func (r *mutationResolver) DeleteAllPokemon(ctx context.Context) ([]*model.Pokemon, error) {
	result, err := r.DB.PokeDeleteAll(ctx)
	pokeToReturn := resource.ResourceToModel(result)
	return pokeToReturn, err
}

func (r *queryResolver) ListAllPokemon(ctx context.Context) ([]*model.Pokemon, error) {
	result, err := r.DB.PokeList(ctx)
	pokeToReturn := resource.ResourceToModel(result)
	return pokeToReturn, err
}

func (r *queryResolver) QueryPokemonID(ctx context.Context, input string) ([]*model.Pokemon, error) {
	result, err := r.DB.PokeFindID(ctx, input)
	pokeToReturn := resource.ResourceToModel(result)
	return pokeToReturn, err
}

func (r *queryResolver) QueryPokemonName(ctx context.Context, input string) ([]*model.Pokemon, error) {
	result, err := r.DB.PokeFindName(ctx, input)
	pokeToReturn := resource.ResourceToModel(result)
	return pokeToReturn, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
