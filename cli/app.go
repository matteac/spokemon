package cli

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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
      if len(args) > i+1 {
        a.Name = args[i+1]
      } else {
        fmt.Println("Missing argument for --name")
        os.Exit(1)
      }
    case "-t", "--types":
      if len(args) > i+1 {
        a.Types = append(a.Types, cases.Title(language.English).String(strings.ToLower(args[i+1])))
      } else {
        fmt.Println("Missing argument for --type")
        os.Exit(1)
      }
    case "-g", "--generation":
      if len(args) > i+1 {
        a.Generation, _ = strconv.Atoi(args[i+1])
      } else {
        fmt.Println("Missing argument for --generation")
        os.Exit(1)
      }
    case "-l", "--legendary":
      if len(args) > i+1 {
        a.Legendary, _ = strconv.ParseBool(args[i+1])
      } else {
        a.Legendary = true
      }
    }
  }
}

func (a *App) HelpMessage() string {
  str := ""
  str += "Usage: spokemon [options]\n"
  str += "Options:\n"
  str += "  -n, --name\t\t<STRING>\tName of the pokemon\n"
  str += "  -t, --type\t\t<STRING>\tType of the pokemon\n"
  str += "  -g, --generation\t<INT>\t\tGeneration of the pokemon\n"
  str += "  -l, --legendary\t<BOOL>\t\tIs the pokemon legendary?\n"
  str += "  -h, --help\t\t\t\tPrint this help message\n"
  return str
}
