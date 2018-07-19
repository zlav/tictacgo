package director_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/zlav/tictacgo/director"
)

var _ = Describe("Director", func() {
	Context("new tictactoe game created", func() {
		game := director.NewDirector()
		game.TurnOn()

		It("can setup a new game", func() {
			Expect(game.SetupGame("2")).To(BeTrue())
		})

		Context("playing on the game", func() {
			It("can place a token on the game", func() {
				Expect(game.Play("1")).To(BeTrue())
			})

			It("returns false if the play is out of bounds", func() {
				Expect(game.Play("-1")).To(BeFalse())
			})

			It("returns false if the cell is set", func() {
				Expect(game.Play("1")).To(BeFalse())
			})

			It("can reset the game", func() {
				game.Reset("Y")
				Expect(game.Play("1")).To(BeTrue())
			})

			It("Can see a win", func() {
				game.Play("1")
				game.Play("2")
				game.Play("4")
				game.Play("3")
				game.Play("7")
				Expect(game.GameOver()).To(BeTrue())
			})

			It("Can see a draw", func() {
				game.Reset("y")
				game.Play("1")
				game.Play("4")
				game.Play("2")
				game.Play("3")
				game.Play("7")
				game.Play("5")
				game.Play("6")
				game.Play("8")
				game.Play("9")
				Expect(game.GameOver()).To(BeTrue())
			})
		})
	})
})
