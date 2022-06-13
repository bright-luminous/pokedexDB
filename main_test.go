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

var currentExpectPokemon model.Pokemon = model.Pokemon{
	ID:          "",
	Name:        "fone",
	Description: "look like pola bare",
	Category:    "infar",
	Type:        model.PokemonTypeDragon,
	Abilities:   []string{"drink coffee"},
}

var inputPokemon = model.PokemonCreateInput{
	Name:        currentExpectPokemon.Name,
	Description: currentExpectPokemon.Description,
	Category:    currentExpectPokemon.Category,
	Type:        currentExpectPokemon.Type,
	Abilities:   currentExpectPokemon.Abilities,
}

// Should create a test function per operation.
// If there are 5 operations, there should be at least 5 test functions for each operation.
func testSetup(dbName string) (context.Context, *resource.PokemonSQLop) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	if cancel != nil {
		fmt.Printf("Context cancel msg : %v\n\n", cancel)
	}
	operator := resource.NewPokemonSQLOperation("sql.DB")
	return ctx, operator
}

func createAdditional(ctx context.Context, op resource.PokemonSQLop, input model.PokemonCreateInput) (*model.Pokemon, error) {
	returnPokemon, err := op.PokeCreate(ctx, &input)
	return returnPokemon, err
}

func TestOpCreateFunc(t *testing.T) {
	ctx, operator := testSetup("sql.DB")

	returnPokemon, err := operator.PokeCreate(ctx, &inputPokemon)
	currentExpectPokemon.ID = returnPokemon.ID
	assert.Equal(t, &currentExpectPokemon, returnPokemon)
	assert.Equal(t, nil, err)
	operator.PokeDeleteAll(ctx)

}

func TestOpUpdateFunc(t *testing.T) {
	ctx, operator := testSetup("sql.DB")
	returnPokemon, err := createAdditional(ctx, *operator, inputPokemon)
	assert.Equal(t, nil, err)
	currentExpectPokemon.ID = returnPokemon.ID

	updateVar := "maybe look like whale"
	updatedPokemon, err := operator.PokeUpdate(ctx, currentExpectPokemon.ID, model.FieldAvailableDescription, updateVar)
	currentExpectPokemon.Description = updateVar
	assert.Equal(t, currentExpectPokemon, *updatedPokemon[0])
	assert.Equal(t, nil, err)
	operator.PokeDeleteAll(ctx)
}

func TestOpUpdateMultiFunc(t *testing.T) {
	ctx, operator := testSetup("sql.DB")
	returnPokemon, err := createAdditional(ctx, *operator, inputPokemon)
	resource.PrintIfErrorExist(err)
	tobeUpdatePokemon := model.Pokemon{
		ID:          returnPokemon.ID,
		Name:        "New",
		Description: "play genshin",
		Category:    "4th year now",
		Type:        model.PokemonTypeBug,
		Abilities:   []string{"clean his glass"},
	}
	currentExpectPokemon = tobeUpdatePokemon

	updateResult, err := operator.PokeUpdateMulti(ctx, tobeUpdatePokemon)
	assert.Equal(t, currentExpectPokemon, *updateResult[0])
	assert.Equal(t, nil, err)
	operator.PokeDeleteAll(ctx)

}

func TestOpDeleteFunc(t *testing.T) {
	ctx, operator := testSetup("sql.DB")
	returnPokemon, err := createAdditional(ctx, *operator, inputPokemon)
	resource.PrintIfErrorExist(err)

	deletedPokemon, err := operator.PokeDelete(ctx, returnPokemon.ID)
	assert.Equal(t, *returnPokemon, *deletedPokemon[0])
	assert.Equal(t, nil, err)
	operator.PokeDeleteAll(ctx)

}

func TestOpDeleteAll(t *testing.T) {
	ctx, operator := testSetup("sql.DB")
	returnPokemon, err := createAdditional(ctx, *operator, inputPokemon)
	resource.PrintIfErrorExist(err)

	result, err := operator.PokeDeleteAll(ctx)
	assert.Equal(t, returnPokemon, result[0])
	assert.Equal(t, nil, err)
	operator.PokeDeleteAll(ctx)

}

func TestOpListAll(t *testing.T) {
	ctx, operator := testSetup("sql.DB")
	returnPokemon, err := createAdditional(ctx, *operator, inputPokemon)
	resource.PrintIfErrorExist(err)

	listResult, err := operator.PokeList(ctx)
	assert.Equal(t, []*model.Pokemon{returnPokemon}, listResult)
	assert.Equal(t, nil, err)
	operator.PokeDeleteAll(ctx)

}

func TestOpListId(t *testing.T) {
	ctx, operator := testSetup("sql.DB")
	returnPokemon, err := createAdditional(ctx, *operator, inputPokemon)
	resource.PrintIfErrorExist(err)

	idResult, err := operator.PokeFindByID(ctx, returnPokemon.ID)
	assert.Equal(t, []*model.Pokemon{returnPokemon}, idResult)
	assert.Equal(t, nil, err)
	operator.PokeDeleteAll(ctx)

}

func TestOpListName(t *testing.T) {
	ctx, operator := testSetup("sql.DB")
	returnPokemon, err := createAdditional(ctx, *operator, inputPokemon)
	resource.PrintIfErrorExist(err)

	nameResult, err := operator.PokeFindByName(ctx, returnPokemon.Name)
	assert.Equal(t, []*model.Pokemon{returnPokemon}, nameResult)
	assert.Equal(t, nil, err)
	operator.PokeDeleteAll(ctx)

}
