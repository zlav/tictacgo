package player

import (
	"github.com/zlav/tictacgo/symbol"
)

type Player struct {
	name     string
	icon     symbol.Symbol
	computer bool
}

func NewPlayer(newSym symbol.Symbol, name string, computer bool) Player {
	return Player{
		name:     name,
		icon:     newSym,
		computer: computer,
	}
}

func (p Player) GetName() string {
	return p.name
}

func (p Player) GetIcon() symbol.Symbol {
	return p.icon
}

func (p Player) IsComputer() bool {
	return p.computer
}
