package main_test

import "github.com/onsi/gomega/gexec"
import . "github.com/onsi/ginkgo"
import . "github.com/onsi/gomega"

var _ = Describe("adapter server executable", func() {
	It("conforms to the server interface, and is buildable", func() {
		_, err := gexec.Build("github.com/degaurab/gbdb-adapter")
		Expect(err).ToNot(HaveOccurred())
	})
})
