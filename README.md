# Tic-Tac-Go 
This repo contains a tic tac toe application built using golang that features an unbeatable computer using the minimax algorithm 

## Notes
This is a server client implementation where the 'client' is in charge of sending 'requests' to the server. In this implementation the server and client are connected by functions and objects however the concept remains the same.
The server sends 'responses' to the client which can be printed or ignored. This is setup to be able to turn into an HTTP response in an actual network client server implementation.

## Install the application
### Install Go
Ensure that you have properly installed go onto your computer following the guidelines from https://golang.org/doc/code.html

### Install the repo
From anywhere on your command line inside the $GOPATH 
`go get github.com/zlav/tictacgo`
This will place all of the packages inside of your go workspace under the src folder

## Test the application
### Install packages
From anywhere on your command line run inside the $GOPATH
`go get github.com/onsi/ginkgo/ginkgo`
`go get github.com/onsi/gomega/...`
This will download the necessary testing framework and assertion library

### Run the tests
Navigate to $GOPATH/src/github.com/zlav/tictacgo
Run the following command and all tests will run, including the other packages
`ginkgo -r` (-r for recursive)
If you want to only run only local test files simply enter
`ginkgo`

Note: tictacgo.go is a driver for the other packages and user input so there is limited test coverage here

## Run the application
Navigate to $GOPATH/src/github.com/zlav/tictacgo
From here you have a few options
1. `go run tictacgo.go`
2. `go build tictacgo.go` `./tictacgo`
3. Move the built tictacgo file in your $GOPATH/bin to run from anywhere in your $GOPATH
4. Move to your local /bin directory to run from anywhere

## Next Development Steps
1. Build out more integration level tests and test user input
2. Add a better graphical interface
3. Drive out better features 
  * Take advantage of symbol package more
  * Offer better player customization
4. Push board and cell to their limits by adding an entirely different game such as checkers
