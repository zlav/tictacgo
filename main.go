package main

import (
	"fmt"

	"github.com/zlav/tictacgo/board"
	"github.com/zlav/tictacgo/player"
)

func main() {
	fmt.Printf("Welcome to Tic Tac Go\n")

	var players []player.Player
	setup := false
	fmt.Printf("How would you like to play? 1 player, 2 player, or watch the computer fight itself?\n")
	for !setup {
		setup = true
		input := ""
		fmt.Scanf("%s", &input)
		switch input[0] {
		case '1':
			players = []player.Player{player.NewHuman("X"), player.NewComputer("O")}
		case '2':
			players = []player.Player{player.NewHuman("X"), player.NewHuman("O")}
		case 'c', 'C':
			players = []player.Player{player.NewComputer("X"), player.NewComputer("O")}
		default:
			fmt.Printf("Invalid Input\n")
			setup = false
		}
	}

	game := board.NewGame(players[0].GetIcon(), players[1].GetIcon())
	fmt.Printf("Here is an example of how to place a token on the grid\n")
	game.PrintHelp()
	fmt.Printf("Good Luck!\n")

	playing := true
	for playing {
		for !game.IsDraw() && !game.IsWon() {
			for _, player := range players {
				game.PrintBoard()
				player.PlayTicTacToe(game)
				if game.IsDraw() {
					break
				}
				if game.IsWon() {
					fmt.Printf("Congrats to %s!\n", player.GetName())
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
	fmt.Printf("Play again? Y/N\n")
	input := ""
	fmt.Scanf("%s", &input)
	return input[0] == 'y' || input[0] == 'Y'
}
