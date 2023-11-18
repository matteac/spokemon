package cli

import (
	"strconv"
	"strings"
)

// declare args for the cli
type App struct {
  Name string
  Types []string
  Generation int
  Legendary bool
}

func NewApp() *App {
  return &App{
    Name: "",
    Types: []string{},
    Generation: 0,
    Legendary: false,
  }
}

func (a *App) ParseArgs(args []string) {
  for i, arg := range args {
    switch arg {
    case "-n", "--name":
      a.Name = args[i+1]
    case "-t", "--types":
      a.Types = strings.Split(args[i+1], ",")
    case "-g", "--generation":
      a.Generation, _ = strconv.Atoi(args[i+1])
    case "-l", "--legendary":
      a.Legendary, _ = strconv.ParseBool(args[i+1])
    }
  }
}

func (a *App) HelpMessage() string {
  str := ""
  str += "Usage: spokemon [options]\n"
  str += "Options:\n"
  str += "  -n, --name\t\t<STRING>\tName of the pokemon\n"
  str += "  -t, --types\t\t<STRING>\tTypes of the pokemon, separated by commas\n"
  str += "  -g, --generation\t<INT>\t\tGeneration of the pokemon\n"
  str += "  -l, --legendary\t<BOOL>\t\tIs the pokemon legendary?\n"
  str += "  -h, --help\t\t\t\tPrint this help message\n"
  return str
}
