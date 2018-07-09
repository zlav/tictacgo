package board

import (
	"fmt"

	"github.com/zlav/tictacgo/cell"
	"github.com/zlav/tictacgo/symbol"
)

const rows = 3
const columns = 3

type Tictacboard struct {
	grid   [rows][columns]cell.Cell
	winner symbol.Symbol
	moves  int
}

func NewGame() *Tictacboard {
	return &Tictacboard{
		winner: symbol.Symbol{},
		moves:  0,
		grid: [rows][columns]cell.Cell{{cell.NewCell(), cell.NewCell(), cell.NewCell()},
			{cell.NewCell(), cell.NewCell(), cell.NewCell()},
			{cell.NewCell(), cell.NewCell(), cell.NewCell()}}}
}

func (board Tictacboard) PrintBoard() {
	for r, row := range board.grid {
		for _, box := range row {
			fmt.Print("|")
			box.Print()
		}
		fmt.Print("|")
		if r < 2 {
			fmt.Print("\n-------")
		}
		fmt.Println()
	}
}

func (board Tictacboard) PrintHelp() {
	count := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Printf("|%d", count)
			count += 1
		}
		fmt.Print("|")
		if i < 2 {
			fmt.Print("\n-------")
		}
		fmt.Println()
	}
}

func (board *Tictacboard) Play(location int, s symbol.Symbol) bool {
	r := (location - 1) / 3
	c := (location - 1) % 3
	if r < 0 || r >= rows || c < 0 || c >= columns {
		fmt.Println("Input must be between 1-9")
		return false
	}
	newPlay := board.grid[r][c].Set(s)
	if !newPlay {
		fmt.Println("Cell already taken")
		return false
	}
	board.moves += 1
	board.updateStatus(r, c, s)
	return newPlay
}

func (board Tictacboard) BestPlay(s symbol.Symbol) int {

	return board.moves
}

func (board Tictacboard) IsWon() bool {
	return board.winner != symbol.Symbol{}
}

func (board Tictacboard) IsDraw() bool {
	return board.moves >= rows*columns
}

func (board *Tictacboard) Reset() {
	board.winner = symbol.Symbol{}
	board.moves = 0
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			board.grid[r][c].Reset()
		}
	}
}

func (board *Tictacboard) updateStatus(r int, c int, s symbol.Symbol) {
	for i := 0; i < columns; i++ {
		if (i != c) && board.grid[r][i].GetValue() != s {
			break
		}
		if i == columns-1 {
			board.winner = s
			return
		}
	}

	for i := 0; i < rows; i++ {
		if (i != r) && board.grid[i][c].GetValue() != s {
			break
		}
		if i == rows-1 {
			board.winner = s
			return
		}
	}

	if r == c {
		for i := 0; i < rows; i++ {
			if (i != r) && board.grid[i][i].GetValue() != s {
				break
			}
			if i == rows-1 {
				board.winner = s
				return
			}
		}
	}

	if r+c == rows-1 {
		for i := 0; i < rows; i++ {
			if (i != r) && board.grid[i][(rows-i)-1].GetValue() != s {
				break
			}
			if i == rows-1 {
				board.winner = s
				return
			}
		}
	}
}
