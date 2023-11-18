package Pokemon

import (
	"strconv"
	"strings"
)

type PokemonType string

const (
	Grass PokemonType = "Grass"
	Fire  PokemonType = "Fire"
	Water PokemonType = "Water"
	Bug   PokemonType = "Bug"

	Normal   PokemonType = "Normal"
	Poison   PokemonType = "Poison"
	Electric PokemonType = "Electric"
	Ground   PokemonType = "Ground"

	Fairy    PokemonType = "Fairy"
	Fighting PokemonType = "Fighting"
	Psychic  PokemonType = "Psychic"
	Rock     PokemonType = "Rock"

	Steel  PokemonType = "Steel"
	Ice    PokemonType = "Ice"
	Ghost  PokemonType = "Ghost"
	Dragon PokemonType = "Dragon"

	Flying PokemonType = "Flying"
	Dark   PokemonType = "Dark"
)

func (t PokemonType) String() string {
	return string(t)
}

type Pokemon struct {
	Name           string
	Types          []PokemonType
	HP             int
	Attack         int
	Defense        int
	SpecialAttack  int
	SpecialDefense int
	Speed          int
	Generation     int
	Legendary      bool
}

func (p *Pokemon) LegendaryString() string {
	return strings.ToUpper(strconv.FormatBool(p.Legendary))
}

func (p *Pokemon) HasTypes(t []PokemonType) bool {
  if len(t) == 0 {
    return true
  }
	for _, type_ := range t {
		for _, pokemonType := range p.Types {
			if pokemonType == type_ {
				return true
			}
		}
	}
	return false
}
