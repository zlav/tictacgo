package board

import (
	"fmt"

	"github.com/zlav/tictacgo/cell"
	"github.com/zlav/tictacgo/symbol"
)

const rows = 3
const columns = 3
const maxPlayers = 2

type Tictacboard struct {
	grid    [rows][columns]cell.Cell
	winner  symbol.Symbol
	players [maxPlayers]symbol.Symbol
	moves   int
}

func NewGame(one symbol.Symbol, two symbol.Symbol) *Tictacboard {
	return &Tictacboard{
		winner:  symbol.Symbol{},
		moves:   0,
		players: [maxPlayers]symbol.Symbol{one, two},
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

func (board *Tictacboard) BestPlay(s symbol.Symbol) int {
	bestVal := -1000
	bestMove := -1

	opponent := board.players[0]
	if board.players[0].Get() == s.Get() {
		opponent = board.players[1]
	}

	for i := 0; i < rows*columns; i++ {
		r := i / 3
		c := i % 3
		if !board.grid[r][c].IsSet() {
			board.grid[r][c].Set(s)
			board.moves += 1

			moveVal := minimax(board, 0, false, s, opponent)

			board.moves -= 1
			board.grid[r][c].Reset()

			if moveVal > bestVal {
				bestMove = i
				fmt.Printf("Best Move: %d and score %d\n", bestMove, moveVal)
				bestVal = moveVal
			}
		}
	}
	fmt.Printf("Best Move is %d\n", bestMove)

	return bestMove + 1
}

func evaluate(board *Tictacboard, player symbol.Symbol, opponent symbol.Symbol) int {
	for row := 0; row < 3; row++ {
		if board.grid[row][0].GetValue().Get() == board.grid[row][1].GetValue().Get() && board.grid[row][1].GetValue() == board.grid[row][2].GetValue() {
			if board.grid[row][0].GetValue().Get() == player.Get() {
				return +10
			} else if board.grid[row][0].GetValue().Get() == opponent.Get() {
				return -10
			}
		}
	}

	for col := 0; col < 3; col++ {
		if board.grid[0][col].GetValue().Get() == board.grid[1][col].GetValue().Get() && board.grid[1][col].GetValue().Get() == board.grid[2][col].GetValue().Get() {
			if board.grid[0][col].GetValue().Get() == player.Get() {
				return +10
			} else if board.grid[0][col].GetValue().Get() == opponent.Get() {
				return -10
			}
		}
	}

	if board.grid[0][0].GetValue().Get() == board.grid[1][1].GetValue().Get() && board.grid[1][1].GetValue().Get() == board.grid[2][2].GetValue().Get() {
		if board.grid[0][0].GetValue().Get() == player.Get() {
			return +10
		} else if board.grid[0][0].GetValue().Get() == opponent.Get() {
			return -10
		}
	}

	if board.grid[0][2].GetValue().Get() == board.grid[1][1].GetValue().Get() && board.grid[1][1].GetValue().Get() == board.grid[2][0].GetValue().Get() {
		if board.grid[0][2].GetValue().Get() == player.Get() {
			return +10
		} else if board.grid[0][2].GetValue().Get() == opponent.Get() {
			return -10
		}
	}

	return 0
}

func minimax(board *Tictacboard, depth int, isMax bool, player symbol.Symbol, opponent symbol.Symbol) int {
	score := evaluate(board, player, opponent)

	if score == 10 {
		return score
	}

	if score == -10 {
		return score
	}

	if board.moves == 9 {
		return 0
	}

	if isMax {
		best := -1000

		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if !board.grid[i][j].IsSet() {

					board.grid[i][j].Set(player)
					board.moves += 1

					temp := minimax(board, depth+1, !isMax, player, opponent)
					if temp > best {
						best = temp
					}

					board.moves -= 1
					board.grid[i][j].Reset()
				}
			}
		}
		return best
	} else {
		best := 1000

		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if !board.grid[i][j].IsSet() {

					board.grid[i][j].Set(opponent)
					board.moves += 1

					temp := minimax(board, depth+1, !isMax, player, opponent)
					if temp < best {
						best = temp
					}

					board.moves -= 1
					board.grid[i][j].Reset()
				}
			}
		}
		return best
	}
}

func (board *Tictacboard) updateStatus(r int, c int, s symbol.Symbol) {
	for i := 0; i < columns; i++ {
		if (i != c) && board.grid[r][i].GetValue().Get() != s.Get() {
			break
		}
		if i == columns-1 {
			board.winner = s
			return
		}
	}

	for i := 0; i < rows; i++ {
		if (i != r) && board.grid[i][c].GetValue().Get() != s.Get() {
			break
		}
		if i == rows-1 {
			board.winner = s
			return
		}
	}

	if r == c {
		for i := 0; i < rows; i++ {
			if (i != r) && board.grid[i][i].GetValue().Get() != s.Get() {
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
			if (i != r) && board.grid[i][(rows-i)-1].GetValue().Get() != s.Get() {
				break
			}
			if i == rows-1 {
				board.winner = s
				return
			}
		}
	}
}
