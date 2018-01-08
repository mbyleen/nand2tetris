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
		} else {
			// Non-A instructions are C instructions
			b = parseCInstruction(line)
		}
		// remove this empty string check when working with symbols
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
	//strip the @ prefix
	s = s[1:]
	//see if the remainder will parse as a number
	n, err := strconv.ParseInt(s, 0, 15)
	if err != nil {
		return ""
	}
	//convert the number to binary
	return fmt.Sprintf("%015b", n)
}

// Parse C instruction using lookup tables
// 1 1 1 a c1 c2 c3 c4 c5 c6 d1 d2 d3 j1 j2 j3
// dest=comp;jump
func parseCInstruction(s string) string {
	ins := "111"
	var dest string
	var comp string
	var jump string

	first := strings.Split(s, "=")
	if len(first) > 1 {
		dest = getDest(first[0])
		ind := strings.Index(s, "=")
		s = s[ind+1:]
	} else {
		dest = getDest("null")
	}

	second := strings.Split(s, ";")
	comp = getComp(second[0])
	if len(second) > 1 {
		jump = getJump(second[1])
	} else {
		jump = getJump("null")
	}

	return ins + comp + dest + jump
}
