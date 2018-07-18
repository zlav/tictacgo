package main

import (
	"fmt"

	"github.com/zlav/tictacgo/board"
	"github.com/zlav/tictacgo/player"
	"github.com/zlav/tictacgo/symbol"
)

func main() {
	fmt.Printf("Welcome to Tic Tac Go\n")

	var players []player.Player
	setup := false
	fmt.Printf("How would you like to play? [1] player, []2 player, or watch the [C]omputer fight itself?\n")
	xSym := symbol.NewSymbol("X")
	oSym := symbol.NewSymbol("O")
	for !setup {
		setup = true
		input := ""
		fmt.Scanf("%s", &input)
		switch input[0] {
		case '1':
			players = []player.Player{player.NewHuman(xSym), player.NewComputer(oSym)}
		case '2':
			players = []player.Player{player.NewHuman(xSym), player.NewHuman(oSym)}
		case 'c', 'C', '3':
			players = []player.Player{player.NewComputer(xSym), player.NewComputer(oSym)}
		default:
			fmt.Printf("Invalid Input\n")
			setup = false
		}
	}

	game := board.NewGame(players[0].GetIcon(), players[1].GetIcon())
	fmt.Printf("Here is an example of how to select a box on the grid\n")
	game.PrintHelp()
	fmt.Printf("Good Luck!\n")

	playing := true
	for playing {
		for !game.IsDraw() && !game.IsWon() {
			for _, player := range players {
				game.PrintBoard()
				player.PlayTicTacToe(game)
				if game.IsDraw() {
					fmt.Printf("Draw! Nobody is a winner\n")
					break
				}
				if game.IsWon() {
					fmt.Printf("Congrats to Player %s!\n", player.GetIcon().Get())
					game.PrintBoard()
					break
				}
			}
		}

		if playAgain() {
			game.Reset()
		} else {
			playing = false
		}
	}
}

func playAgain() bool {
	fmt.Printf("Play again? Yes/No\n")
	input := ""
	fmt.Scanf("%s", &input)
	return input[0] == 'y' || input[0] == 'Y'
}
