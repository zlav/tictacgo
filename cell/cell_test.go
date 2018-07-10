package cell_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zlav/tictacgo/cell"
	"github.com/zlav/tictacgo/symbol"
)

var _ = Describe("Cell", func() {
	Context("when a cell is created", func() {
		testCell := cell.NewCell()

		It("initializes to a default value", func() {
			Expect(testCell.GetValue()).To(Equal(symbol.NewSymbol(cell.DefaultState)))
		})

		It("can set a new value and return it", func() {
			setVal := symbol.NewSymbol("T")
			testCell.Set(setVal)
			Expect(testCell.GetValue()).To(Equal(setVal))
		})

		It("can reset itself to the default state", func() {
			testCell.Reset()
			Expect(testCell.GetValue()).To(Equal(symbol.NewSymbol(cell.DefaultState)))
		})

		It("accurately returns if it is set or not", func() {
			testCell.Reset()
			Expect(testCell.IsSet()).To(BeFalse())

			testCell.Set(symbol.NewSymbol("K"))
			Expect(testCell.IsSet()).To(BeTrue())
		})
	})
})
