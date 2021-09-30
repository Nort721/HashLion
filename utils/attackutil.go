package utils

import (
	"fmt"
	generator "hashlion/types"
	"strconv"
	"time"
)

var lines []string

func AttackDictionary(hashalgo string, target string, wordListPath string, print bool, layers int) {

	hashes := make(map[string]generator.Generator)

	hashes["sha1"] = generator.HashSha1{}
	hashes["sha256"] = generator.HashSha256{}
	hashes["sha512"] = generator.HashSha512{}
	hashes["md5"] = generator.HashMd5{}
	hashes["bcrypt"] = generator.HashBcrypt{}

	fmt.Println()

	var counter int = 0

	var testHash string

	// read words file
	PrintMsg("loading words file . . .\n")
	lines = ScanFile(wordListPath)

	cracked := false

	startDate := time.Now()
	fmt.Println("Starting cracking attempts at " + startDate.Format("01-02-2006 15:04:05"))

	for _, line := range lines {

		counter++

		gen := hashes[hashalgo]

		testHash = generator.Run_Hash(gen, line)

		if print {
			fmt.Printf("trying -> %v : %v | %v\n", testHash, line, strconv.Itoa(counter))
		}

		if testHash == target {
			finishAttack(startDate, target, line, counter)
			cracked = true
			break
		}
	}

	if !cracked {
		fmt.Println("Failed to crack, reason: hash was not in dictionary | " + strconv.Itoa(counter) + " size: " + strconv.Itoa(len(lines)))
	}

}

func AttackBruteForce(hashalgo string, target string, format string, print bool, layers int) {
	fmt.Printf("attackmode: brutefoece, algo: %v, target: %v, format: %v, print: %v, layers: %v\n",
		hashalgo, target, format, print, layers)
}

func AttackSSH(wordListPath string, ip string, port string, print bool) {

}

// this will print out the finishing attack message
func finishAttack(startDate time.Time, target string, line string, counter int) {
	line1 := "Hash cracked successfully! " + target + " : " + line + " | " + strconv.Itoa(counter)
	line2 := "Started cracking attempts at " + startDate.Format("01-02-2006 15:04:05")
	finishDate := time.Now()
	line3 := "Finished cracking at " + finishDate.Format("01-02-2006 15:04:05")
	line4 := "Time took to crack: " + GetTimeBetweenDates(finishDate, startDate)

	text := []string{line1, line2, line3, line4}

	for _, line := range text {
		fmt.Println(line)
	}
	fmt.Println()

	WriteFile("hashlion_log", text)
}

// this will print out the finishing attack message
func finishAttackQuick(startDate time.Time, target string, line string) {
	line1 := "Hash cracked successfully! " + target + " : " + line
	finishDate := time.Now()
	line2 := "Finished cracking at " + finishDate.Format("01-02-2006 15:04:05")
	line3 := "Time took to crack: " + GetTimeBetweenDates(finishDate, startDate)

	text := []string{line1, line2, line3}

	for _, line := range text {
		fmt.Println(line)
	}
	fmt.Println()

	WriteFile("hashlion_log", text)
}
