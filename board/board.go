package board

import (
	"github.com/zlav/tictacgo/cell"
	"github.com/zlav/tictacgo/symbol"
)

type Board struct {
	grid [][]cell.Cell
}

func NewBoard(r, c int) *Board {
	temp := make([][]cell.Cell, r)
	for i := range temp {
		temp[i] = make([]cell.Cell, c)
	}
	return &Board{
		grid: temp,
	}
}

func (board *Board) SetCell(r, c int, s symbol.Symbol) bool {
	if board.outOfBounds(r, c) {
		return false
	}
	return board.grid[r][c].Set(s)
}

func (board Board) GetCellValue(r, c int) symbol.Symbol {
	if board.outOfBounds(r, c) {
		return symbol.Symbol{}
	}
	return board.grid[r][c].GetValue()
}

func (board Board) IsCellSet(r, c int) bool {
	if board.outOfBounds(r, c) {
		return false
	}
	return board.grid[r][c].IsSet()
}

func (board *Board) ResetCell(r, c int) bool {
	if board.outOfBounds(r, c) {
		return false
	}
	board.grid[r][c].Reset()
	return true
}

func (board *Board) Reset() {
	for r := 0; r < len(board.grid); r++ {
		for c := 0; c < len(board.grid[0]); c++ {
			board.grid[r][c].Reset()
		}
	}
}

func (board *Board) outOfBounds(r, c int) bool {
	if r >= len(board.grid) || r < 0 || c < 0 || c >= len(board.grid[0]) {
		return true
	}
	return false
}
