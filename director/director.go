package director

import (
	"fmt"

	"github.com/zlav/tictacgo/games"
)

type Director struct {
	game games.Tictactoegame
}

func NewDirector() *Director {
	return &Director{
		game: *games.NewTicTacToe(),
	}
}

func (d Director) TurnOn() {
	d.game.TurnOn()
}

func (d *Director) SetupGame(input string) bool {
	return d.game.Setup(input)
}

func (d *Director) Play(input string) bool {
	return d.game.Play(input)
}

func (d *Director) PlayerTurn() bool {
	return d.game.PlayerTurn()
}

func (d *Director) Status() bool {
	return d.game.CurrentStatus()
}

func (d *Director) PrintGame() {
	d.game.PrintGame()
}

func (d *Director) PrintHelp() {
	d.game.PrintHelp()
}

func (d Director) GameOver() bool {
	if d.game.IsDraw() || d.game.IsWon() {
		return true
	}
	return false
}

func (d *Director) Reset(input string) bool {
	if input[0] == 'y' || input[0] == 'Y' {
		d.game.Reset()
		return true
	}
	fmt.Printf("Thank you for playing!\n")
	return false
}
