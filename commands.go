package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name		string
	description string
	callback 	func() error
}


func getCommands() map[string]cliCommand {
	cmds := make(map[string]cliCommand)

	cmds["help"] = cliCommand{
		name: "help",
		description: "Displays a help message",
		callback: commandHelp,
	}

	cmds["exit"] = cliCommand{
		name: "exit",
		description: "Exit the Pokedex",
		callback: commandExit,
	}

	return cmds
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")

	os.Exit(0)

	return nil
}


func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}


