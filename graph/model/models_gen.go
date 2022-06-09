// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type DeleteIDInput struct {
	ID string `json:"ID"`
}

type Pokemon struct {
	ID          string      `json:"ID"`
	Name        string      `json:"Name"`
	Description string      `json:"Description"`
	Category    string      `json:"Category"`
	Type        PokemonType `json:"Type"`
	Abilities   []string    `json:"Abilities"`
}

type PokemonCreateInput struct {
	ID          *string     `json:"ID"`
	Name        string      `json:"Name"`
	Description string      `json:"Description"`
	Category    string      `json:"Category"`
	Type        PokemonType `json:"Type"`
	Abilities   []string    `json:"Abilities"`
}

type PokemonMapUpdateInput struct {
	ID          string      `json:"ID"`
	Name        string      `json:"Name"`
	Description string      `json:"Description"`
	Category    string      `json:"Category"`
	Type        PokemonType `json:"Type"`
	Abilities   []string    `json:"Abilities"`
}

type PokemonUpdateInput struct {
	ID        string         `json:"ID"`
	UpdateKey FieldAvailable `json:"UpdateKey"`
	UpdateVal string         `json:"UpdateVal"`
}

type FieldAvailable string

const (
	FieldAvailableID          FieldAvailable = "ID"
	FieldAvailableName        FieldAvailable = "Name"
	FieldAvailableDescription FieldAvailable = "Description"
	FieldAvailableCategory    FieldAvailable = "Category"
	FieldAvailableType        FieldAvailable = "Type"
	FieldAvailableAbilities   FieldAvailable = "Abilities"
)

var AllFieldAvailable = []FieldAvailable{
	FieldAvailableID,
	FieldAvailableName,
	FieldAvailableDescription,
	FieldAvailableCategory,
	FieldAvailableType,
	FieldAvailableAbilities,
}

func (e FieldAvailable) IsValid() bool {
	switch e {
	case FieldAvailableID, FieldAvailableName, FieldAvailableDescription, FieldAvailableCategory, FieldAvailableType, FieldAvailableAbilities:
		return true
	}
	return false
}

func (e FieldAvailable) String() string {
	return string(e)
}

func (e *FieldAvailable) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FieldAvailable(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FieldAvailable", str)
	}
	return nil
}

func (e FieldAvailable) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PokemonType string

const (
	PokemonTypeBug      PokemonType = "Bug"
	PokemonTypeDark     PokemonType = "Dark"
	PokemonTypeDragon   PokemonType = "Dragon"
	PokemonTypeElectric PokemonType = "Electric"
	PokemonTypeFairy    PokemonType = "Fairy"
	PokemonTypeFighting PokemonType = "Fighting"
	PokemonTypeFire     PokemonType = "Fire"
	PokemonTypeFlying   PokemonType = "Flying"
	PokemonTypeGhost    PokemonType = "Ghost"
	PokemonTypeGrass    PokemonType = "Grass"
	PokemonTypeGround   PokemonType = "Ground"
	PokemonTypeIce      PokemonType = "Ice"
	PokemonTypeNormal   PokemonType = "Normal"
	PokemonTypePoison   PokemonType = "Poison"
	PokemonTypePsychic  PokemonType = "Psychic"
	PokemonTypeRock     PokemonType = "Rock"
	PokemonTypeSteel    PokemonType = "Steel"
	PokemonTypeWater    PokemonType = "Water"
)

var AllPokemonType = []PokemonType{
	PokemonTypeBug,
	PokemonTypeDark,
	PokemonTypeDragon,
	PokemonTypeElectric,
	PokemonTypeFairy,
	PokemonTypeFighting,
	PokemonTypeFire,
	PokemonTypeFlying,
	PokemonTypeGhost,
	PokemonTypeGrass,
	PokemonTypeGround,
	PokemonTypeIce,
	PokemonTypeNormal,
	PokemonTypePoison,
	PokemonTypePsychic,
	PokemonTypeRock,
	PokemonTypeSteel,
	PokemonTypeWater,
}

func (e PokemonType) IsValid() bool {
	switch e {
	case PokemonTypeBug, PokemonTypeDark, PokemonTypeDragon, PokemonTypeElectric, PokemonTypeFairy, PokemonTypeFighting, PokemonTypeFire, PokemonTypeFlying, PokemonTypeGhost, PokemonTypeGrass, PokemonTypeGround, PokemonTypeIce, PokemonTypeNormal, PokemonTypePoison, PokemonTypePsychic, PokemonTypeRock, PokemonTypeSteel, PokemonTypeWater:
		return true
	}
	return false
}

func (e PokemonType) String() string {
	return string(e)
}

func (e *PokemonType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PokemonType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid pokemonType", str)
	}
	return nil
}

func (e PokemonType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
