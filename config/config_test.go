package config_test

import (
	"io"
	"log"

	"github.com/degaurab/gbdb-adapter/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("LoadConfig", func() {
	var (
		stderr           = gbytes.NewBuffer()
		sampleConfigPath string
		logger           *log.Logger
	)

	BeforeEach(func() {
		logger = log.New(io.MultiWriter(stderr, GinkgoWriter), "[gbpd-service-adapter]", log.LstdFlags)
	})

	Context("when config path is correct", func() {
		BeforeEach(func() {
			sampleConfigPath = "samples/service-config.yml"
		})

		Context("whn config is correct", func() {
			It("should return config template", func() {
				config, err := config.LoadConfig(sampleConfigPath, logger)

				Expect(err).NotTo(HaveOccurred())
				Expect(config.InstanceIP).To(Equal("localhost"))
			})
		})

	})

	Context("when config path is incorrect", func() {
		BeforeEach(func() {
			sampleConfigPath = "wrong-path-to/service-config.yml"
		})

		It("should return error", func() {
			_, err := config.LoadConfig(sampleConfigPath, logger)

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("no such file or director"))
		})
	})

})
