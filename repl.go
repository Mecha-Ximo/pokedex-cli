package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
	  fmt.Print(prompt)
  
	  scanner.Scan()
	  rawCmd := scanner.Text()
	  cmdWords := cleanInput(rawCmd)
  
	  if len(cmdWords) == 0 {
		continue
	  }
	  
	  cmdKey := cmdWords[0]
	  cmd, exists := getCommands()[cmdKey]

	  if !exists {
		fmt.Println("Unknown command")
		continue
	  }
	  
	  cmd.callback()
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
  
	return words
}