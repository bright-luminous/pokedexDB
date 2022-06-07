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

type pokemonSQLop struct {
	ctx        context.Context
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
		panic(err)
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
	ctx := context.Background()
	modelToUse := new(pokemon)
	op.db = db
	op.ctx = ctx
	op.modelToUse = modelToUse
}

func (op *pokemonSQLop) createTable() {
	_, err := op.db.NewCreateTable().Model((*pokemon)(nil)).Exec(op.ctx)
	errCheck(err)
}

func (op *pokemonSQLop) pokeCreate(name string, description string, category string, pokeType pokemonType, abilities []string) {
	pokemonTobeAdd := pokemon{Name: name, Description: description, Category: category, Type: pokeType, Abilities: abilities}
	_, err := op.db.NewInsert().Model(&pokemonTobeAdd).Exec(op.ctx)
	errCheck(err)
}

func (op *pokemonSQLop) pokeUpdate(name string, updateField string, updateVal string) {
	_, err := op.db.NewUpdate().Model(op.modelToUse).Set("?= ?", updateField, updateVal).Where("Name = ?", name).Exec(op.ctx)
	errCheck(err)
}

func (op *pokemonSQLop) pokeDelete(name string) {
	_, err := op.db.NewDelete().Model(op.modelToUse).Where("Name = ?", name).Exec(op.ctx)
	errCheck(err)
}

func (op *pokemonSQLop) pokeList() {
	arrModel := new([]pokemon)
	err := op.db.NewSelect().Model(arrModel).Scan(op.ctx, arrModel)
	errCheck(err)
	fmt.Printf("All Pokemon: %v\n\n", *arrModel)
}

func (op *pokemonSQLop) pokeFindID(ID int) {
	err := op.db.NewSelect().Model(op.modelToUse).Where("id = ?", ID).Scan(op.ctx, op.modelToUse)
	errCheck(err)
	fmt.Printf("Pokemon by ID: %v\n\n", *op.modelToUse)
}

func (op *pokemonSQLop) pokeFindName(Name string) {
	err := op.db.NewSelect().Model(op.modelToUse).Where("Name = ?", Name).Scan(op.ctx, op.modelToUse)
	errCheck(err)
	fmt.Printf("Pokemon by name: %v\n\n", *op.modelToUse)
}

func main() {
	operator := new(pokemonSQLop)
	operator.init("sql.DB")
	// operator.pokeCreate("Pikacu", "rat", "atk", electric, []string{"run", "lighting bolt"})
	// operator.pokeDelete("Pikacu")
	// operator.pokeUpdate("Pikacu", "description", "ha ha i update you")
	// operator.pokeList()
	// operator.pokeFindID(2)
	operator.pokeFindName("Pikacu")
}
