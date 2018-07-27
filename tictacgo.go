package main

import (
	"github.com/zlav/tictacgo/client"
	"github.com/zlav/tictacgo/director"
)

func main() {
	gameServer := director.NewDirector()
	ticTacClient := client.NewTicTacClient(gameServer)
	ticTacClient.Run()
}
