package main

import (
	"fmt"
	"strconv"
	"strings"
)

var symb = make(map[string]string)
var extra = make(map[string]int)

func parse(input []string) ([]string, error) {
	var binary []string
	for _, line := range input {
		line = removeComments(line)
		line = removeWhitespace(line)
		if line == "" {
			continue
		}
		// identify symbols to catalog
		if strings.HasPrefix(line, "(") && strings.HasSuffix(line, ")") {
			catalog(line, len(binary))
			continue
		}
		b := parseLine(line, len(binary))
		binary = append(binary, b)
	}
	for key := range extra {
		index := extra[key]
		val := symb[key]
		v, err := numToBinary(val)
		if err != nil {
			// how to handle this? can assume well-formed input for exercise
		}
		binary[index] = v
	}
	return binary, nil
}

func parseLine(line string, i int) string {
	var b string
	// identify A instructions
	if strings.HasPrefix(line, "@") {
		b = parseAInstruction(line, i)
	} else {
		// Non-A instructions are C instructions
		b = parseCInstruction(line)
	}
	return b
}

func removeComments(s string) string {
	ind := strings.Index(s, "//")
	if ind != -1 {
		s = s[:ind]
	}
	return s
}

func removeWhitespace(s string) string {
	return strings.TrimSpace(s)
}

// Store symbol values
func catalog(line string, i int) {
	// strip parentheses
	symbol := line[1 : len(line)-1]
	// Address referred to by (symbol) is that of next line
	symb[symbol] = strconv.Itoa(i)
}

// Parse an A instruction
func parseAInstruction(s string, i int) string {
	//strip the @ prefix
	s = s[1:]
	// Check for a number
	val, err := numToBinary(s)
	if err != nil {
		// Check if the symbol has a translation stored
		v, ok := symb[s]
		if ok == false {
			// store the symbol for later translation
			extra[s] = i
			return s
		}
		v, err := numToBinary(v)
		if err != nil {
			// ???
		}
		val = v
	}
	return val
}

func numToBinary(s string) (string, error) {
	n, err := strconv.ParseInt(s, 0, 16)
	if err != nil {
		return "", err
	}
	return "0" + fmt.Sprintf("%015b", n), nil
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
