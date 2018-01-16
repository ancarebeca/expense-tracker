package etl_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestEtl(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Etl Suite")
}
