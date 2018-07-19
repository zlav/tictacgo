package board_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/zlav/tictacgo/board"
	"github.com/zlav/tictacgo/symbol"
)

var _ = Describe("Board", func() {
	player1 := symbol.NewSymbol("X")

	Context("new tictactoe board created", func() {
		board := board.NewBoard(3, 3)
		board.Reset()

		Context("playing on the board", func() {
			It("can place a token on the board", func() {
				Expect(board.SetCell(1, 1, player1)).To(BeTrue())
			})

			It("can retrieve a placed token on the board", func() {
				Expect(board.GetCellValue(1, 1)).To(Equal(player1))
			})

			It("returns false if the play is out of bounds", func() {
				Expect(board.SetCell(-1, -1, player1)).To(BeFalse())
			})

			It("it returns false if the cell is set and you try to set", func() {
				Expect(board.SetCell(1, 1, player1)).To(BeFalse())
			})

			It("it returns true if the cell is set and check", func() {
				Expect(board.IsCellSet(1, 1)).To(BeTrue())
			})

			It("can reset the board", func() {
				board.Reset()
				Expect(board.GetCellValue(1, 1)).ToNot(Equal(player1))
			})
		})
	})
})
