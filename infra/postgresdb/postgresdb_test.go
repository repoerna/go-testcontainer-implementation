package postgresdb_test

import (
	"context"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/testcontainers/testcontainers-go"
	. "go-testcontainer-implementation/infra/postgresdb"
	"os"
)

// postgresdb test case in one file  is to simplified test creation

var _ = Describe("Setup Postgres Container", func() {
	var err error
	var container testcontainers.Container
	var ctx = context.Background()
	var dsn string

	Context("spin up postgres container", Ordered, func() {
		AfterAll(func() {
			DeferCleanup(func() {
				container.Terminate(ctx)
			})
		})

		It("no error occurred", func() {
			container, err = SpinUpPostgres(ctx)
			Expect(err).NotTo(HaveOccurred())
			dsn = GetMappedDBURL()
		})

		It("set the DATABASE_URL env", func() {
			dsn := os.Getenv("DATABASE_URL")
			Expect(dsn).NotTo(BeEmpty())
			Expect(dsn).Should(ContainSubstring("postgres"))
		})

		When("container ready", Ordered, func() {
			It("can establish connection", func() {
				conn, err := Conn(dsn)
				Expect(err).NotTo(HaveOccurred())

				db, err := conn.DB()
				Expect(err).NotTo(HaveOccurred())

				// Ping the postgresdb to test the connection
				err = db.Ping()
				Expect(err).NotTo(HaveOccurred())
			})

			It("can run migration", func() {
				err := RunMigration("testmigration")
				Expect(err).NotTo(HaveOccurred())
			})

			It("can truncate table", func() {
				err := TruncateTable("table_test")
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})
})
