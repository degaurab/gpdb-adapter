package helper_test

import (
	"github.com/degaurab/gpdb-adapter/helper"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func init() {
	Describe("RandStringBytes", func() {
		const randoStringSize = 10
		randStringList := make(map[string]int)

		Context("when random string requested", func() {
			It("generated random string between requests", func() {
				for i := 0; i < randoStringSize; i++ {
					randString := helper.RandStringBytes(randoStringSize)
					if _, ok := randStringList[randString]; ok {
						Expect(ok).ToNot(Equal(true))
					} else {
						randStringList[randString] = 1
					}
				}
			})
		})
	})
}
