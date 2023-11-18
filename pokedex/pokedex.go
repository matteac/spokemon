package pokedex

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	poke "github.com/matteac/spokemon/pokemon"
)

type Pokedex []poke.Pokemon

func (p Pokedex) Len() int {
  return len(p)
}

func (p Pokedex) Filter(pokemon poke.Pokemon) []poke.Pokemon{
  results := []poke.Pokemon{}
  for _, p := range p {
    if strings.Contains(strings.ToLower(p.Name), strings.ToLower(pokemon.Name)) &&
    p.HasTypes(pokemon.Types) && 
    p.Legendary == pokemon.Legendary&&
    (p.Generation == pokemon.Generation || pokemon.Generation == 0) {
      results = append(results, p)
    }
  }
  return results
}

func (p Pokedex) ReadCSV(path string) Pokedex {
  file, err := os.Open(path)
  if err != nil {
    panic(err)
  }
  defer file.Close()
  reader := csv.NewReader(file)
  records, err := reader.ReadAll()
  if err != nil {
    panic(err)
  }
  pokemons := []poke.Pokemon{}
  for _, record := range records {
    name := record[1]
    type1 := record[2]
    type2 := record[3]

    hp, err := strconv.Atoi(record[5])
    if err != nil {
      panic(err)
    }
    atk, err := strconv.Atoi(record[6])
    if err != nil {
      panic(err)
    }
    def, err := strconv.Atoi(record[7])
    if err != nil {
      panic(err)
    }
    spatk, err := strconv.Atoi(record[8])
    if err != nil {
      panic(err)
    }
    spdef, err := strconv.Atoi(record[9])
    if err != nil {
      panic(err)
    }
    spd, err := strconv.Atoi(record[10])
    if err != nil {
      panic(err)
    }
    gen, err := strconv.Atoi(record[11])
    if err != nil {
      panic(err)
    }
    legendary, err := strconv.ParseBool(record[12])
    if err != nil {
      panic(err)
    }
    pokemons = append(pokemons, poke.Pokemon{
      Name: name,
      Types: []poke.PokemonType{
        poke.PokemonType(type1),
        poke.PokemonType(type2),
      },
      HP: hp,
      Attack: atk,
      Defense: def,
      SpecialAttack: spatk,
      SpecialDefense: spdef,
      Speed: spd,
      Generation: gen,
      Legendary: legendary,
    })

  }
  fmt.Println("Pokedex populated with", len(pokemons), "pokemons")
  return pokemons
}

var Pokedex_ = Pokedex{}


var Pokemon = poke.Pokemon{
  Types: []poke.PokemonType{
    poke.Grass,
  },
}

var Results = []poke.Pokemon{}



