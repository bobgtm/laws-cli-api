package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startAPI() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Laws CLI > ")
		reader.Scan()

		userInput := cleanInput(reader.Text())
		if len(userInput) == 0 {
			continue
		}

		commandName := userInput[0]
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown Command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "displays a help message",
			callback:    commandHelp,
		},
		"show me": {
			name:        "show me",
			description: "shows most recent legislation",
			callback:    commandShowMaster,
		},
	}
}
