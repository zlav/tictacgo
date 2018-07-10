package board_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/zlav/tictacgo/board"
	"github.com/zlav/tictacgo/symbol"
)

var _ = Describe("Board", func() {
	player1 := symbol.NewSymbol("X")
	player2 := symbol.NewSymbol("O")
	Context("new tictactoe board created", func() {
		board := board.NewGame(player1, player2)

		Context("playing on the board", func() {
			It("can place a token on the board", func() {
				Expect(board.Play(1, player1)).To(BeTrue())
				Expect(board.GetCellValue(1)).To(Equal(player1))
			})

			It("returns false if the play is out of bounds", func() {
				Expect(board.Play(-1, player1)).To(BeFalse())
			})

			It("returns false if the cell is set", func() {
				Expect(board.Play(1, player1)).To(BeFalse())
			})

			It("can reset the board", func() {
				board.Reset()
				Expect(board.GetCellValue(1)).ToNot(Equal(player1))
			})

		})

		Context("the board is intelligent", func() {
			BeforeEach(func() {
				board.Reset()
			})
			It("knows the next best play on the board", func() {
				board.Play(1, player1)
				board.Play(3, player2)
				board.Play(5, player1)
				Expect(board.BestPlay(player2)).To(Equal(9))
			})
			It("knows if the board is a winner", func() {
				makeWin(board, player1)
				Expect(board.IsWon()).To(BeTrue())
			})
			It("knows if the board is a draw", func() {
				makeDraw(board, player1, player2)
				Expect(board.IsDraw()).To(BeTrue())
			})
		})
	})

})

func makeWin(b *board.Tictacboard, p1 symbol.Symbol) {
	b.Play(1, p1)
	b.Play(2, p1)
	b.Play(3, p1)
}

func makeDraw(b *board.Tictacboard, p1, p2 symbol.Symbol) {
	b.Play(1, p1)
	b.Play(2, p1)
	b.Play(3, p2)
	b.Play(4, p2)
	b.Play(5, p2)
	b.Play(6, p1)
	b.Play(7, p1)
	b.Play(8, p1)
	b.Play(9, p2)
}
