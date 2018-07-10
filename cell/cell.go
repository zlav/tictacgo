package cell

import (
	"github.com/zlav/tictacgo/symbol"
)

const DefaultState = " "

type Cell struct {
	value symbol.Symbol
}

func NewCell() Cell {
	return Cell{value: symbol.NewSymbol(DefaultState)}
}

func (c *Cell) Set(newSym symbol.Symbol) bool {
	if c.value.Get() != DefaultState {
		return false
	}
	c.value.Set(newSym.Get())
	return true
}

func (c Cell) GetValue() symbol.Symbol {
	return c.value
}

func (c Cell) IsSet() bool {
	if c.value.Get() == DefaultState {
		return false
	}
	return true
}

func (c *Cell) Reset() {
	c.value.Set(DefaultState)
}
