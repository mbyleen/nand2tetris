package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "ghcc"
	app.Usage = "Cross-compile HACK assembly language to machine code readable by the HACK computer"
	app.Action = func(c *cli.Context) error {
		filename := c.Args().Get(0)
		input, err := read(filename)
		if err != nil {
			fmt.Println(err)
			return err
		}
		output, err := parse(input)
		if err != nil {
			fmt.Println(err)
			return err
		}
		if err := write(output, filename); err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	}

	app.Run(os.Args)
}
