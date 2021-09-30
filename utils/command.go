package utils

import (
	generator "hashlion/types"
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
			" - crack/c -dictionary/d <wordlist> <hashalgo> <target> -print/p <layers(def:1)>\n" +
			" - crack/c -brutefoce/b <format> <hashalgo> <target> -print/p <layers(def:1)>\n" +
			" - attackssh/assh <wordlist> <host:port> -print/p\n" +
			" - hash <hashalgo> <string>\n" +
			" - exit\n" +
			"\n" +
			"Supported hashes:\n" +
			" - md5\n" +
			" - sha1\n" +
			" - sha256\n" +
			" - sha384\n" +
			" - sha512\n" +
			" - bcrypt\n" +
			"====================\n")
}

//---

// logs command
//---
type CommandExit struct{}

// impleneting the func that the interface demands
func (c CommandExit) OnCommand(args []string) {
	PrintMsg("exiting . . .")
	os.Exit(1)
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

	if len(args) == 1 {
		PrintMsg("Incorrect args length, Type help for help.\n")
		return
	}

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

		} else {
			PrintMsg("Incorrect syntax, attack mode " + attackMode + " doesn't exist")
		}

	} else {
		PrintMsg("Incorrect args length, Type help for help.\n")
	}
}

//---

// attackssh command
//                        0         1          2         3
// syntax dictionary: attackssh <wordlist> <ip:port> -print/p
//---
type CommandAttackSSH struct{}

// impleneting the func that the interface demands
func (c CommandAttackSSH) OnCommand(args []string) {

	if len(args) == 4 {

		wordListPath := args[1]

		// if the user didn't write .txt at the end, we will just add it our self
		if !strings.Contains(wordListPath, ".txt") {
			wordListPath += ".txt"
		}

		address := args[2]

		ip := address

		port := address

		print := (strings.Contains(args[3], "p"))

		AttackSSH(wordListPath, ip, port, print)

	} else {
		PrintMsg("Incorrect args length, Type help for help.\n")
	}
}

//---

// hash command
//                     0        1         2
// syntax dictionary: hash <hashalgo> <string>
//---
type CommandHash struct{}

// impleneting the func that the interface demands
func (c CommandHash) OnCommand(args []string) {

	if len(args) == 3 {

		hashalgo := args[1]

		str := args[2]

		PrintMsg("Hashed(" + str + "): " + generator.GenerateHash(str, hashalgo) + "\n")
	} else {
		PrintMsg("Incorrect args length, Type help for help.\n")
	}
}

//---
