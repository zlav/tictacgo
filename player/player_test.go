package player_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/zlav/tictacgo/player"
	"github.com/zlav/tictacgo/symbol"
)

var _ = Describe("Player", func() {
	testIcon := symbol.NewSymbol("X")

	Context("non computer player is created", func() {
		playerTest := player.NewPlayer(testIcon, "Player 1", false)

		It("initializes with the correct icon passed", func() {
			Expect(playerTest.GetIcon()).To(Equal(testIcon))
		})

		It("returns the correct default name", func() {
			Expect(playerTest.GetName()).To(Equal("Player 1"))
		})

		It("returns the correct computer setting", func() {
			Expect(playerTest.IsComputer()).To(Equal(false))
		})
	})

	Context("computer player is created", func() {
		playerTest := player.NewPlayer(testIcon, "Player 1", true)

		It("returns the correct computer setting", func() {
			Expect(playerTest.IsComputer()).To(Equal(true))
		})
	})
})
