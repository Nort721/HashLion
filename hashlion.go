package main

import (
	"fmt"
	command "hashlion/utils"
	consoleutil "hashlion/utils"
	fileutil "hashlion/utils"
	hashutil "hashlion/utils"
	timeutil "hashlion/utils"
	"strconv"
	"strings"
	"time"
)

var lines []string
var Running bool = true

func main() {

	// logs
	// fixed os.open issue when copy-pasting file locations on unix based systems
	// added bcrypt

	str := hashutil.GenerateBcrypt("test")
	fmt.Println(str)

	const version string = "Build 2 Alpha"
	const program_name string = "HashLion"

	//consoleutil.SetConsoleTitle(program_name + " " + version)
	fmt.Println("===================================")
	fmt.Println(program_name + " " + version + ", By Nort721")
	fmt.Println("===================================")

	commands := make(map[string]command.Command)

	commands["help"] = command.CommandHelp{}
	commands["exit"] = command.CommandExit{}
	commands["attack"] = command.CommandAttack{}

	var input string

	for Running {

		fmt.Print("input -> ")

		input = consoleutil.ScanInput()

		args := strings.Split(input, " ")

		cmd, ok := commands[args[0]]

		if ok {
			command.Run_Command(cmd, args)
		} else {
			consoleutil.PrintMsg("Unknown command. Type help for help.\n")
		}

	}
}

// attacks the given target and provides live info of the attack
func AttackAndShow(hashType string, target string) bool {
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

	} else if hashType == "bcrypt" {

		for _, line := range lines {

			counter++

			testHash = hashutil.GenerateBcrypt(line)

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

	} else if hashType == "bcrypt" {

		for _, line := range lines {

			testHash = hashutil.GenerateBcrypt(line)

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
	line1 := "Hash cracked successfully! " + target + " : " + line + " | " + strconv.Itoa(counter)
	line2 := "Started cracking attempts at " + startDate.Format("01-02-2006 15:04:05")
	finishDate := time.Now()
	line3 := "Finished cracking at " + finishDate.Format("01-02-2006 15:04:05")
	line4 := "Time took to crack: " + timeutil.GetTimeBetweenDates(finishDate, startDate)

	text := []string{line1, line2, line3, line4}

	for _, line := range text {
		fmt.Println(line)
	}
	fmt.Println()

	fileutil.WriteFile("hashlion_log", text)
}

// this will print out the finishing attack message
func finishAttackQuick(startDate time.Time, target string, line string) {
	line1 := "Hash cracked successfully! " + target + " : " + line
	finishDate := time.Now()
	line2 := "Finished cracking at " + finishDate.Format("01-02-2006 15:04:05")
	line3 := "Time took to crack: " + timeutil.GetTimeBetweenDates(finishDate, startDate)

	text := []string{line1, line2, line3}

	for _, line := range text {
		fmt.Println(line)
	}
	fmt.Println()

	fileutil.WriteFile("hashlion_log", text)
}
