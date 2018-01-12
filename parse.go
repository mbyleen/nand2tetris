package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parse(input []string) ([]string, error) {

	var labels = make(map[string]string)
	var extra = make(map[string][]int)
	var extraKeys = make([]string, 0)
	var varAddr = 16

	var binary []string
	for _, line := range input {

		line = removeComments(line)
		line = removeWhitespace(line)
		if line == "" {
			continue
		}

		// identify symbols to catalog
		if strings.HasPrefix(line, "(") && strings.HasSuffix(line, ")") {
			labels = catalog(line, len(binary), labels)
			continue
		}

		// parse A and C instructions
		var b string
		// identify A instructions
		if strings.HasPrefix(line, "@") {
			extra, extraKeys, b = parseAInstruction(line, len(binary), labels, extra, extraKeys)
		} else {
			// Non-A instructions are C instructions
			b = parseCInstruction(line)
		}
		binary = append(binary, b)
	}

	for _, key := range extraKeys {
		indices := extra[key]
		// Check for a label symbol
		val, ok := labels[key]
		if ok == true {
			v, err := numToBinary(val)
			if err != nil {
				// how to handle this? can assume well-formed input for exercise
			}
			for _, index := range indices {
				binary[index] = v
			}
			continue
		}

		// Assign a variable space
		addr := strconv.Itoa(varAddr)
		v, err := numToBinary(addr)
		if err != nil {
			// ??
		}
		for _, index := range indices {
			binary[index] = v
		}
		varAddr++
	}
	return binary, nil
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
func catalog(line string, i int, labels map[string]string) map[string]string {
	// strip parentheses
	symbol := line[1 : len(line)-1]
	// Address referred to by (symbol) is that of next line
	labels[symbol] = strconv.Itoa(i)
	return labels
}

// Parse an A instruction
func parseAInstruction(s string, i int, labels map[string]string, extra map[string][]int, extraKeys []string) (map[string][]int, []string, string) {
	//strip the @ prefix
	s = s[1:]

	// Check for a numerical value
	v, err := numToBinary(s)
	if err == nil {
		return extra, extraKeys, v
	}

	// Check if the symbol is predefined
	v = getSymbol(s)
	if v != "" {
		val, err := numToBinary(v)
		if err != nil {
			// ???
		}
		return extra, extraKeys, val
	}

	// Check if the symbol has a translation stored (is a label symbol)
	v, ok := labels[s]
	if ok == true {
		val, err := numToBinary(v)
		if err != nil {
			// ???
		}
		return extra, extraKeys, val
	}

	// store the symbol for later translation
	if extra[s] == nil {
		extra[s] = make([]int, 0)
	}
	extra[s] = append(extra[s], i)
	if !contains(extraKeys, s) {
		extraKeys = append(extraKeys, s)
	}
	return extra, extraKeys, s
}

// TODO make this take an int instead of a string and convert only
// numbers read from the input file
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

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
