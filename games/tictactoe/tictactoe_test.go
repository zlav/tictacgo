package tictactoe_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zlav/tictacgo/games/tictactoe"
)

var _ = Describe("Director", func() {
	Context("new player vs computer tictactoe game board created", func() {
		testGame := tictactoe.NewTicTacToe()
		testGame.TurnOn()
		var expect bool

		It("can setup a new game", func() {
			expect, _ = testGame.Setup("1")
			Expect(expect).To(BeTrue())
		})

		Context("when playing pieces on the board", func() {
			It("knows when it is the players turn", func() {
				Expect(testGame.PlayerTurn()).To(BeTrue())
			})

			It("can play a token on the board", func() {
				expect, _ = testGame.Play("1")
				Expect(expect).To(BeTrue())
			})

			It("knows when it is not the players turn", func() {
				Expect(testGame.PlayerTurn()).To(BeFalse())
			})

			It("cannot play a token on the board where it has already been played", func() {
				expect, _ = testGame.Play("1")
				Expect(expect).To(BeFalse())
			})

			It("cannot play a token on the board out of bounds", func() {
				expect, _ = testGame.Play("-1")
				Expect(expect).To(BeFalse())
			})
		})

		Context("when the computer is playing", func() {
			It("knows the best place to play", func() {
				testGame.Reset()
				expect, _ = testGame.Play("1")
				Expect(expect).To(BeTrue())
				expect, _ = testGame.CurrentStatus()
				Expect(expect).To(BeTrue())
				expect, _ = testGame.Play("5")
				Expect(expect).To(BeFalse())
			})
		})
	})

	Context("new player vs player tictactoe game board created", func() {
		testGame := tictactoe.NewTicTacToe()
		testGame.Setup("2")
		testGame.TurnOn()

		It("can detect a draw", func() {
			testGame.Play("1")
			testGame.Play("2")
			testGame.Play("4")
			testGame.Play("3")
			testGame.Play("7")
			Expect(testGame.IsWon()).To(BeTrue())
		})

		It("can detect a win", func() {
			testGame.Reset()
			testGame.Play("1")
			testGame.Play("4")
			testGame.Play("2")
			testGame.Play("3")
			testGame.Play("7")
			testGame.Play("5")
			testGame.Play("6")
			testGame.Play("8")
			testGame.Play("9")
			Expect(testGame.IsDraw()).To(BeTrue())
		})
	})
})
