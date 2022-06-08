package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

func errCheck(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func (op *pokemonSQLop) init(dbName string) {
	sqldb, err := sql.Open(sqliteshim.ShimName, dbName)
	errCheck(err)
	db := bun.NewDB(sqldb, sqlitedialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))
	modelToUse := new(pokemon)
	op.db = db
	op.modelToUse = modelToUse
}

func (op *pokemonSQLop) createTable(ctx context.Context) (sql.Result, error) {
	sqlResult, err := op.db.NewCreateTable().Model((*pokemon)(nil)).Exec(ctx)
	return sqlResult, err
}

func (op *pokemonSQLop) pokeCreate(ctx context.Context, name string, description string, category string, pokeType pokemonType, abilities []string) (pokemon, error) {
	pokemonTobeAdd := pokemon{
		Name:        name,
		Description: description,
		Category:    category,
		Type:        pokeType,
		Abilities:   abilities}
	_, err := op.db.NewInsert().Model(&pokemonTobeAdd).Exec(ctx)
	return pokemonTobeAdd, err
}

func (op *pokemonSQLop) pokeUpdate(ctx context.Context, ID int, updateField string, updateVal string) ([]pokemon, error) {
	_, err := op.db.NewUpdate().Model(op.modelToUse).Set("?= ?", updateField, updateVal).Where("id = ?", ID).Exec(ctx)
	resultPokemon, err := op.pokeFindID(ctx, ID)
	return resultPokemon, err
}

func (op *pokemonSQLop) pokeDelete(ctx context.Context, ID int) ([]pokemon, error) {
	_, err := op.db.NewDelete().Model(op.modelToUse).Where("id = ?", ID).Exec(ctx)
	resultPokemon, err := op.pokeFindID(ctx, ID)
	return resultPokemon, err
}

func (op *pokemonSQLop) pokeList(ctx context.Context) ([]pokemon, error) {
	arrModel := new([]pokemon)
	err := op.db.NewSelect().Model(arrModel).Scan(ctx, arrModel)
	return *arrModel, err
}

func (op *pokemonSQLop) pokeFindID(ctx context.Context, ID int) ([]pokemon, error) {
	arrModel := new([]pokemon)
	err := op.db.NewSelect().Model(op.modelToUse).Where("id = ?", ID).Scan(ctx, arrModel)
	errCheck(err)
	return *arrModel, err
}

func (op *pokemonSQLop) pokeFindName(ctx context.Context, Name string) ([]pokemon, error) {
	arrModel := new([]pokemon)
	err := op.db.NewSelect().Model(op.modelToUse).Where("Name = ?", Name).Scan(ctx, arrModel)
	errCheck(err)
	return *arrModel, err
}

func (op *pokemonSQLop) pokeDeleteAll(ctx context.Context) {
	pokeArr, err := op.pokeList(ctx)
	errCheck(err)
	for _, v := range pokeArr {
		_, err := op.pokeDelete(ctx, int(v.ID))
		errCheck(err)
	}
}

func Add(a int, b int) int {
	return a + b
}
