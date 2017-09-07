package main

import (
	"katip/commands"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Katip"
	app.Usage = "Micro backups with historical changes"
	app.Version = "0.0.1"

	app.Commands = commands.GetAllCommands()

	app.Run(os.Args)
}
