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

func PrintIfErrorExist(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func CheckPokeType(inStr string) model.PokemonType {
	var typeTobeReturn model.PokemonType
	for _, v := range model.AllPokemonType {
		if string(v) == inStr {
			fmt.Println(v)
			typeTobeReturn = v
		}
	}
	if typeTobeReturn == "" {
		panic("wrong type")
	}
	return typeTobeReturn
}

func ReferencePokemon(pokemon []model.Pokemon) []*model.Pokemon {
	var pokeToReturn []*model.Pokemon
	for _, v := range pokemon {
		typeTobeUse := v.Type
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

func NewPokemonSQLOperation(dbName string) *PokemonSQLop {
	sqldb, err := sql.Open(sqliteshim.ShimName, dbName)
	PrintIfErrorExist(err)
	db := bun.NewDB(sqldb, sqlitedialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))
	return &PokemonSQLop{
		modelToUse: new(model.Pokemon),
		db:         db,
	}
}

func (op *PokemonSQLop) CreateTable(ctx context.Context) (sql.Result, error) {
	sqlResult, err := op.db.NewCreateTable().Model((*model.Pokemon)(nil)).Exec(ctx)
	return sqlResult, err
}

func (op *PokemonSQLop) PokeCreate(ctx context.Context, pokeInput model.PokemonCreateInput) (model.Pokemon, error) {
	newID := uuid.New().String()
	pokemonTobeAdd := model.Pokemon{
		ID:          newID,
		Name:        pokeInput.Name,
		Description: pokeInput.Description,
		Category:    pokeInput.Category,
		Type:        pokeInput.Type,
		Abilities:   pokeInput.Abilities}
	_, err := op.db.NewInsert().Model(&pokemonTobeAdd).Exec(ctx)
	return pokemonTobeAdd, err
}

func (op *PokemonSQLop) PokeUpdate(ctx context.Context, ID string, updateField model.FieldAvailable, updateVal string) ([]model.Pokemon, error) {
	var err error
	if updateField == "Type" {
		newType := CheckPokeType(updateVal)
		_, err = op.db.NewUpdate().Model(op.modelToUse).Set("Type= ?", newType).Where("id = ?", ID).Exec(ctx)
	} else {
		_, err = op.db.NewUpdate().Model(op.modelToUse).Set("?= ?", updateField, updateVal).Where("id = ?", ID).Exec(ctx)
	}
	PrintIfErrorExist(err)
	resultPokemon, err := op.PokeFindByID(ctx, ID)
	return resultPokemon, err
}

func (op *PokemonSQLop) PokeUpdateMulti(ctx context.Context, updateInput model.Pokemon) ([]model.Pokemon, error) {
	_, err := op.db.NewUpdate().Model(&updateInput).Where("id = ?", updateInput.ID).Exec(ctx)
	resultPokemon, _ := op.PokeFindByID(ctx, updateInput.ID)
	return resultPokemon, err
}

func (op *PokemonSQLop) PokeDelete(ctx context.Context, ID string) ([]model.Pokemon, error) {
	_, err := op.db.NewDelete().Model(op.modelToUse).Where("id = ?", ID).Exec(ctx)
	PrintIfErrorExist(err)
	resultPokemon, err := op.PokeFindByID(ctx, ID)
	return resultPokemon, err
}

func (op *PokemonSQLop) PokeList(ctx context.Context) ([]model.Pokemon, error) {
	pokemons := new([]model.Pokemon)
	err := op.db.NewSelect().Model(pokemons).Scan(ctx, pokemons)
	return *pokemons, err
}

func (op *PokemonSQLop) PokeFindByID(ctx context.Context, ID string) ([]model.Pokemon, error) {
	arrModel := new([]model.Pokemon)
	err := op.db.NewSelect().Model(op.modelToUse).Where("id = ?", ID).Scan(ctx, arrModel)
	PrintIfErrorExist(err)
	return *arrModel, err
}

func (op *PokemonSQLop) PokeFindByName(ctx context.Context, Name string) ([]model.Pokemon, error) {
	arrModel := new([]model.Pokemon)
	err := op.db.NewSelect().Model(op.modelToUse).Where("Name = ?", Name).Scan(ctx, arrModel)
	PrintIfErrorExist(err)
	return *arrModel, err
}

func (op *PokemonSQLop) PokeDeleteAll(ctx context.Context) ([]model.Pokemon, error) {
	pokeArr, err := op.PokeList(ctx)
	PrintIfErrorExist(err)
	for _, v := range pokeArr {
		_, err := op.PokeDelete(ctx, v.ID)
		PrintIfErrorExist(err)
	}
	return pokeArr, err
}
