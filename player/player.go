package player

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func NewHuman(newSym symbol.Symbol) Player {
	return human{
		name: "Player",
		icon: newSym,
	}
}

func NewComputer(newSym symbol.Symbol) Player {
	return computer{
		name: "Computer",
		icon: newSym,
	}
}

func (h human) PlayTicTacToe(board *board.Tictacboard) {
	fmt.Printf("%s %s's turn\n", h.name, h.icon.Get())
	validPlay := false

	for !validPlay {
		reader := bufio.NewReader(os.Stdin)
		input, tooLong, err := reader.ReadLine()
		if !tooLong || err != nil {
			result, err := strconv.Atoi(string(input))
			if err != nil {
				// TODO: Error message from board about the numebers allowed
				fmt.Println("Incorrect input: Please enter the correct number")
			} else {
				validPlay = board.Play(result, h.icon)
			}
		}
	}
}

func (c computer) PlayTicTacToe(board *board.Tictacboard) {
	fmt.Printf("%s's Turn\n", c.name)
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
