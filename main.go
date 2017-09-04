package main

import (
	"fmt"
	"os"
	"strings"

	_ "github.com/urfave/cli"
)

func main() {
	fmt.Println("Hello, World")
	//cli.NewApp().Run(os.Args)
	fmt.Println(strings.Join(os.Args, ", "))
}
