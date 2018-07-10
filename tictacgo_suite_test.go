package tictacgo_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTictacgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tictacgo Suite")
}
