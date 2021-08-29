package utils

import (
	"os"
	"strconv"
	"strings"
)

/* commands system */
// interface example --->
type Command interface {
	OnCommand(args []string)
}

// this will not accept the argument if the func is not implemented
func Run_Command(cmd Command, args []string) {
	cmd.OnCommand(args)
}

// help command
//---
type CommandHelp struct{}

// impleneting the func that the interface demands
func (c CommandHelp) OnCommand(args []string) {
	PrintMsg(
		"\n====================\n" +
			"exit\n" +
			"attack -dictionary/d <wordlist> <hashalgo> <target> -print/p <layers(def:1)>\n" +
			"attack -brutefoce/b <format> <hashalgo> <target> -print/p <layers(def:1)>\n" +
			"====================\n")
}

//---

// logs command
//---
type CommandExit struct{}

// impleneting the func that the interface demands
func (c CommandExit) OnCommand(args []string) {
	os.Exit(3)
}

//---

// attack command
//                      0          1           2          3         4        5        6
// syntax dictionary: attack -dictionary/d <wordlist> <hashalgo> <target> -print/p <layers>
// syntax bruteforce: attack -brutefoce/b <format> <hashalgo> <target> -print/p <layers>
//---
type CommandAttack struct{}

// impleneting the func that the interface demands
func (c CommandAttack) OnCommand(args []string) {

	attackMode := args[1]

	length := len(args)

	if length > 3 {

		target := args[4]
		hashalgo := args[3]

		// get printing mode
		var print bool = false

		if length > 5 {
			print = (strings.Contains(args[5], "p"))
		}

		// get hashing layers
		layers := 1

		if length > 6 {
			layers, _ = strconv.Atoi(args[6])
		}

		if strings.Contains(attackMode, "-d") {

			wordListPath := args[2]

			// if the user didn't write .txt at the end, we will just add it our self
			if !strings.Contains(wordListPath, ".txt") {
				wordListPath += ".txt"
			}

			AttackDictionary(hashalgo, target, wordListPath, print, layers)

		} else if strings.Contains(attackMode, "-b") {

			format := args[2]

			AttackBruteForce(hashalgo, target, format, print, layers)

		}

	}
}

//---
