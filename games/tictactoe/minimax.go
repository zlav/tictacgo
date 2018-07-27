package tictactoe

import (
	"github.com/zlav/tictacgo/board"
	"github.com/zlav/tictacgo/symbol"
)

func bestPlay(d *Tictactoegame, s symbol.Symbol) int {
	bestVal := -1000
	bestMove := -1

	opponent := d.players[0]
	if d.players[0].GetIcon() == s {
		opponent = d.players[1]
	}

	for i := 0; i < rows*columns; i++ {
		r := i / 3
		c := i % 3
		if !d.board.IsCellSet(r, c) {
			d.board.SetCell(r, c, s)
			d.moves++

			moveVal := minimax(d, 0, false, s, opponent.GetIcon())
			d.moves--
			d.board.ResetCell(r, c)

			if moveVal > bestVal {
				bestMove = i
				bestVal = moveVal
			}
		}
	}

	return bestMove + 1
}

func minimax(d *Tictactoegame, depth int, isMax bool, player symbol.Symbol, opponent symbol.Symbol) int {
	score := evaluateGrid(d.board, player, opponent)

	if score == 10 {
		return score
	}

	if score == -10 {
		return score
	}

	if d.moves == 9 {
		return 0
	}

	if isMax {
		best := -1000

		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if !d.board.IsCellSet(i, j) {

					d.board.SetCell(i, j, player)
					d.moves++

					temp := minimax(d, depth+1, !isMax, player, opponent)
					if temp > best {
						best = temp
					}

					d.moves--
					d.board.ResetCell(i, j)
				}
			}
		}
		return best
	}

	best := 1000
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !d.board.IsCellSet(i, j) {

				d.board.SetCell(i, j, opponent)
				d.moves++

				temp := minimax(d, depth+1, !isMax, player, opponent)
				if temp < best {
					best = temp
				}

				d.moves--
				d.board.ResetCell(i, j)
			}
		}
	}
	return best
}

func evaluateGrid(board *board.Board, player symbol.Symbol, opponent symbol.Symbol) int {
	for row := 0; row < 3; row++ {
		if board.GetCellValue(row, 0) == board.GetCellValue(row, 1) && board.GetCellValue(row, 1) == board.GetCellValue(row, 2) {
			if board.GetCellValue(row, 0) == player {
				return +10
			} else if board.GetCellValue(row, 0) == opponent {
				return -10
			}
		}
	}

	for col := 0; col < 3; col++ {
		if board.GetCellValue(0, col) == board.GetCellValue(1, col) && board.GetCellValue(1, col) == board.GetCellValue(2, col) {
			if board.GetCellValue(0, col) == player {
				return +10
			} else if board.GetCellValue(0, col) == opponent {
				return -10
			}
		}
	}

	if board.GetCellValue(0, 0) == board.GetCellValue(1, 1) && board.GetCellValue(1, 1) == board.GetCellValue(2, 2) {
		if board.GetCellValue(0, 0) == player {
			return +10
		} else if board.GetCellValue(0, 0) == opponent {
			return -10
		}
	}

	if board.GetCellValue(0, 2) == board.GetCellValue(1, 1) && board.GetCellValue(1, 1) == board.GetCellValue(2, 0) {
		if board.GetCellValue(0, 2) == player {
			return +10
		} else if board.GetCellValue(0, 2) == opponent {
			return -10
		}
	}

	return 0
}
