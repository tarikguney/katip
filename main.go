package main

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/urfave/cli"
)

// I think the reason why this pkg folder exists is because it contains pre-compiled object files
// that are not going to be compiled again for speed during the whole compilation process.
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
