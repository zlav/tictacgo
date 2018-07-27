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
	d.board.Reset()
	return "Welcome to Tic Tac Go\nHow would you like to play? [1] 1 player, [2] 2 player, or watch the [3] Computer fight itself? "
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

func (d *Tictactoegame) CurrentStatus() (bool, string) {
	if d.IsWon() {
		return true, "Congratulations to " + d.winner.GetName() + "\nWould you like to play again? Y/N: "
	} else if d.IsDraw() {
		return true, "Draw! Nobody is a winner\nWould you like to play again? Y/N: "
	} else if d.players[d.currPlayer].IsComputer() {
		response := fmt.Sprintf("%s's Turn\n", d.players[d.currPlayer].GetName())
		time.Sleep(500 * time.Millisecond)
		d.Play(strconv.Itoa(bestPlay(d, d.players[d.currPlayer].GetIcon())))
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

func (d *Tictactoegame) Reset() {
	d.winner = player.Player{}
	d.moves = 0
	d.currPlayer = 0
	d.board.Reset()
}

func (d Tictactoegame) PlayerTurn() bool {
	return !d.players[d.currPlayer].IsComputer()
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
