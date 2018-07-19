package client_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/zlav/tictacgo/client"
	"github.com/zlav/tictacgo/director"
)

var _ = Describe("Client", func() {
	Context("new client created", func() {
		game := director.NewDirector()
		game.TurnOn()
		game.SetupGame("2")

		It("successfully creates the client", func() {
			client := client.NewTicTacClient(game)
			Expect(client).ShouldNot(BeNil())
		})
	})
})
