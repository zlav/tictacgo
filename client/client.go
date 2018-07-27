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
	c.setup()

	for c.playRound() {
	}
}

func (c *tictacclient) setup() {
	fmt.Print(c.server.TurnOn())

	setup := false
	var output string
	for !setup {
		setup, output = c.server.SetupGame(userInputString())
		fmt.Print(output)
	}
}

func (c *tictacclient) playRound() bool {
	fmt.Print(c.server.PrintHelp())
	c.playGame(tutorial())

	_, output := c.server.Status()
	fmt.Print(output)

	replay, output := c.server.Reset(userInputString())
	fmt.Print(output)
	return replay
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

		printBoard(c, tutorial)
	}
}

func userInputString() string {
	input := ""
	fmt.Scanf("%s", &input)
	return input
}

func printBoard(c *tictacclient, tutorial bool) {
	if tutorial {
		fmt.Print(c.server.PrintHelp())
	}
	fmt.Print(c.server.PrintGame())
}

func tutorial() bool {
	fmt.Printf("Enable tutorial mode for a guided walkthrough how to place tokens? Y/N: ")
	input := userInputString()
	if input[0] == 'y' || input[0] == 'Y' {
		return true
	}
	return false
}
