package main

import (
	"strings"
	"testing"
)

func TestAinstructionNumbers(t *testing.T) {
	// What is the number of greatest magnitude expected?
	//negative numbers are not legal--can expect a valid input for this exercise
	testCases := []string{"@84", "@0", "@notNumber", "notAType"}
	expected := []string{"000000001010100", "000000000000000"}

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
