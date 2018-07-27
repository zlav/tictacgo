package client

import (
	"fmt"

	"github.com/zlav/tictacgo/director"
)

type tictacclient struct {
	server *director.Director
}

func NewTicTacClient(server *director.Director) *tictacclient {
	return &tictacclient{server: server}
}

func (c *tictacclient) Run() {
	fmt.Print(c.server.TurnOn())

	setup := false
	output := ""
	for !setup {
		setup, output = c.server.SetupGame(userInputString())
		fmt.Print(output)
	}

	connected := true
	for connected {
		fmt.Print(c.server.PrintHelp())
		c.playGame(tutorial())
		_, output := c.server.Status()
		fmt.Print(output)
		connected, output = c.server.Reset(userInputString())
		fmt.Print(output)
	}
}

func (c *tictacclient) playGame(tutorial bool) {
	fmt.Print(c.server.PrintGame())

	for !c.server.GameOver() {
		status, output := c.server.Status()
		fmt.Print(output)
		if status {
		} else if c.server.PlayerTurn() {
			hasPlayed, output := c.server.Play(userInputString())
			for !hasPlayed {
				fmt.Print(output)
				hasPlayed, output = c.server.Play(userInputString())
			}
		}

		if tutorial {
			fmt.Print(c.server.PrintHelp())
		}
		fmt.Print(c.server.PrintGame())
	}
}

func userInputString() string {
	input := ""
	fmt.Scanf("%s", &input)
	return input
}

func tutorial() bool {
	fmt.Printf("Enable tutorial mode for a guided walkthrough how to place tokens? Y/N: ")
	input := userInputString()
	if input[0] == 'y' || input[0] == 'Y' {
		return true
	}
	return false
}
