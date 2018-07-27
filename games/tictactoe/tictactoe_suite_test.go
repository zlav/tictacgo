package tictactoe_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTictactoe(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tictactoe Suite")
}
