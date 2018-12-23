package commands

import (
	"fmt"
	"github.com/urfave/cli"
)

func GetAllCommands() []cli.Command {
	return []cli.Command{
		{
			Name:    "watch",
			Aliases: []string{"w"},
			Usage:   "Starts a new watch. By default, it commits changes to the history as any change occurs, and folder path is the current directory.",
			Action: func(c *cli.Context) error {
				success, err := runWatchCommand()
				if !success {
					fmt.Println(err)
				} else {
					fmt.Println("Watch has started... Go ahead and make some changes.")
				}
				return nil
			},
		},
		{
			Name:    "stop",
			Aliases: []string{"s"},
			Usage:   "Stops the watch.",
			Action: func(c *cli.Context) error {
				fmt.Println("Not yet implemented: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "destroy",
			Aliases: []string{"d"},
			Usage:   "Destroys all the history recorded. It also stops the application if it is still running.",
			Action: func(c *cli.Context) error {
				fmt.Println("Not yet implemented: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "pause",
			Aliases: []string{"p"},
			Usage:   "Pauses the app.",
			Action: func(c *cli.Context) error {
				fmt.Println("Not yet implemented: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "resume",
			Aliases: []string{"d"},
			Usage:   "Resumes the app with all the arguments passed initially when the app was first started.",
			Action: func(c *cli.Context) error {
				fmt.Println("Not yet implemented: ", c.Args().First())
				return nil
			},
		},
	}
}
