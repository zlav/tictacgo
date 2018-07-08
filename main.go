package main

import (
	"fmt"

	"github.com/zlav/tictacgo/board"
)

func main() {
	fmt.Println("Welcome to Tic Tac Go")
	game := board.NewGame()
	game.PrintBoard()

	// TODO: Send it to whichever game you'd like to see 1 player, 2 player, watch computers?
	// This can be abstracted later. Build it here now

	// TODO: Print the results and options for next game
}
