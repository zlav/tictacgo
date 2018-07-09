package cell

import "github.com/zlav/tictacgo/symbol"

type cell struct {
	value symbol.Symbol
}

func NewCell() cell {
	return cell{value: symbol.NewSymbol()}
}

func (c cell) Set(newVal string) {
	c.value.Set(newVal)
}

func (c cell) GetValue() symbol.Symbol {
	return c.value
}

func (c cell) Print() {
	c.value.Print()
}

func (c cell) Reset() {
	c.value.Reset()
}
