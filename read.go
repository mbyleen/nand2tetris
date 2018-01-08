package main

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func read(filename string) ([]string, error) {
	if strings.TrimSuffix(filename, ".asm") == filename {
		return nil, errors.New("read: given file lacks .asm filetype")
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
