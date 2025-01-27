package main

import (
	"fmt"
	"os"

	"github.com/Mecha-Ximo/pokedex-cli/pokemon_api"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var getNext, getPrev = generateMapGetters()

func getCommands() map[string]cliCommand {
	cmds := make(map[string]cliCommand)


	cmds["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}

	cmds["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}

	cmds["map"] = cliCommand{
		name:        "map",
		description: "Get location areas",
		callback:    getNext,
	}

	cmds["mapb"] = cliCommand{
		name:        "mapb",
		description: "Get previous location areas",
		callback:    getPrev,
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


func generateMapGetters() (func() error, func() error) {
	next := location_areas_url
	prev := ""

	getNext := func() error {
		if next == "" {
			fmt.Println("You are in the last page!")
			return nil
		}

		locations := []pokemon_api.LocationArea{}
		locations, next, prev = pokemon_api.RequestLocationAreas(next)

		printLocations(locations)

		return nil
	}

	getPrevious := func() error {
		if prev == "" {
			fmt.Println("you're on the first page")
			return nil
		}
		
		locations := []pokemon_api.LocationArea{}
		locations, next, prev = pokemon_api.RequestLocationAreas(prev)

		printLocations(locations)

		return nil
	}

	return getNext, getPrevious
}

func printLocations(locations []pokemon_api.LocationArea) {
	for _, location := range locations {
		fmt.Println(location.Name)
	}
}