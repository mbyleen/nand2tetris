package main

import (
	"errors"
	"fmt"
	"github.com/urfave/cli"
	"os"
	"strings"
)

func main() {
	app := cli.NewApp()
	app.Name = "gha"
	app.Usage = "Assemble HACK assembly language to machine code readable by the HACK computer"
	app.Action = func(c *cli.Context) error {

		filename := c.Args().Get(0)
		if strings.TrimSuffix(filename, ".asm") == filename {
			return errors.New("read: given file lacks .asm filetype")
		}

		input, err := read(filename)
		if err != nil {
			fmt.Println(err)
			return err
		}

		output := parse(input)

		if err := write(output, filename); err != nil {
			fmt.Println(err)
			return err
		}

		return nil
	}

	app.Run(os.Args)
}
