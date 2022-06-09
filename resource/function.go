package resource

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/bright-luminous/pokedexDB/graph/model"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

func ErrCheck(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func typeConvert(inType model.PokemonType) model.PokemonType {
	tobeCon := inType
	tobeReturn := model.PokemonType(tobeCon)

	return tobeReturn
}

func ReferencePokemon(result []model.Pokemon) []*model.Pokemon {
	var pokeToReturn []*model.Pokemon
	for _, v := range result {
		typeTobeUse := typeConvert(v.Type)
		pokeToAppend := model.Pokemon{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Category:    v.Category,
			Type:        typeTobeUse,
			Abilities:   v.Abilities,
		}
		pokeToReturn = append(pokeToReturn, &pokeToAppend)
	}
	return pokeToReturn
}

func (op *PokemonSQLop) Init(dbName string) {
	sqldb, err := sql.Open(sqliteshim.ShimName, dbName)
	ErrCheck(err)
	db := bun.NewDB(sqldb, sqlitedialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))
	modelToUse := new(model.Pokemon)
	op.db = db
	op.modelToUse = modelToUse
}

func (op *PokemonSQLop) CreateTable(ctx context.Context) (sql.Result, error) {
	sqlResult, err := op.db.NewCreateTable().Model((*model.Pokemon)(nil)).Exec(ctx)
	return sqlResult, err
}

func (op *PokemonSQLop) PokeCreate(ctx context.Context, name string, description string, category string, pokeType model.PokemonType, abilities []string) (model.Pokemon, error) {
	newID := uuid.New().String()
	pokemonTobeAdd := model.Pokemon{
		ID:          newID,
		Name:        name,
		Description: description,
		Category:    category,
		Type:        pokeType,
		Abilities:   abilities}
	queryResult, err := op.db.NewInsert().Model(&pokemonTobeAdd).Exec(ctx)
	fmt.Println(queryResult)
	return pokemonTobeAdd, err
}

func (op *PokemonSQLop) PokeUpdate(ctx context.Context, ID string, updateField string, updateVal string) ([]model.Pokemon, error) {
	var err error
	if updateField == "Type" {
		_, err = op.db.NewUpdate().Model(op.modelToUse).Set("Type= ?", updateVal).Where("id = ?", ID).Exec(ctx)
	} else {
		_, err = op.db.NewUpdate().Model(op.modelToUse).Set("?= ?", updateField, updateVal).Where("id = ?", ID).Exec(ctx)
	}
	resultPokemon, _ := op.PokeFindID(ctx, ID)
	return resultPokemon, err
}

func (op *PokemonSQLop) PokeUpdateMulti(ctx context.Context, ID string, updateName string, updateDescription string, updateCategory string, updateType model.PokemonType, updateAbilities []string) ([]model.Pokemon, error) {
	_, err := op.db.NewUpdate().
		Model(op.modelToUse).
		Set("ID= ?", ID).
		Set("Name= ?", updateName).
		Set("Description= ?", updateDescription).
		Set("Category= ?", updateCategory).
		Set("Type= ?", updateType).
		Set("Abilities= ?", updateAbilities).
		Where("id = ?", ID).Exec(ctx)
	resultPokemon, _ := op.PokeFindID(ctx, ID)
	return resultPokemon, err
}

func (op *PokemonSQLop) PokeDelete(ctx context.Context, ID string) ([]model.Pokemon, error) {
	op.db.NewDelete().Model(op.modelToUse).Where("id = ?", ID).Exec(ctx)
	resultPokemon, err := op.PokeFindID(ctx, ID)
	return resultPokemon, err
}

func (op *PokemonSQLop) PokeList(ctx context.Context) ([]model.Pokemon, error) {
	arrModel := new([]model.Pokemon)
	err := op.db.NewSelect().Model(arrModel).Scan(ctx, arrModel)
	return *arrModel, err
}

func (op *PokemonSQLop) PokeFindID(ctx context.Context, ID string) ([]model.Pokemon, error) {
	arrModel := new([]model.Pokemon)
	err := op.db.NewSelect().Model(op.modelToUse).Where("id = ?", ID).Scan(ctx, arrModel)
	ErrCheck(err)
	return *arrModel, err
}

func (op *PokemonSQLop) PokeFindName(ctx context.Context, Name string) ([]model.Pokemon, error) {
	arrModel := new([]model.Pokemon)
	err := op.db.NewSelect().Model(op.modelToUse).Where("Name = ?", Name).Scan(ctx, arrModel)
	ErrCheck(err)
	return *arrModel, err
}

func (op *PokemonSQLop) PokeDeleteAll(ctx context.Context) ([]model.Pokemon, error) {
	pokeArr, err := op.PokeList(ctx)
	ErrCheck(err)
	for _, v := range pokeArr {
		_, err := op.PokeDelete(ctx, v.ID)
		ErrCheck(err)
	}
	return pokeArr, err
}
