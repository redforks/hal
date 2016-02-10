package timeth

import (
	"github.com/onsi/ginkgo"

	"testing"
)

var t = ginkgo.GinkgoT

func TestTimeth(t *testing.T) {
	ginkgo.RunSpecs(t, "Timeth Suite")
}
