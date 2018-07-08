package main

type symbol struct {
	char string
}

func NewSymbol() symbol {
	return symbol{char: " "}
}

func (s symbol) ResetSymbol() {
	s.char = " "
}

// TODO: Correct error return
func (s symbol) SetSymbol(val string) bool {
	if len(val) == 1 {
		s.char = val
		return true
	}
	return false
}
