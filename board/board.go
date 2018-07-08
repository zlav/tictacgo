package board

import "fmt"

type tictacboard struct {
	board [][]string
}

func NewGame() tictacboard {
	var b tictacboard
	b.board = [][]string{{"-", "-", "-"}, {"-", "-", "-"}, {"-", "-", "-"}}
	return b
}

// Helper option if help mode is turned on that prints all the current openings next to the board
func (board tictacboard) PrintBoard() {
	for _, row := range board.board {
		for _, box := range row {
			fmt.Printf("%s ", box)
		}
		fmt.Println()
	}
}

// TODO: Everything
func (board tictacboard) Play(x int) bool {

	return true
}

// TODO: Everything
func (board tictacboard) BestPlay() int {

	return 1
}
