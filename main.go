package main

import (
	"github.com/frodion/pnp/algorithms/other"
	"github.com/frodion/pnp/algorithms/search"
	"github.com/frodion/pnp/algorithms/sort"
	"github.com/urfave/cli"
	"log"
	"os"
)

var app = cli.NewApp()

func main() {
	info()
	commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func info() {
	app.Name = "PNP - Algorithms archive"
	app.Usage = "CLI for algorithm execution"
	app.Authors = []*cli.Author{
		{
			Name: "Rodion Fedotov",
		},
	}
	app.Version = "1.0.0"
}

func commands() {
	app.Commands = append(app.Commands, search.GetSearchCommands()...)
	app.Commands = append(app.Commands, sort.GetSortCommands()...)
	app.Commands = append(app.Commands, other.GetOtherCommands()...)
}
