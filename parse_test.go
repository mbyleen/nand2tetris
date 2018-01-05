package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestAinstructionNumbers(t *testing.T) {
	//what are the pos and neg numbers of greatest magnitude expected?
	//negative numbers are not legal--can expect a valid input for this exercise
	input := []string{"@84", "@0", "@notNumber", "notAType"}
	// parse the test file
	output, err := parse(input)
	if err != nil {
		t.Error()
	}
	// test against expected output
	expected := []string{"000000001010100", "000000000000000"}
	for i, _ := range output {
		fmt.Println(output[i], expected[i])

		if strings.Compare(output[i], expected[i]) != 0 {
			t.Error("Match failed:", output[i], expected[i])
		}
	}
}
