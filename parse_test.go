package main

import (
	"strings"
	"testing"
)

func TestAinstructionNumbers(t *testing.T) {
	// What is the number of greatest magnitude expected?
	//negative numbers are not legal--can expect a valid input for this exercise
	input := []string{"@84", "@0", "@notNumber", "notAType"}
	// parse the test file
	output, err := parse(input)
	if err != nil {
		t.Error()
	}
	// test against expected output
	expected := []string{"000000001010100", "000000000000000"}
	for i, _ := range expected {
		if strings.Compare(output[i], expected[i]) != 0 {
			t.Error("Match failed:", output[i], expected[i])
		}
	}
}

func TestCinstruction(t *testing.T) {
	// limited test cases--will test looku tables with longer examples
	testCases := []string{"A=D+1", "0;JMP", "D;JNE"}
	expected := []string{"1110011111100000", "1110101010000111", "1110001100000101"}
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
