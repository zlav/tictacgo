package symbol

import "fmt"

type Symbol struct {
	char string
}

func NewSymbol(c string) Symbol {
	return Symbol{char: c}
}

// TODO: Correct error return
func (s *Symbol) Set(val string) {
	s.char = val
}

func (s Symbol) Get() string {
	return s.char
}

func (s Symbol) Print() {
	fmt.Print(s.char)
}
