package postgresdb_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPostgresdb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Postgresdb Suite")
}
