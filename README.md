# Tic-Tac-Go 
This repo contains a tic tac toe application built using golang that features an unbeatable computer using the minimax algorithm 

## Install the application
### Install Go
Ensure that you have properly installed go onto your computer following the guidelines from https://golang.org/doc/code.html

### Install the repo
From anywhere on your command line run 
`go get github.com/zlav/tictacgo`
This will place all of the packages inside of your go workspace under the src folder

## Run the application
Navigate to $GOPATH/src/github.com/zlav/tictacgo
From here you have a few options
1. `go run tictacgo.go`
2. `go build tictacgo.go` `./tictacgo`
3. Move the built tictacgo file in your $GOPATH/bin to run from anywhere in your $GOPATH
4. Move to your local /bin directory to run from anywhere

## Next Development Steps
1. Build out a full test suite
2. Add a better graphical interface
3. Drive out better features 
  * Take advantage of symbol package more
  * Offer better player customization
4. Push board and cell to their limits by adding an entirely different game such as checkers
