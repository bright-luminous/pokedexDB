package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/bright-luminous/pokedexDB/graph/generated"
	"github.com/bright-luminous/pokedexDB/graph/model"
)

func (r *mutationResolver) CreateTable(ctx context.Context) (*model.Pokemon, error) {
	var nilReturn model.Pokemon
	_, err := r.DB.CreateTable(ctx)
	return &nilReturn, err
}

func (r *mutationResolver) CreatePokemon(ctx context.Context, input model.PokemonCreateInput) (*model.Pokemon, error) {
	if input.ID != nil {
		return nil, fmt.Errorf("id must be null")
	}
	result, err := r.DB.PokeCreate(ctx, &input)
	return result, err
}

func (r *mutationResolver) UpdatePokemon(ctx context.Context, input model.PokemonUpdateInput) ([]*model.Pokemon, error) {
	result, err := r.DB.PokeUpdate(ctx, input.ID, input.UpdateKey, input.UpdateVal)
	return result, err
}

func (r *mutationResolver) UpdatePokemonAbility(ctx context.Context, input model.PokemonUpdateAbilityInput) ([]*model.Pokemon, error) {
	result, err := r.DB.PokeUpdateAbility(ctx, input.ID, input.NewAbility)
	return result, err
}

func (r *mutationResolver) UpdatePokemonMap(ctx context.Context, input model.PokemonMapUpdateInput) ([]*model.Pokemon, error) {
	result, err := r.DB.PokeUpdateMulti(ctx, model.Pokemon(input))
	return result, err
}

func (r *mutationResolver) DeletePokemon(ctx context.Context, input model.DeleteIDInput) ([]*model.Pokemon, error) {
	result, err := r.DB.PokeDelete(ctx, input.ID)
	return result, err
}

func (r *mutationResolver) DeleteAllPokemon(ctx context.Context) ([]*model.Pokemon, error) {
	result, err := r.DB.PokeDeleteAll(ctx)
	return result, err
}

func (r *queryResolver) ListAllPokemon(ctx context.Context) ([]*model.Pokemon, error) {
	result, err := r.DB.PokeList(ctx)
	return result, err
}

func (r *queryResolver) QueryPokemonID(ctx context.Context, input string) ([]*model.Pokemon, error) {
	result, err := r.DB.PokeFindByID(ctx, input)
	return result, err
}

func (r *queryResolver) QueryPokemonName(ctx context.Context, input string) ([]*model.Pokemon, error) {
	result, err := r.DB.PokeFindByName(ctx, input)
	return result, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
