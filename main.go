package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/matteac/spokemon/pokedex"
	"github.com/matteac/spokemon/ui"
	"github.com/rivo/tview"
)

var app *tview.Application = tview.NewApplication()
// on the left the Pokedex.PokemonForm and on the right the Pokedex.PokemonList
var grid *tview.Grid = tview.NewGrid().
  SetSize(1, 3, 0, 0).
  AddItem(ui.Form, 0, 0, 1, 1, 0, 0, false).
  AddItem(ui.Table, 0, 1, 1, 2, 2, 2, false)

func capture(event *tcell.EventKey) *tcell.EventKey {
	if event.Rune() == 'q' {
		app.Stop()
	} 
	return event
}

func main() {
  pokedex.Pokedex_ = pokedex.Pokedex_.ReadCSV("pokemon.csv")

  app.SetRoot(grid, true)
  app.EnableMouse(true)
  app.SetFocus(ui.Form)
	app.SetInputCapture(capture)
	app.Run()

}
