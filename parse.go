package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parse(input []string) ([]string, error) {
	var binary []string
	for _, line := range input {
		var b string
		// identify A instructions
		if strings.HasPrefix(line, "@") {
			b = parseAInstruction(line)
		}
		if b != "" {
			binary = append(binary, b)
		}
	}
	return binary, nil
}

// Parse an A instruction
// currently returns an empty string for non-numerical instructions
// will need to change this when work with symbols later
func parseAInstruction(s string) string {
	//strip the prefix
	s = s[1:]
	//see if the remainder will parse as a number
	n, err := strconv.ParseInt(s, 0, 15)
	if err != nil {
		return ""
	}
	//convert the number to binary
	return fmt.Sprintf("%015b", n)
}
