package cell

import (
	"github.com/zlav/tictacgo/symbol"
)

const defaultSymbol = " "

type Cell struct {
	value symbol.Symbol
}

func NewCell() Cell {
	return Cell{value: symbol.NewSymbol(defaultSymbol)}
}

func (c *Cell) Set(newSym symbol.Symbol) bool {
	if c.value.Get() != defaultSymbol {
		return false
	}
	c.value.Set(newSym.Get())
	return true
}

func (c Cell) GetValue() symbol.Symbol {
	return c.value
}

func (c Cell) IsSet() bool {
	if c.value.Get() == defaultSymbol {
		return false
	}
	return true
}

func (c *Cell) Reset() {
	c.value.Set(defaultSymbol)
}
