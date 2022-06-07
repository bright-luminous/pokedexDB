package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

type pokemonSQLop struct {
	db         *bun.DB
	modelToUse *pokemon
}

type pokemon struct {
	ID          int64 `bun:",pk,autoincrement"`
	Name        string
	Description string
	Category    string
	Type        pokemonType
	Abilities   []string
}

type pokemonType string

type pokemonUpdateInput struct {
	Name        string
	Description string
	Category    string
	Type        pokemonType
	Abilities   []string
}

const (
	bug      pokemonType = "bug"
	dark     pokemonType = "dark"
	dragon   pokemonType = "dragon"
	electric pokemonType = "lighting"
	fairy    pokemonType = "fairy"
	fighting pokemonType = "fighting"
	fire     pokemonType = "fire"
	flying   pokemonType = "flying"
	ghost    pokemonType = "ghost"
	grass    pokemonType = "grass"
	ground   pokemonType = "ground"
	ice      pokemonType = "ice"
	normal   pokemonType = "normal"
	poison   pokemonType = "poison"
	psychic  pokemonType = "psychic"
	rock     pokemonType = "rock"
	steel    pokemonType = "steel"
	water    pokemonType = "water"
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

func (op *pokemonSQLop) pokeUpdate(ctx context.Context, name string, updateField string, updateVal string) (pokemon, error) {
	_, err := op.db.NewUpdate().Model(op.modelToUse).Set("?= ?", updateField, updateVal).Where("Name = ?", name).Exec(ctx)
	return *op.modelToUse, err
}

func (op *pokemonSQLop) pokeDelete(ctx context.Context, name string) (pokemon, error) {
	_, err := op.db.NewDelete().Model(op.modelToUse).Where("Name = ?", name).Exec(ctx)
	return *op.modelToUse, err
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

func main() {
	operator := new(pokemonSQLop)
	operator.init("sql.DB")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	if cancel != nil {
		fmt.Printf("Context cancel msg : %v\n\n", cancel)
	}
	toPrint, err := operator.pokeFindName(ctx, "Chamander")
	errCheck(err)
	fmt.Printf("Pokemon by name: %v\n\n", toPrint)
}

// test push
