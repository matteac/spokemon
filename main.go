package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/matteac/spokemon/cli"
	"github.com/matteac/spokemon/pokedex"
	poke "github.com/matteac/spokemon/pokemon"
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
	if len(os.Args) > 1 {

		cli_app := cli.NewApp()
		if os.Args[1] == "--help" || os.Args[1] == "-h" {
			fmt.Println(cli_app.HelpMessage())
			return
		}
		cli_app.ParseArgs(os.Args[1:])
		types := []poke.PokemonType{}
		for _, t := range cli_app.Types {
			types = append(types, poke.PokemonType(t))
		}
		pokedex.Pokemon.Name = cli_app.Name
		pokedex.Pokemon.Types = types
		pokedex.Pokemon.Legendary = cli_app.Legendary
		pokedex.Pokemon.Generation = cli_app.Generation
		pokedex.Results = pokedex.Pokedex_.Filter(pokedex.Pokemon)
		for _, pokemon := range pokedex.Results {
			fmt.Println(pokemon.String())
		}
		return
	}

	app.SetRoot(grid, true)
	app.EnableMouse(true)
	app.SetFocus(ui.Form)
	app.SetInputCapture(capture)
	app.Run()

}
