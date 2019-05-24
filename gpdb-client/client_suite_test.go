package gpdb_client

import (
. "github.com/onsi/ginkgo"
. "github.com/onsi/gomega"

"testing"
)

func TestGPDBClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GBDB Client Suite")
}
