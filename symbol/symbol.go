package symbol

import "fmt"

type Symbol struct {
	char string
}

func NewSymbol() Symbol {
	return Symbol{char: " "}
}

func (s Symbol) Reset() {
	s.char = " "
}

// TODO: Correct error return
func (s Symbol) Set(val string) bool {
	if len(val) == 1 {
		s.char = val
		return true
	}
	return false
}

func (s Symbol) Get() string {
	return s.char
}

func (s Symbol) Print() {
	fmt.Print(s.char)
}
