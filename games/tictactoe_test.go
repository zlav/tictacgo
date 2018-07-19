package games_test

import "github.com/zlav/tictacgo/games"
import . "github.com/onsi/ginkgo"
import . "github.com/onsi/gomega"

var _ = Describe("Director", func() {
	Context("new player vs computer tictactoe game board created", func() {
		testGame := games.NewTicTacToe()
		testGame.TurnOn()

		It("can setup a new game", func() {
			Expect(testGame.Setup("1")).To(BeTrue())
		})

		Context("when playing pieces on the board", func() {
			It("knows when it is the players turn", func() {
				Expect(testGame.PlayerTurn()).To(BeTrue())
			})

			It("can play a token on the board", func() {
				Expect(testGame.Play("1")).To(BeTrue())
			})

			It("knows when it is not the players turn", func() {
				Expect(testGame.PlayerTurn()).To(BeFalse())
			})

			It("cannot play a token on the board where it has already been played", func() {
				Expect(testGame.Play("1")).To(BeFalse())
			})

			It("cannot play a token on the board out of bounds", func() {
				Expect(testGame.Play("-1")).To(BeFalse())
			})
		})
		Context("when the computer is playing", func() {
			It("knows the best place to play", func() {
				testGame.Reset()
				Expect(testGame.Play("1")).To(BeTrue())
				Expect(testGame.CurrentStatus()).To(BeTrue())
				Expect(testGame.Play("5")).To(BeFalse())
			})
		})
	})

	Context("new player vs player tictactoe game board created", func() {
		testGame := games.NewTicTacToe()
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
