package symbol

type Symbol struct {
	char string
}

func NewSymbol(c string) Symbol {
	return Symbol{char: c}
}

func (s *Symbol) Set(val string) {
	s.char = val
}

func (s Symbol) Get() string {
	return s.char
}
