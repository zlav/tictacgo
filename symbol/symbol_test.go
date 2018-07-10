package symbol_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zlav/tictacgo/symbol"
)

var _ = Describe("Symbol", func() {
	Context("When a symbol is created", func() {
		initialValue := "T"
		sym := symbol.NewSymbol(initialValue)

		It("returns the correct value", func() {
			Expect(sym.Get()).To(Equal(initialValue))
		})

		It("can change its value", func() {
			newValue := "N"
			sym.Set(newValue)
			Expect(sym.Get()).To(Equal(newValue))
		})
	})
})
