package main

import (
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	filename := "testfiles/add/Add.asm"
	compareFile := "testfiles/add/Add.hack"
	testFromFile(t, filename, compareFile)
}

func TestMaxL(t *testing.T) {
	filename := "testfiles/max/MaxL.asm"
	compareFile := "testfiles/max/MaxL.hack"
	testFromFile(t, filename, compareFile)
}

func TestMax(t *testing.T) {
	filename := "testfiles/max/Max.asm"
	compareFile := "testfiles/max/Max.hack"
	testFromFile(t, filename, compareFile)
}

func TestRectL(t *testing.T) {
	filename := "testfiles/rect/RectL.asm"
	compareFile := "testfiles/rect/RectL.hack"
	testFromFile(t, filename, compareFile)
}

func TestRect(t *testing.T) {
	filename := "testfiles/rect/Rect.asm"
	compareFile := "testfiles/rect/Rect.hack"
	testFromFile(t, filename, compareFile)
}

func TestPongL(t *testing.T) {
	filename := "testfiles/pong/PongL.asm"
	compareFile := "testfiles/pong/PongL.hack"
	testFromFile(t, filename, compareFile)
}

func TestPong(t *testing.T) {
	filename := "testfiles/pong/Pong.asm"
	compareFile := "testfiles/pong/Pong.hack"
	testFromFile(t, filename, compareFile)
}

func testFromFile(t *testing.T, filename string, compareFile string) {
	code, err := read(filename)
	if err != nil {
		t.Error("type: %T; value: %q\n", err, err)
	}

	output, err := parse(code)
	if err != nil {
		t.Error("Parsing error")
	}

	compare, err := read(compareFile)
	if err != nil {
		t.Error("type: %T; value: %q\n", err, err)
	}

	for i := range compare {
		if strings.Compare(output[i], compare[i]) != 0 {
			t.Errorf("Comparision failure at line", i+1)
		}
	}
}
