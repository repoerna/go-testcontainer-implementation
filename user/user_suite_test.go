package user_test

import (
	"context"
	"github.com/testcontainers/testcontainers-go"
	"go-testcontainer-implementation/infra/postgresdb"
	"gorm.io/gorm"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestUser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "User Suite")
}

var ctx = context.Background()
var container testcontainers.Container
var err error
var conn *gorm.DB

var _ = BeforeSuite(func() {
	container, err = postgresdb.SpinUpPostgres(ctx)
	Expect(err).NotTo(HaveOccurred())

	//conn, err := postgresdb.Conn(ctx)
	//Expect(err).NotTo(HaveOccurred())

	err = postgresdb.RunMigration(ctx, "../db/migration")
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	err = container.Terminate(ctx)
})
