package utils

import (
	"bufio"
	"log"
	"os"
)

// loads the text file thats in the given path
// and reads all the text in it, returns
// a string slice with all the text from the file
func ScanFile(path string) []string {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
