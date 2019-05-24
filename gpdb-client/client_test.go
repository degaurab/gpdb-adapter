package gpdb_client

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"io"
	"log"
)

var _ = Describe("GPDBClient", func() {
	var (
		stderr   = gbytes.NewBuffer()
		logger   *log.Logger
		dbDriver DBDriver
	)

	BeforeEach(func() {
		logger = log.New(io.MultiWriter(stderr, GinkgoWriter), "[gbpd-service-adapter]", log.LstdFlags)
		dbDriver = DBDriver{
			User: "postgres",
			Password: "password",
			Hostname: "localhost",
			Port: 5432,
		}
	})

	Describe("TestConnection", func() {
		Context("when DB connection is working", func() {
			It("", func() {
				connResponse := dbDriver.TestConnection(logger)
				Expect(connResponse).To(BeNil())
			})
		})

		Context("when config is wrond", func() {
			BeforeEach(func() {
				dbDriver = DBDriver{
					User: "postgres",
					Password: "password",
					Hostname: "wrong-host",
					Port: 5432,
				}
			})

			It("", func() {
				connResponse := dbDriver.TestConnection(logger)
				Expect(connResponse).To(MatchError("test connection failed"))
			})

		})
	})

	//TODO: we are going to mostly test pq lib at this point by injecting SQL queries.
	XDescribe("InitializeDBForUser", func() {
		Context("when db and username provided", func() {
			BeforeEach(func() {
				//clean-up
			})
			It("should create user", func() {

			})
		})
	})

	//TODO: we are going to mostly test pq lib at this point by injecting SQL queries.
	//      testing templates might make more sense
	XDescribe("DeleteDatabase", func() {
		Context("when db and username provided", func() {
			BeforeEach(func() {
				//clean-up
			})
			It("should create user", func() {

			})
		})
	})
})
