package postgresdb_test

import (
	"context"
	_ "github.com/lib/pq"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/testcontainers/testcontainers-go"
	. "go-testcontainer-implementation/infra/postgresdb"
	"os"
)

var _ = Describe("SpinUpPostgres", func() {
	var err error
	var container testcontainers.Container
	var ctx = context.Background()

	It("No error occurred", func() {

		container, err = SpinUpPostgres(ctx)
		Expect(err).NotTo(HaveOccurred())

		dsn := os.Getenv("DATABASE_URL")
		Expect(dsn).NotTo(BeEmpty())
		Expect(dsn).Should(ContainSubstring("postgres"))

		container.Terminate(ctx)
	})

	//It("successfully run migration", func() {
	//
	//})

})
