package ui

import (
	"strconv"
  "github.com/gdamore/tcell/v2"
	pokedex "github.com/matteac/spokemon/pokedex"
	poke "github.com/matteac/spokemon/pokemon"
	"github.com/rivo/tview"
)

var Table *tview.Table = tview.NewTable().
  SetBorders(true).
  SetSelectable(true, false).
  SetCell(0, 0, tview.NewTableCell("Name")).
  SetCell(0, 1, tview.NewTableCell("Type")).
  SetCell(0, 2, tview.NewTableCell("HP")).
  SetCell(0, 3, tview.NewTableCell("ATK")).
  SetCell(0, 4, tview.NewTableCell("DEF")).
  SetCell(0, 5, tview.NewTableCell("SP.ATK")).
  SetCell(0, 6, tview.NewTableCell("SP.DEF")).
  SetCell(0, 7, tview.NewTableCell("SPD")).
  SetCell(0, 8, tview.NewTableCell("Gen")).
  SetCell(0, 9, tview.NewTableCell("Legendary")).
  SetFixed(1, 0, )


var Form = tview.NewForm().
  AddInputField("Name", "", 20, nil, func(name string) {
    pokedex.Pokemon.Name = name
  }).
  AddDropDown("Type", []string{"None","Grass", "Fire", "Water", "Bug", "Normal", "Poison", "Electric", "Ground", "Fairy", "Fighting", "Psychic", "Rock", "Steel", "Ice", "Ghost", "Dragon", "Flying", "Dark"}, 0, func(option string, _ int) {
    if option == "None" {
      pokedex.Pokemon.Types = []poke.PokemonType{}
    } else {
      pokedex.Pokemon.Types = []poke.PokemonType{poke.PokemonType(option)}
    }
  }).
  AddDropDown("Generation", []string{"None","1", "2", "3", "4", "5", "6"}, 0, func(option string, _ int) {
    if option == "None" {
      pokedex.Pokemon.Generation = 0
    }
    pokedex.Pokemon.Generation, _ = strconv.Atoi(option)
  }).
  AddDropDown("Legendary", []string{"No", "Yes"}, 0, func(option string, _ int) {
    if option == "Yes" {
      pokedex.Pokemon.Legendary = true
    } else {
      pokedex.Pokemon.Legendary = false

    }
  }).
  AddButton("Search", func() {
    Table.Clear()
    Table.
  SetCell(0, 0, tview.NewTableCell("Name")).
  SetCell(0, 1, tview.NewTableCell("Type")).
  SetCell(0, 2, tview.NewTableCell("HP")).
  SetCell(0, 3, tview.NewTableCell("ATK")).
  SetCell(0, 4, tview.NewTableCell("DEF")).
  SetCell(0, 5, tview.NewTableCell("SP.ATK")).
  SetCell(0, 6, tview.NewTableCell("SP.DEF")).
  SetCell(0, 7, tview.NewTableCell("SPD")).
  SetCell(0, 8, tview.NewTableCell("Gen")).
  SetCell(0, 9, tview.NewTableCell("Legendary"))
    pokedex.Results = pokedex.Pokedex_.Filter(pokedex.Pokemon)
    for i, pokemon := range pokedex.Results {
      i := i + 1
      Table.
        SetCell(i, 0, tview.NewTableCell(pokemon.Name))
      Table.
        SetCell(i, 1, tview.NewTableCell(pokemon.Types[0].String()))
      Table.
        SetCell(i, 2, tview.NewTableCell(strconv.Itoa(pokemon.HP)))
      Table.
        SetCell(i, 3, tview.NewTableCell(strconv.Itoa(pokemon.Attack)))
      Table.
        SetCell(i, 4, tview.NewTableCell(strconv.Itoa(pokemon.Defense)))
      Table.
        SetCell(i, 5, tview.NewTableCell(strconv.Itoa(pokemon.SpecialAttack)))
      Table.
        SetCell(i, 6, tview.NewTableCell(strconv.Itoa(pokemon.SpecialDefense)))
      Table.
        SetCell(i, 7, tview.NewTableCell(strconv.Itoa(pokemon.Speed)))
      Table.
        SetCell(i, 8, tview.NewTableCell(strconv.Itoa(pokemon.Generation)))
      Table.
        SetCell(i, 9, tview.NewTableCell(pokemon.LegendaryString()))
    }
  }).
  SetFieldTextColor(tcell.ColorWhite).
  SetFieldBackgroundColor(tcell.ColorDarkMagenta).
  SetButtonBackgroundColor(tcell.ColorDarkMagenta)


