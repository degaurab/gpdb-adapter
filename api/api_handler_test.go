package api


import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"io"
	"log"
)


var _ = Describe("API Handler", func() {
	var (
		stderr = gbytes.NewBuffer()
		logger *log.Logger

	)

	BeforeEach(func() {
		logger = log.New(io.MultiWriter(stderr, GinkgoWriter), "[gpdb-service-adapter]", log.LstdFlags)
	})

	Describe("serviceCatalog", func() {

	})
})
