package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Parser struct {
	labels    map[string]string
	extra     map[string][]int
	extraKeys []string
	varAddr   int
}

func NewParser() *Parser {
	return &Parser{
		labels:    make(map[string]string),
		extra:     make(map[string][]int),
		extraKeys: make([]string, 0),
		varAddr:   16,
	}
}

func parse(input []string) ([]string, error) {

	p := NewParser()

	var binary []string
	for _, line := range input {

		// remove comments and whitespace
		line = delNonCode(line)
		if line == "" {
			continue
		}

		// identify labels for later translation
		if strings.HasPrefix(line, "(") && strings.HasSuffix(line, ")") {
			p.catalog(line, len(binary))
			continue
		}

		// parse A and C instructions
		var b string
		if strings.HasPrefix(line, "@") {
			b = p.parseAInstruction(line, len(binary))
		} else {
			b = parseCInstruction(line)
		}

		binary = append(binary, b)
	}

	// translate reserved variables and labels
	binary = p.fillInVar(binary)

	return binary, nil
}

func (p *Parser) fillInVar(binary []string) []string {
	for _, key := range p.extraKeys {
		indices := p.extra[key]
		// Check for a label symbol
		val, ok := p.labels[key]
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
		addr := strconv.Itoa(p.varAddr)
		v, err := numToBinary(addr)
		if err != nil {
			// ??
		}
		for _, index := range indices {
			binary[index] = v
		}
		p.varAddr++
	}
	return binary
}

// Store symbol values
func (p *Parser) catalog(line string, i int) {
	// strip parentheses
	symbol := line[1 : len(line)-1]
	// Address referred to by (symbol) is that of next line
	p.labels[symbol] = strconv.Itoa(i)
}

// Parse an A instruction
func (p *Parser) parseAInstruction(s string, i int) string {
	//strip the @ prefix
	s = s[1:]

	// Check for a numerical value
	v, err := numToBinary(s)
	if err == nil {
		return v
	}

	// Check if the symbol is predefined
	v = getSymbol(s)
	if v != "" {
		val, err := numToBinary(v)
		if err != nil {
			// ???
		}
		return val
	}

	// Check if the symbol has a translation stored (is a label symbol)
	v, ok := p.labels[s]
	if ok == true {
		val, err := numToBinary(v)
		if err != nil {
			// ???
		}
		return val
	}

	// store the symbol for later translation
	if p.extra[s] == nil {
		p.extra[s] = make([]int, 0)
	}
	p.extra[s] = append(p.extra[s], i)
	if !contains(p.extraKeys, s) {
		p.extraKeys = append(p.extraKeys, s)
	}
	return s
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

func delNonCode(s string) string {
	ind := strings.Index(s, "//")
	if ind != -1 {
		s = s[:ind]
	}
	return strings.TrimSpace(s)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
