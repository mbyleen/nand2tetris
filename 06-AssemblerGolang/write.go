package main

import (
	"os"
	"strings"
)

func write(output []string, filename string) error {
	name := strings.TrimSuffix(filename, ".asm")
	filename = name + ".hack"

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, line := range output {
		if _, err := f.WriteString(line + "\n"); err != nil {
			return err
		}
	}
	return nil
}
