package utils

import (
	"bufio"
	"log"
	"os"
	"time"
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

func WriteFile(name string, text []string) {
	f, err := os.Create(name + ".txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	now := time.Now()

	_, err1 := f.WriteString("=+= Generated by hashlion " + now.Format("01-02-2006 15:04:05") + " =+=\n")

	if err1 != nil {
		log.Fatal(err1)
	}

	for _, line := range text {
		_, err2 := f.WriteString(line + "\n")

		if err2 != nil {
			log.Fatal(err2)
		}
	}

}
