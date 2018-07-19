package games_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGames(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Games Suite")
}
