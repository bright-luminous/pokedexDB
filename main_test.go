package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/bright-luminous/pokedexDB/graph/model"
	"github.com/bright-luminous/pokedexDB/resource"
	"github.com/stretchr/testify/assert"
)

func TestOpFunc(t *testing.T) {
	myPokemon := []model.Pokemon{
		{
			ID:          "1",
			Name:        "phum",
			Description: "look like pola bare",
			Category:    "infar",
			Type:        model.PokemonTypeDragon,
			Abilities:   []string{"drink coffee"},
		},
		{
			ID:          "2",
			Name:        "fone",
			Description: "have her own keyboard",
			Category:    "frontend",
			Type:        model.PokemonTypeGhost,
			Abilities:   []string{"red screen"},
		},
		{
			ID:          "3",
			Name:        "chic",
			Description: "have Ipad never use it",
			Category:    "backend",
			Type:        model.PokemonTypeGrass,
			Abilities:   []string{"driving"},
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	if cancel != nil {
		fmt.Printf("Context cancel msg : %v\n\n", cancel)
	}

	operator := new(resource.PokemonSQLop)
	operator.Init("sql.DB")

	//create table (not working)
	// queryResult, err := operator.createTable(ctx)
	// assert.Equal(t, 1, queryResult)
	// assert.Equal(t, nil, err)

	//create pokemon
	returnPokemon, err := operator.PokeCreate(ctx, myPokemon[0].Name, myPokemon[0].Description, myPokemon[0].Category, myPokemon[0].Type, myPokemon[0].Abilities)
	assert.Equal(t, myPokemon[0], returnPokemon)
	assert.Equal(t, nil, err)

	operator.PokeCreate(ctx, myPokemon[1].Name, myPokemon[1].Description, myPokemon[1].Category, myPokemon[1].Type, myPokemon[1].Abilities)
	operator.PokeCreate(ctx, myPokemon[2].Name, myPokemon[2].Description, myPokemon[2].Category, myPokemon[2].Type, myPokemon[2].Abilities)

	//update pokemon
	updatedPokemon, err := operator.PokeUpdate(ctx, "1", "Description", "maybe look like whale")
	afterUpdatePokemon := []model.Pokemon{
		{
			ID:          "1",
			Name:        "phum",
			Description: "maybe look like whale",
			Category:    "infar",
			Type:        model.PokemonTypeBug,
			Abilities:   []string{"drink coffee"},
		},
	}
	assert.Equal(t, afterUpdatePokemon, updatedPokemon)
	assert.Equal(t, nil, err)

	listResult, err := operator.PokeList(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(listResult)

	//delete pokemon
	deletedPokemon, err := operator.PokeDelete(ctx, "1")
	assert.Equal(t, []model.Pokemon(nil), deletedPokemon)
	assert.Equal(t, nil, err)

	//list all pokemon
	listResult, err = operator.PokeList(ctx)
	assert.Equal(t, []model.Pokemon{myPokemon[1], myPokemon[2]}, listResult)
	assert.Equal(t, nil, err)

	//list ID pokemon
	idResult, err := operator.PokeFindID(ctx, "2")
	assert.Equal(t, []model.Pokemon{myPokemon[1]}, idResult)
	assert.Equal(t, nil, err)

	//list name pokemon
	nameResult, err := operator.PokeFindName(ctx, "chic")
	assert.Equal(t, []model.Pokemon{myPokemon[2]}, nameResult)
	assert.Equal(t, nil, err)

	//clear the DB
	// operator.PokeDeleteAll(ctx)
}
