package cell_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCell(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cell Suite")
}
