package player

import (
	"fmt"
	"time"

	"github.com/zlav/tictacgo/board"
	"github.com/zlav/tictacgo/symbol"
)

type Player interface {
	GetName() string
	GetIcon() symbol.Symbol
	PlayTicTacToe(*board.Tictacboard)
}

type human struct {
	name string
	icon symbol.Symbol
}

type computer struct {
	name string
	icon symbol.Symbol
}

func NewHuman(icon string) Player {
	return human{
		name: "Player",
		icon: symbol.NewSymbol(icon)}
}

func NewComputer(icon string) Player {
	return computer{
		name: "Player",
		icon: symbol.NewSymbol(icon)}
}

func (h human) PlayTicTacToe(board *board.Tictacboard) {
	fmt.Printf("Player %s's turn\n", h.icon.Get())
	validPlay := false
	for !validPlay {
		var input int
		fmt.Scanf("%d", &input)
		validPlay = board.Play(input, h.icon)
	}
}

func (c computer) PlayTicTacToe(board *board.Tictacboard) {
	fmt.Println("Computers Turn")
	time.Sleep(500 * time.Millisecond)
	board.Play(board.BestPlay(c.icon), c.icon)
}

func (h human) GetName() string {
	return h.name
}

func (c computer) GetName() string {
	return c.name
}

func (h human) GetIcon() symbol.Symbol {
	return h.icon
}

func (c computer) GetIcon() symbol.Symbol {
	return c.icon
}
