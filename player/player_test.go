package player_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/zlav/tictacgo/board"
	"github.com/zlav/tictacgo/player"
	"github.com/zlav/tictacgo/symbol"
)

var _ = Describe("Player", func() {
	humanIcon := symbol.NewSymbol("X")
	computerIcon := symbol.NewSymbol("O")
	gameBoard := board.NewGame(humanIcon, computerIcon)

	AfterEach(func() {
		gameBoard.Reset()
	})

	Context("human player is created", func() {
		humanTest := player.NewHuman(humanIcon)

		It("initializes with the correct icon passed", func() {
			Expect(humanTest.GetIcon()).To(Equal(humanIcon))
		})

		It("returns the correct default name", func() {
			Expect(humanTest.GetName()).To(Equal("Player"))
		})
	})

	Context("computer player is created", func() {
		computerTest := player.NewComputer(computerIcon)

		It("initializes with the correct icon passed", func() {
			Expect(computerTest.GetIcon()).To(Equal(computerIcon))
		})

		It("returns the correct default name", func() {
			Expect(computerTest.GetName()).To(Equal("Computer"))
		})

		It("plays correctly", func() {
			gameBoard.Play(1, humanIcon)
			gameBoard.Play(3, computerIcon)
			gameBoard.Play(5, humanIcon)
			computerTest.PlayTicTacToe(gameBoard)
			Expect(gameBoard.GetCellValue(9)).To(Equal(computerIcon))
		})
	})
})
