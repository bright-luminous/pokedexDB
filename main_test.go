package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOpFunc(t *testing.T) {
	myPokemon := []pokemon{
		{
			ID:          1,
			Name:        "phum",
			Description: "look like pola bare",
			Category:    "infar",
			Type:        fairy,
			Abilities:   []string{"drink coffee"},
		},
		{
			ID:          2,
			Name:        "fone",
			Description: "have her own keyboard",
			Category:    "frontend",
			Type:        ghost,
			Abilities:   []string{"red screen"},
		},
		{
			ID:          3,
			Name:        "chic",
			Description: "have Ipad never use it",
			Category:    "backend",
			Type:        flying,
			Abilities:   []string{"driving"},
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	if cancel != nil {
		fmt.Printf("Context cancel msg : %v\n\n", cancel)
	}

	operator := new(pokemonSQLop)
	operator.init("sql.DB")

	//create table (not working)
	// queryResult, err := operator.createTable(ctx)
	// assert.Equal(t, 1, queryResult)
	// assert.Equal(t, nil, err)

	//create pokemon
	returnPokemon, err := operator.pokeCreate(ctx, myPokemon[0].Name, myPokemon[0].Description, myPokemon[0].Category, myPokemon[0].Type, myPokemon[0].Abilities)
	assert.Equal(t, myPokemon[0], returnPokemon)
	assert.Equal(t, nil, err)

	operator.pokeCreate(ctx, myPokemon[1].Name, myPokemon[1].Description, myPokemon[1].Category, myPokemon[1].Type, myPokemon[1].Abilities)
	operator.pokeCreate(ctx, myPokemon[2].Name, myPokemon[2].Description, myPokemon[2].Category, myPokemon[2].Type, myPokemon[2].Abilities)

	//update pokemon
	updatedPokemon, err := operator.pokeUpdate(ctx, 1, "Description", "maybe look like whale")
	afterUpdatePokemon := []pokemon{
		{
			ID:          1,
			Name:        "phum",
			Description: "maybe look like whale",
			Category:    "infar",
			Type:        fairy,
			Abilities:   []string{"drink coffee"},
		},
	}
	assert.Equal(t, afterUpdatePokemon, updatedPokemon)
	assert.Equal(t, nil, err)

	//delete pokemon
	deletedPokemon, err := operator.pokeDelete(ctx, 1)
	assert.Equal(t, []pokemon(nil), deletedPokemon)
	assert.Equal(t, nil, err)

	//list all pokemon
	listResult, err := operator.pokeList(ctx)
	assert.Equal(t, []pokemon{myPokemon[1], myPokemon[2]}, listResult)
	assert.Equal(t, nil, err)

	//list ID pokemon
	idResult, err := operator.pokeFindID(ctx, 2)
	assert.Equal(t, []pokemon{myPokemon[1]}, idResult)
	assert.Equal(t, nil, err)

	//list name pokemon
	nameResult, err := operator.pokeFindName(ctx, "chic")
	assert.Equal(t, []pokemon{myPokemon[2]}, nameResult)
	assert.Equal(t, nil, err)
}
