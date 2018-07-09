package player

import "fmt"

type Player interface {
	PlayTicTacToe()
}

type human struct {
	name string
}

type computer struct {
	name string
}

func NewHuman() Player {
	return human{name: "Player"}
}

func NewComputer() Player {
	return computer{name: "Computer"}
}

func (h human) PlayTicTacToe() {
	fmt.Println("Play Human")
}

func (c computer) PlayTicTacToe() {
	fmt.Println("Play Player")
}
