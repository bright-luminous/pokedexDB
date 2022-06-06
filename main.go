package main

import (
	"context"
	"database/sql"

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

var pokemon1 = pokemon{Name: "Pikacu", Description: "lighting rat", Category: "lighting", Type: fire, Abilities: []string{"run", "lighting bolt"}}
var pokemon2 = pokemon{Name: "Pikacu", Description: "lighting", Category: "lighting", Type: "animal", Abilities: []string{"run", "lighting bolt"}}

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

func pokeList(ctx context.Context, db *bun.DB, model *pokemon, query string) {
	rows, err := db.QueryContext(ctx, "SELECT * FROM pokemons")
	if err != nil {
		panic(err)
	}
	db.ScanRow(ctx, rows)

	// err = db.ScanRows(ctx, rows, &model)
}

func pokeFindID() {

}

func pokeFindName() {

}

func main() {
	sqldb, err := sql.Open(sqliteshim.ShimName, "sql.DB")
	if err != nil {
		panic(err)
	}
	db := bun.NewDB(sqldb, sqlitedialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))
	ctx := context.TODO()

	// pokeCreate(ctx, db, &pokemon1)
	// pokeUpdate(ctx, db, &pokemon2)
	// pokeDelete(ctx, db, &pokemon1)
	pokeList(ctx, db, &pokemon1, "")
}
