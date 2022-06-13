package main

// import (
// 	"context"
// 	"fmt"
// 	"testing"
// 	"time"

// 	"github.com/bright-luminous/pokedexDB/graph/model"
// 	"github.com/bright-luminous/pokedexDB/resource"
// 	"github.com/stretchr/testify/assert"
// )

// // Should create a test function per operation.
// // If there are 5 operations, there should be at least 5 test functions for each operation.
// func TestOpCreateFunc(t *testing.T) {
// 	myPokemon := []model.PokemonCreateInput{
// 		{
// 			Name:        "phum",
// 			Description: "look like pola bare",
// 			Category:    "infar",
// 			Type:        model.PokemonTypeDragon,
// 			Abilities:   []string{"drink coffee"},
// 		},
// 		{
// 			Name:        "fone",
// 			Description: "have her own keyboard",
// 			Category:    "frontend",
// 			Type:        model.PokemonTypeGhost,
// 			Abilities:   []string{"red screen"},
// 		},
// 		{
// 			Name:        "chic",
// 			Description: "have Ipad never use it",
// 			Category:    "backend",
// 			Type:        model.PokemonTypeGrass,
// 			Abilities:   []string{"driving"},
// 		},
// 	}
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	if cancel != nil {
// 		fmt.Printf("Context cancel msg : %v\n\n", cancel)
// 	}
// 	operator := resource.NewPokemonSQLOperation("sql.DB")

// 	returnPokemon, err := operator.PokeCreate(ctx, myPokemon[0])
// 	assert.Equal(t, myPokemon[0], returnPokemon)
// 	assert.Equal(t, nil, err)
// }

// func TestOpUpdateFunc(t *testing.T) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	if cancel != nil {
// 		fmt.Printf("Context cancel msg : %v\n\n", cancel)
// 	}
// 	operator := resource.NewPokemonSQLOperation("sql.DB")

// 	updatedPokemon, err := operator.PokeUpdate(ctx, "1", "Description", "maybe look like whale")
// 	afterUpdatePokemon := []model.Pokemon{
// 		{
// 			ID:          "1",
// 			Name:        "phum",
// 			Description: "maybe look like whale",
// 			Category:    "infar",
// 			Type:        model.PokemonTypeBug,
// 			Abilities:   []string{"drink coffee"},
// 		},
// 	}
// 	assert.Equal(t, afterUpdatePokemon, updatedPokemon)
// 	assert.Equal(t, nil, err)

// 	listResult, err := operator.PokeList(ctx)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(listResult)
// }

// func TestOpUpdateMultiFunc(t *testing.T) {
// 	tobeUpdatePokemon := model.PokemonMapUpdateInput{
// 		ID:          "",
// 		Name:        "New",
// 		Description: "play genshin",
// 		Category:    "4th year now",
// 		Type:        model.PokemonTypeBug,
// 		Abilities:   []string{"clean his glass"},
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	if cancel != nil {
// 		fmt.Printf("Context cancel msg : %v\n\n", cancel)
// 	}
// 	operator := resource.NewPokemonSQLOperation("sql.DB")

// 	updateResult, err := operator.PokeUpdateMulti(ctx, model.Pokemon(tobeUpdatePokemon))
// 	assert.Equal(t, []model.Pokemon(nil), updateResult)
// 	assert.Equal(t, nil, err)

// }

// func TestOpDeleteFunc(t *testing.T) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	if cancel != nil {
// 		fmt.Printf("Context cancel msg : %v\n\n", cancel)
// 	}
// 	operator := resource.NewPokemonSQLOperation("sql.DB")

// 	deletedPokemon, err := operator.PokeDelete(ctx, "1")
// 	assert.Equal(t, []model.Pokemon(nil), deletedPokemon)
// 	assert.Equal(t, nil, err)

// }

// func TestOpDeleteAll(t *testing.T) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	if cancel != nil {
// 		fmt.Printf("Context cancel msg : %v\n\n", cancel)
// 	}
// 	operator := resource.NewPokemonSQLOperation("sql.DB")

// 	result, err := operator.PokeDeleteAll(ctx)
// 	assert.Equal(t, []model.Pokemon(nil), result)
// 	assert.Equal(t, nil, err)

// }

// func TestOpListAll(t *testing.T) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	if cancel != nil {
// 		fmt.Printf("Context cancel msg : %v\n\n", cancel)
// 	}
// 	operator := resource.NewPokemonSQLOperation("sql.DB")
// 	listResult, err := operator.PokeList(ctx)
// 	assert.Equal(t, []model.Pokemon{}, listResult)
// 	assert.Equal(t, nil, err)

// }

// func TestOpListId(t *testing.T) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	if cancel != nil {
// 		fmt.Printf("Context cancel msg : %v\n\n", cancel)
// 	}
// 	operator := resource.NewPokemonSQLOperation("sql.DB")

// 	idResult, err := operator.PokeFindByID(ctx, "2")
// 	assert.Equal(t, []model.Pokemon{}, idResult)
// 	assert.Equal(t, nil, err)

// }

// func TestOpListName(t *testing.T) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	if cancel != nil {
// 		fmt.Printf("Context cancel msg : %v\n\n", cancel)
// 	}
// 	operator := resource.NewPokemonSQLOperation("sql.DB")

// 	nameResult, err := operator.PokeFindByName(ctx, "chic")
// 	assert.Equal(t, []model.Pokemon{}, nameResult)
// 	assert.Equal(t, nil, err)

// }
