package main

import (
	"fmt"
	command "hashlion/utils"
	consoleutil "hashlion/utils"
	"strings"
)

func main() {

	// logs
	// fixed os.open issue when copy-pasting file locations on unix based systems
	// added bcrypt

	const version string = "Build 3 Alpha"
	const program_name string = "HashLion"

	//consoleutil.SetConsoleTitle(program_name + " " + version)
	fmt.Println("===================================")
	fmt.Println(program_name + " " + version + ", By Nort721")
	fmt.Println("===================================")

	commands := make(map[string]command.Command)

	commands["help"] = command.CommandHelp{}
	commands["exit"] = command.CommandExit{}

	commands["attack"] = command.CommandAttack{}
	commands["a"] = command.CommandAttack{}

	commands["attackssh"] = command.CommandAttackSSH{}
	commands["assh"] = command.CommandAttackSSH{}

	commands["hash"] = command.CommandHash{}

	var input string

	for {

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
