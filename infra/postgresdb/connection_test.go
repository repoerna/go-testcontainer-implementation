package postgresdb_test

import (
	"context"
	"database/sql"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/testcontainers/testcontainers-go"
	. "go-testcontainer-implementation/infra/postgresdb"
	"log"
	"os"
)

var _ = Describe("Conn", func() {
	var err error
	var container testcontainers.Container
	var ctx = context.Background()

	BeforeSuite(func() {
		container, err = SpinUpPostgres(ctx)
		if err != nil {
			log.Fatal(err)
		}
	})

	AfterSuite(func() {
		container.Terminate(ctx)
	})

	It("Can establish connection", func() {
		dsn := os.Getenv("DATABASE_URL")

		// Open a connection to the postgresdb
		db, err := sql.Open("postgres", dsn)
		defer db.Close()
		Expect(err).NotTo(HaveOccurred())

		// Ping the postgresdb to test the connection
		err = db.Ping()
		Expect(err).NotTo(HaveOccurred())
	})
})
