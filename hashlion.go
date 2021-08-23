package main

import (
	"fmt"
	consoleutil "hashlion/utils"
	fileutil "hashlion/utils"
	hashutil "hashlion/utils"
	timeutil "hashlion/utils"
	"strconv"
	"strings"
	"time"
)

func main() {

	const version string = "Build 1 Alpha"
	const program_name string = "HashLion"

	consoleutil.SetConsoleTitle(program_name + " " + version)
	fmt.Println("===================================")
	fmt.Println(program_name + " " + version + ", By Nort721")
	fmt.Println("===================================")
	fmt.Println("")

	fmt.Print("enter word list path(.txt): ")
	input := consoleutil.ScanInput()

	// if the user didn't write .txt at the end, we will just add it our self
	if !strings.Contains(input, ".txt") {
		input += ".txt"
	}

	// read words file
	fmt.Println("loading file . . .")
	fmt.Println()
	var lines []string = fileutil.ScanFile(input)

	// main(input) loop variables
	var running bool = true

	var hashType string

	for running {
		// get the hash type
		fmt.Print("enter hash type(sha1/sha256/sha512/md5): ")
		input = consoleutil.ScanInput()
		hashType = input

		if input == "exit" {
			running = false
		}

		if input != "sha1" && input != "sha256" && input != "sha512" && input != "md5" {
			fmt.Println("error -> bad syntax, args:{sha1,sha256,sha512,md5}")
		} else {
			break
		}
	}

	// input loop
	for running {

		// get the target
		fmt.Print("enter target hash: ")

		input = consoleutil.ScanInput()

		if input == "exit" {
			running = false
			continue
		}

		// the target hash
		target := input

		// ask if to show attack live info
		fmt.Println("Tip! attack is quicker on hide mode")
		fmt.Print("show attack live info?(show/hide): ")

		input = consoleutil.ScanInput()

		if input == "exit" {
			running = false
			continue
		}

		// if no input, choose hide by default as its the fastest
		if len(input) == 0 {
			//fmt.Print("hide")
			input = "hide"
		}

		var cracked bool = false

		fmt.Println()
		if input == "show" {
			cracked = attackAndShow(lines, hashType, target)
		} else if input == "hide" {
			cracked = attackQuick(lines, hashType, target)
		} else {
			fmt.Println("error -> bad syntax, args:{show,hide}")
			continue
		}

		if !cracked {
			fmt.Println("Failed to crack, reason: hash was not in dictionary")
		}
	}
}

// attacks the given target and provides live info of the attack
func attackAndShow(lines []string, hashType string, target string) bool {
	startDate := time.Now()
	fmt.Println("Starting cracking attempts at " + startDate.Format("01-02-2006 15:04:05"))

	var counter int = 0

	var testHash string

	// the hash check is intentionally here instead of having it inside a central
	// having function so we can check the hash here once outside of the attack loop
	// making the attack faster
	if hashType == "sha1" {

		for _, line := range lines {

			counter++

			testHash = hashutil.GenerateSha1(line)

			fmt.Printf("trying -> %v : %v | %v\n", testHash, line, strconv.Itoa(counter))
			//fmt.Println("trying -> " + testHash + " : " + line + " | " + strconv.Itoa(counter))

			if testHash == target {
				finishAttack(startDate, target, line, counter)
				return true
			}
		}

	} else if hashType == "sha256" {

		for _, line := range lines {

			counter++

			testHash = hashutil.GenerateSha256(line)

			fmt.Printf("trying -> %v : %v | %v\n", testHash, line, strconv.Itoa(counter))
			//fmt.Println("trying -> " + testHash + " : " + str + " | " + strconv.Itoa(counter))

			if testHash == target {
				finishAttack(startDate, target, line, counter)
				return true
			}
		}

	} else if hashType == "sha512" {

		for _, line := range lines {

			counter++

			testHash = hashutil.GenerateSha512(line)

			fmt.Printf("trying -> %v : %v | %v\n", testHash, line, strconv.Itoa(counter))

			if testHash == target {
				finishAttack(startDate, target, line, counter)
				return true
			}
		}

	} else if hashType == "md5" {

		for _, line := range lines {

			counter++

			testHash = hashutil.GenerateMD5(line)

			fmt.Printf("trying -> %v : %v | %v\n", testHash, line, strconv.Itoa(counter))

			if testHash == target {
				finishAttack(startDate, target, line, counter)
				return true
			}
		}

	}

	return false
}

// attacks the given hash target as fast as possible
// giving up whatever is possible to speed up the attack
func attackQuick(lines []string, hashType string, target string) bool {
	startDate := time.Now()
	fmt.Println("Starting cracking attempts at " + startDate.Format("01-02-2006 15:04:05"))

	var testHash string

	// the hash check is intentionally here instead of having it inside a central
	// having function so we can check the hash here once outside of the attack loop
	// making the attack faster
	if hashType == "sha1" {

		for _, line := range lines {

			testHash = hashutil.GenerateSha1(line)

			if testHash == target {
				finishAttackQuick(startDate, target, line)
				return true
			}
		}

	} else if hashType == "sha256" {

		for _, line := range lines {

			testHash = hashutil.GenerateSha256(line)

			if testHash == target {
				finishAttackQuick(startDate, target, line)
				return true
			}
		}

	} else if hashType == "sha512" {

		for _, line := range lines {

			testHash = hashutil.GenerateSha512(line)

			if testHash == target {
				finishAttackQuick(startDate, target, line)
				return true
			}
		}

	} else if hashType == "md5" {

		for _, line := range lines {

			testHash = hashutil.GenerateMD5(line)

			if testHash == target {
				finishAttackQuick(startDate, target, line)
				return true
			}
		}

	}

	return false
}

// this will print out the finishing attack message
func finishAttack(startDate time.Time, target string, line string, counter int) {
	fmt.Println("Hash cracked successfully! " + target + " : " + line + " | " + strconv.Itoa(counter))
	fmt.Println("Started cracking attempts at " + startDate.Format("01-02-2006 15:04:05"))
	finishDate := time.Now()
	fmt.Println("Finished cracking at " + finishDate.Format("01-02-2006 15:04:05"))
	fmt.Println("Time took to crack: " + timeutil.GetTimeBetweenDates(finishDate, startDate))
	fmt.Println()
}

// this will print out the finishing attack message
func finishAttackQuick(startDate time.Time, target string, line string) {
	fmt.Println("Hash cracked successfully! " + target + " : " + line)
	finishDate := time.Now()
	fmt.Println("Finished cracking at " + finishDate.Format("01-02-2006 15:04:05"))
	fmt.Println("Time took to crack: " + timeutil.GetTimeBetweenDates(finishDate, startDate))
	fmt.Println()
}
