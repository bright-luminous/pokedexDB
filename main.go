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

type pokemon struct {
	ID          int64 `bun:",pk,autoincrement"`
	Name        string
	Description string
	Category    string
	Type        pokemonType
	Abilities   []string
}

type pokemonType string

var pokemon1 = pokemon{Name: "chamander", Description: "fire lizrd", Category: "fire", Type: fire, Abilities: []string{"fire ball", "fly"}}
var pokemon2 = pokemon{Name: "Pikacu", Description: "lighting", Category: "lighting", Type: lighting, Abilities: []string{"run", "lighting bolt"}}

const (
	fire     pokemonType = "fire"
	lighting pokemonType = "lighting"
)

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func createTable(ctx context.Context, db *bun.DB, model *pokemon) {
	_, err := db.NewCreateTable().Model((*pokemon)(nil)).Exec(ctx)
	errCheck(err)
}

func pokeCreate(ctx context.Context, db *bun.DB, model *pokemon) {
	_, err := db.NewInsert().Model(model).Exec(ctx)
	errCheck(err)
}

func pokeUpdate(ctx context.Context, db *bun.DB, model *pokemon) {
	_, err := db.NewUpdate().Model(model).Where("Name LIKE 'Pikacu'").Exec(ctx)
	errCheck(err)
}

func pokeDelete(ctx context.Context, db *bun.DB, model *pokemon) {
	_, err := db.NewDelete().Model(model).Where("Name LIKE 'Pikacu'").Exec(ctx)
	errCheck(err)
}

func pokeList(ctx context.Context, db *bun.DB, model *[]pokemon) {
	err := db.NewSelect().Model(model).Scan(ctx, model)
	errCheck(err)
	fmt.Printf("All Pokemon: %v\n\n", *model)
}

func pokeFindID(ctx context.Context, db *bun.DB, model *pokemon, ID int) {
	err := db.NewSelect().Model(model).Where("id = ?", ID).Scan(ctx, model)
	errCheck(err)
	fmt.Printf("Pokemon by ID: %v\n\n", *model)
}

func pokeFindName(ctx context.Context, db *bun.DB, model *pokemon, Name string) {
	err := db.NewSelect().Model(model).Where("Name = ?", Name).Scan(ctx, model)
	errCheck(err)
	fmt.Printf("Pokemon by name: %v\n\n", *model)
}

func main() {
	sqldb, err := sql.Open(sqliteshim.ShimName, "sql.DB")
	errCheck(err)
	db := bun.NewDB(sqldb, sqlitedialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))
	ctx := context.Background()
	modelToUse := new(pokemon)

	// pokeCreate(ctx, db, &pokemon1)
	// pokeUpdate(ctx, db, &pokemon2)
	// pokeDelete(ctx, db, &pokemon1)
	// pokeList(ctx, db, modelToUse)
	// pokeFindID(ctx, db, modelToUse, 2)
	pokeFindName(ctx, db, modelToUse, "Pikacu")
}
