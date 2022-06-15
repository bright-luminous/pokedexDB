package resource

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/bright-luminous/pokedexDB/graph/model"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

func PrintIfErrorExist(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func CheckPokeType(inStr string) (model.PokemonType, error) {
	var typeTobeReturn model.PokemonType
	var err error
	for _, v := range model.AllPokemonType {
		if string(v) == inStr {
			fmt.Println(v)
			typeTobeReturn = v
		}
	}
	if typeTobeReturn == "" {
		return "", errors.New("wrong type")
	}
	return typeTobeReturn, err
}

func NewPokemonSQLiteOperation(dbName string) (*PokemonSQLop, error) {
	sqldb, err := sql.Open(sqliteshim.ShimName, dbName)
	db := bun.NewDB(sqldb, sqlitedialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))
	return &PokemonSQLop{
		modelToUse: new(model.Pokemon),
		db:         db,
	}, err
}

func NewPokemonPostgresOperation(inDSN string) (*PokemonSQLop, error) {
	db, err := sql.Open("postgres", inDSN)
	bunDB := bun.NewDB(db, pgdialect.New())
	bunDB.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))
	return &PokemonSQLop{
		modelToUse: new(model.Pokemon),
		db:         bunDB,
	}, err
}

func (op *PokemonSQLop) CreateTable(ctx context.Context) (sql.Result, error) {
	sqlResult, err := op.db.NewCreateTable().Model((*model.Pokemon)(nil)).Exec(ctx)
	return sqlResult, err
}

func (op *PokemonSQLop) PokeCreate(ctx context.Context, pokeInput *model.PokemonCreateInput) (*model.Pokemon, error) {
	newID := uuid.New().String()
	pokemonTobeAdd := model.Pokemon{
		ID:          newID,
		Name:        pokeInput.Name,
		Description: pokeInput.Description,
		Category:    pokeInput.Category,
		Type:        pokeInput.Type,
		Abilities:   pokeInput.Abilities}
	_, err := op.db.NewInsert().Model(&pokemonTobeAdd).Exec(ctx)
	return &pokemonTobeAdd, err
}

func (op *PokemonSQLop) PokeUpdate(ctx context.Context, ID string, updateField model.FieldAvailable, updateVal string) ([]*model.Pokemon, error) {
	var err error
	if updateField == "Type" {
		newType, err := CheckPokeType(updateVal)
		if err != nil {
			return []*model.Pokemon(nil), err
		}
		_, err = op.db.NewUpdate().Model(op.modelToUse).Set("Type= ?", newType).Where("id = ?", ID).Exec(ctx)
		if err != nil {
			return []*model.Pokemon(nil), err
		}
	} else {
		tobeQuery := fmt.Sprintf("%s= '%s'", updateField, updateVal)
		fmt.Println(tobeQuery)
		_, err = op.db.NewUpdate().Model(op.modelToUse).Set(tobeQuery).Where("id = ?", ID).Exec(ctx)
		if err != nil {
			return []*model.Pokemon(nil), err
		}
	}
	resultPokemon, err := op.PokeFindByID(ctx, ID)
	return resultPokemon, err
}

func (op *PokemonSQLop) PokeUpdateAbility(ctx context.Context, ID string, newAbility []string) ([]*model.Pokemon, error) {
	_, err := op.db.NewUpdate().Model(op.modelToUse).Set("Abilities= ?", newAbility).Where("id = ?", ID).Exec(ctx)
	if err != nil {
		return []*model.Pokemon(nil), err
	}
	resultPokemon, err := op.PokeFindByID(ctx, ID)
	return resultPokemon, err
}

func (op *PokemonSQLop) PokeUpdateMulti(ctx context.Context, updateInput model.Pokemon) ([]*model.Pokemon, error) {
	_, err := op.db.NewUpdate().Model(&updateInput).Where("id = ?", updateInput.ID).Exec(ctx)
	resultPokemon, _ := op.PokeFindByID(ctx, updateInput.ID)
	return resultPokemon, err
}

func (op *PokemonSQLop) PokeDelete(ctx context.Context, ID string) ([]*model.Pokemon, error) {
	resultPokemon, err := op.PokeFindByID(ctx, ID)
	PrintIfErrorExist(err)
	_, err = op.db.NewDelete().Model(op.modelToUse).Where("id = ?", ID).Exec(ctx)
	return resultPokemon, err
}

func (op *PokemonSQLop) PokeList(ctx context.Context) ([]*model.Pokemon, error) {
	pokemons := new([]*model.Pokemon)
	err := op.db.NewSelect().Model(pokemons).Scan(ctx, pokemons)
	return *pokemons, err
}

func (op *PokemonSQLop) PokeFindByID(ctx context.Context, ID string) ([]*model.Pokemon, error) {
	arrModel := new([]*model.Pokemon)
	err := op.db.NewSelect().Model(op.modelToUse).Where("id = ?", ID).Scan(ctx, arrModel)
	return *arrModel, err
}

func (op *PokemonSQLop) PokeFindByName(ctx context.Context, Name string) ([]*model.Pokemon, error) {
	arrModel := new([]*model.Pokemon)
	err := op.db.NewSelect().Model(op.modelToUse).Where("Name = ?", Name).Scan(ctx, arrModel)
	PrintIfErrorExist(err)
	return *arrModel, err
}

func (op *PokemonSQLop) PokeDeleteAll(ctx context.Context) ([]*model.Pokemon, error) {
	pokeArr, err := op.PokeList(ctx)
	PrintIfErrorExist(err)
	for _, v := range pokeArr {
		_, err := op.PokeDelete(ctx, v.ID)
		PrintIfErrorExist(err)
	}
	return pokeArr, err
}
