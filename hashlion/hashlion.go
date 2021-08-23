package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

func main() {

	const version string = "Build 1 Alpha"
	const program_name string = "HashLion"

	SetConsoleTitle(program_name + " " + version)
	fmt.Println("===================================")
	fmt.Println(program_name + " " + version + ", By Nort721")
	fmt.Println("===================================")
	fmt.Println("")

	fmt.Print("enter word list path(.txt): ")
	input := scanInput()

	// if the user didn't write .txt at the end, we will just add it our self
	if !strings.Contains(input, ".txt") {
		input += ".txt"
	}

	// read words file
	fmt.Println("loading file . . .")
	fmt.Println()
	var lines []string = scanFile(input)

	// main(input) loop variables
	var running bool = true

	var hashType string

	for running {
		// get the hash type
		fmt.Print("enter hash type(sha1/sha256/sha512/md5): ")
		input = scanInput()
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

		input = scanInput()

		if input == "exit" {
			running = false
			continue
		}

		// the target hash
		target := input

		// ask if to show attack live info
		fmt.Println("Tip! attack is quicker on hide mode")
		fmt.Print("show attack live info?(show/hide): ")

		input = scanInput()

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

			testHash = generateSha1(line)

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

			testHash = generateSha256(line)

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

			testHash = generateSha512(line)

			fmt.Printf("trying -> %v : %v | %v\n", testHash, line, strconv.Itoa(counter))

			if testHash == target {
				finishAttack(startDate, target, line, counter)
				return true
			}
		}

	} else if hashType == "md5" {

		for _, line := range lines {

			counter++

			testHash = generateMD5(line)

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

			testHash = generateSha1(line)

			if testHash == target {
				finishAttackQuick(startDate, target, line)
				return true
			}
		}

	} else if hashType == "sha256" {

		for _, line := range lines {

			testHash = generateSha256(line)

			if testHash == target {
				finishAttackQuick(startDate, target, line)
				return true
			}
		}

	} else if hashType == "sha512" {

		for _, line := range lines {

			testHash = generateSha512(line)

			if testHash == target {
				finishAttackQuick(startDate, target, line)
				return true
			}
		}

	} else if hashType == "md5" {

		for _, line := range lines {

			testHash = generateMD5(line)

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
	fmt.Println("Time took to crack: " + getTimeBetweenDates(finishDate, startDate))
	fmt.Println()
}

// this will print out the finishing attack message
func finishAttackQuick(startDate time.Time, target string, line string) {
	fmt.Println("Hash cracked successfully! " + target + " : " + line)
	finishDate := time.Now()
	fmt.Println("Finished cracking at " + finishDate.Format("01-02-2006 15:04:05"))
	fmt.Println("Time took to crack: " + getTimeBetweenDates(finishDate, startDate))
	fmt.Println()
}

// will scan input for full sentences
// the normal scan command only scans until it gets to a space
// effectively only scanning one word, we want to scan the whole input
// in case we have commands with arguments
func scanInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		return scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return "-999"
}

// generates the sha1 hash of a given string
func generateSha1(text string) string {
	hash := sha1.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// generates the sha256 hash of a given string
func generateSha256(text string) string {
	hasher := sha256.New()

	hasher.Write([]byte(text))

	return hex.EncodeToString(hasher.Sum(nil))
}

// generates the sha256 hash of a given string
func generateSha512(text string) string {
	hasher := sha512.New()

	hasher.Write([]byte(text))

	return hex.EncodeToString(hasher.Sum(nil))
}

// generates the md5 hash of a given string
func generateMD5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// loads the text file thats in the given path
// and reads all the text in it, returns
// a string slice with all the text from the file
func scanFile(path string) []string {
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

// returns the amount of time between two dates
// with the time unit attached
func getTimeBetweenDates(finishDate time.Time, startDate time.Time) string {

	timeDiff := finishDate.Sub(startDate)

	diff := timeDiff.Hours()

	unit := "hours"

	// ToDo: find a better way to do this
	if diff < 1 {

		diff = timeDiff.Minutes()

		unit = "minutes"

		if diff < 1 {

			unit = "seconds"

			diff = timeDiff.Seconds()

			if diff < 1 {

				unit = "milliseconds"

				diff = float64(timeDiff.Milliseconds())

				if diff < 1 {

					unit = "microseconds"

					diff = float64(timeDiff.Microseconds())
				}

			}
		}
	}

	return strconv.FormatFloat(diff, 'f', 1, 64) + " " + unit
}

// sets the windows console title to a given string
func SetConsoleTitle(title string) (int, error) {
	handle, err := syscall.LoadLibrary("Kernel32.dll")
	if err != nil {
		return 0, err
	}
	defer syscall.FreeLibrary(handle)
	proc, err := syscall.GetProcAddress(handle, "SetConsoleTitleW")
	if err != nil {
		return 0, err
	}
	r, _, err := syscall.Syscall(proc, 1, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))), 0, 0)
	return int(r), err
}
