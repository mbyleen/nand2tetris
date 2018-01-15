package main

import (
	"io/ioutil"
	"testing"
)

func TestRead(t *testing.T) {

	// create a short test file
	filepath := "/tmp/test1.asm"
	text := []byte("line 1\nline2\nline3")
	err := ioutil.WriteFile(filepath, text, 0644)
	if err != nil {
		t.Log(err)
	}

	// run the read function on the test file
	input, err := read(filepath)
	if err != nil {
		t.Log(err)
	}

	// compare it to desired output
	d := [...]string{"line 1", "line 2", "line 3"}
	for i := 0; i < len(d); i++ {
		if d[i] != input[i] {
			t.Log(err)
		}
	}
}
