package player_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/zlav/tictacgo/player"
)

var _ = Describe("Player", func() {
	Context("human player is created" func(){
		humanIcon := "X"
		humanTest := player.NewHuman(humanIcon)
		It("initializes with the correct icon passed", func(){
			Expect(humanTest.GetIcon()).To(Equal(symbol.NewSymbol(humanIcon)
		})
	})
})
