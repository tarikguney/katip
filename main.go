package main

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/urfave/cli"
)

// Bu uygulama bakalim hic basarili olacak mi? Shower thoughts olarak aklima geldi.
func main() {
	fmt.Println("Hello, World")
	app := cli.NewApp()
	app.Name = "Katip"
	app.Usage = "Micro backups with historical changes"
	command := exec.Command("git", "log")
	var out bytes.Buffer

	command.Stdout = &out
	var err error = command.Run()
	// app.Action = func(context *cli.Context) error {
	// 	fmt.Println("Hello, World")
	// 	return nil
	// }
	fmt.Println(out.String())
	fmt.Println(err)
	// app.Run(os.Args)
}
