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
	for !c.server.SetupGame(userInputString()) {
	}

	connected := true
	for connected {
		c.server.PrintHelp()
		c.playGame(tutorial())
		c.server.Status()
		connected = c.server.Reset(userInputString())
	}
}

func (c *tictacclient) playGame(tutorial bool) {
	c.server.PrintGame()

	for !c.server.GameOver() {
		if c.server.Status() {
		} else if c.server.PlayerTurn() {
			for !c.server.Play(userInputString()) {
			}
		}

		if tutorial {
			c.server.PrintHelp()
		}
		c.server.PrintGame()
	}
}

func userInputString() string {
	input := ""
	fmt.Scanf("%s", &input)
	return input
}

func tutorial() bool {
	fmt.Printf("Enable tutorial mode for a guided walkthrough how to place tokens? Y/N:")
	input := userInputString()
	if input[0] == 'y' || input[0] == 'Y' {
		return true
	}
	return false
}
