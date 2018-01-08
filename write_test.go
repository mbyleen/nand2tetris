package main

import (
	"testing"
)

func TestWrite(t *testing.T) {

	// create a short test output
	d := []string{"line 1", "line 2", "line 3"}
	filename := "test.asm"

	// run the write function on the test file
	err := write(d, filename)
	if err != nil {
		t.Log(err)
	}
	// compare manually
	// create a test by comparing to a test file
}
