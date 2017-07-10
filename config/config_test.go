package config

import (
	. "github.com/onsi/gomega"
	"testing"
)

func Test_WhenConfigFileIsGiven_ItLoadsTheConfig(t *testing.T) {
	RegisterTestingT(t)

	path := "../fixtures/configTest.yml"
	cfg := Conf{}

	err := cfg.LoadConfig(path)
	Expect(err).ToNot(HaveOccurred(), "Expected LoadConfig(path) to not return an error")

	Expect(cfg.UserDb).To(Equal("root"), "Expected UserDb field to be root")
	Expect(cfg.PassDb).To(Equal("mySecretPass"), "Expected PassDb field to be mySecretPass")
	Expect(cfg.Database).To(Equal("mydatabase"), "Expected Database field to be mydatabase")
	Expect(cfg.FilePath).To(Equal("file.csv"), "Expected FilePath field to be file.csv")
}

func Test_WhenConfigFileIsNotPresent_FailsToLoad(t *testing.T) {
	RegisterTestingT(t)

	path := "test_config_invalid.yml"
	cfg := Conf{}
	err := cfg.LoadConfig(path)

	Expect(err).To(HaveOccurred(), "Expected LoadConfig(path) to return an error")
}
