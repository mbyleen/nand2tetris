package main

import (
	"strings"
	"testing"
)

func TestParseLabelInOrder(t *testing.T) {
	testCase := []string{"(symbol)", "@symbol"}
	expected := []string{"0000000000000000"}

	testParse(t, testCase, expected)
}

func TestParseLabelReverseOrder(t *testing.T) {
	testCase := []string{"@symbol", "(symbol)", "@8"}
	expected := []string{"0000000000000001", "0000000000001000"}

	testParse(t, testCase, expected)
}

func TestParsePredefined(t *testing.T) {
	// Limited examples - test lookup table with provided .asm files
	testCase := []string{"@R2", "@KBD"}
	expected := []string{"0000000000000010", "0110000000000000"}

	testParse(t, testCase, expected)
}

func TestAinstructionNumbers(t *testing.T) {
	// What is the number of greatest magnitude expected?
	//negative numbers are not legal--can expect a valid input for this exercise
	testCases := []string{"(notNumber)", "@84", "@0", "@16384", "@notNumber"}
	expected := []string{"0000000001010100", "0000000000000000", "0100000000000000"}

	testParse(t, testCases, expected)
}

func TestCinstruction(t *testing.T) {
	// limited test cases--will test lookup tables with longer examples
	testCases := []string{"A=D+1", "0;JMP", "D;JNE"}
	expected := []string{"1110011111100000", "1110101010000111", "1110001100000101"}
	testParse(t, testCases, expected)
}

func TestRemoveComments(t *testing.T) {
	testCases := []string{"// This is a comment", "A=D+1 // This is an inline comment"}
	expected := []string{"", "A=D+1 "}

	for i, line := range testCases {
		output := removeComments(line)
		if strings.Compare(expected[i], output) != 0 {
			t.Error("Match failed:", output, expected[i])
		}
	}
}

func testParse(t *testing.T, testCases []string, expected []string) {
	output, err := parse(testCases)
	if err != nil {
		t.Error()
	}
	for i, _ := range expected {
		if strings.Compare(expected[i], output[i]) != 0 {
			t.Error("Match failed:", output[i], expected[i])
		}
	}
}
