package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)

  for {
    fmt.Print(prompt)

    scanner.Scan()
    rawCmd := scanner.Text()
    cmdWords := cleanInput(rawCmd)

    if len(cmdWords) == 0 {
      continue
    }

    fmt.Printf("Your command was: %s\n", cmdWords[0])
  }

}

