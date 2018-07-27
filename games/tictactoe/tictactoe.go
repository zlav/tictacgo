package tictactoe

import (
	"fmt"
	"strconv"
	"time"

	"github.com/zlav/tictacgo/board"
	"github.com/zlav/tictacgo/player"
	"github.com/zlav/tictacgo/symbol"
)

const rows = 3
const columns = 3
const maxPlayers = 2

type Tictactoegame struct {
	moves      int
	winner     player.Player
	players    [maxPlayers]player.Player
	board      *board.Board
	currPlayer int
}

func NewTicTacToe() *Tictactoegame {
	return &Tictactoegame{
		winner:     player.Player{},
		moves:      0,
		currPlayer: 0,
		players:    [maxPlayers]player.Player{},
		board:      board.NewBoard(rows, columns),
	}
}

func (d Tictactoegame) TurnOn() string {
	response := "Welcome to Tic Tac Go\nHow would you like to play? [1] 1 player, [2] 2 player, or watch the [3] Computer fight itself? "
	d.board.Reset()
	return response
}

func (d *Tictactoegame) Setup(gameMode string) (bool, string) {

	xSym := symbol.NewSymbol("X")
	oSym := symbol.NewSymbol("O")
	switch gameMode {
	case "1":
		d.players = [maxPlayers]player.Player{player.NewPlayer(xSym, "Player 1", false), player.NewPlayer(oSym, "Computer", true)}
	case "2":
		d.players = [maxPlayers]player.Player{player.NewPlayer(xSym, "Player 1", false), player.NewPlayer(oSym, "Player 2", false)}
	case "c", "C", "3":
		d.players = [maxPlayers]player.Player{player.NewPlayer(xSym, "Computer 1", true), player.NewPlayer(oSym, "Computer 2", true)}
	default:
		return false, "Invalid Input: [1] 1 Player, [2] 2 Player, or [3] Computer: "
	}

	return true, "Good Luck!\n\n"
}

func (d *Tictactoegame) Play(input string) (bool, string) {
	location, err := integerValidation(input)
	if err != "" {
		return false, err
	}

	r := (location - 1) / rows
	c := (location - 1) % columns
	if !d.board.SetCell(r, c, d.players[d.currPlayer].GetIcon()) {
		return false, "Incorrect input: Cell is already taken: "
	}

	d.checkWinner(r, c)

	d.currPlayer++
	if d.currPlayer >= maxPlayers {
		d.currPlayer = 0
	}

	d.moves++
	return true, ""
}

func (d Tictactoegame) PlayerTurn() bool {
	return !d.players[d.currPlayer].IsComputer()
}

func (d *Tictactoegame) CurrentStatus() (bool, string) {
	if d.IsWon() {
		return true, "Congratulations to " + d.winner.GetName() + "\nWould you like to play again? Y/N: "
	} else if d.IsDraw() {
		return true, "Draw! Nobody is a winner\nWould you like to play again? Y/N: "
	} else if d.players[d.currPlayer].IsComputer() {
		response := fmt.Sprintf("%s's Turn\n", d.players[d.currPlayer].GetName())
		time.Sleep(500 * time.Millisecond)
		d.Play(strconv.Itoa(d.bestPlay(d.players[d.currPlayer].GetIcon())))
		return true, response
	}

	return false, d.players[d.currPlayer].GetName() + ": "
}

func (d *Tictactoegame) IsDraw() bool {
	return d.moves >= rows*columns
}

func (d *Tictactoegame) IsWon() bool {
	return d.winner.GetIcon() != symbol.Symbol{}
}

func (d *Tictactoegame) PrintGame() string {
	response := ""
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			response += fmt.Sprintf("|%s", d.board.GetCellValue(r, c).Get())
		}
		response += fmt.Sprint("|")
		if r < 2 {
			response += fmt.Sprintf("\n-------")
		}
		response += fmt.Sprintf("\n")
	}
	response += fmt.Sprintf("\n")
	return response
}

func (d *Tictactoegame) PrintHelp() string {
	response := fmt.Sprintf("Example: Select the numbered spot you wish to place your piece\n")
	count := 1
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			response += fmt.Sprintf("|%d", count)
			count++
		}
		response += fmt.Sprint("|")
		if r < 2 {
			response += fmt.Sprintf("\n-------")
		}
		response += "\n"
	}
	response += fmt.Sprintf("\n\n")
	return response
}

func (d *Tictactoegame) Reset() {
	d.winner = player.Player{}
	d.moves = 0
	d.currPlayer = 0
	d.board.Reset()
}

func (d *Tictactoegame) bestPlay(s symbol.Symbol) int {
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

func integerValidation(input string) (int, string) {
	location, err := strconv.Atoi(input)
	if err != nil {
		return location, fmt.Sprintf("Incorrect input: Please enter an Integer: ")
	}
	if location > rows*columns || location < 1 {
		return location, fmt.Sprintf("Incorrect input: Please enter an Integer between 1 and %d: ", rows*columns)

	}
	return location, ""
}

func (d *Tictactoegame) checkWinner(r, c int) {
	player := d.players[d.currPlayer]
	for i := 0; i < columns; i++ {
		if (i != c) && d.board.GetCellValue(r, i) != player.GetIcon() {
			break
		}
		if i == columns-1 {
			d.winner = player
			return
		}
	}

	for i := 0; i < rows; i++ {
		if (i != r) && d.board.GetCellValue(i, c) != player.GetIcon() {
			break
		}
		if i == rows-1 {
			d.winner = player
			return
		}
	}

	if r == c {
		for i := 0; i < rows; i++ {
			if (i != r) && d.board.GetCellValue(i, i) != player.GetIcon() {
				break
			}
			if i == rows-1 {
				d.winner = player
				return
			}
		}
	}

	if r+c == rows-1 {
		for i := 0; i < rows; i++ {
			if (i != r) && d.board.GetCellValue(i, (rows-i)-1) != player.GetIcon() {
				break
			}
			if i == rows-1 {
				d.winner = player
				return
			}
		}
	}
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
