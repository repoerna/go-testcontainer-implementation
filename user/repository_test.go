package user_test

import (
	"context"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go-testcontainer-implementation/infra/postgresdb"
	. "go-testcontainer-implementation/user"
	"gorm.io/gorm"
)

var _ = Describe("Repository", func() {
	var err error
	var conn *gorm.DB
	ctx := context.Background()

	id := uuid.New()
	newUser := User{
		ID:       id,
		Email:    "new_user@exmple.com",
		Password: "password",
	}

	It("prepare connection", func() {
		conn, err = postgresdb.Conn(ctx)
		Expect(err).NotTo(HaveOccurred())
	})

	Context("SaveUser", func() {
		It("successfully insert user data db", func() {
			err := SaveUser(conn, newUser)
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("FindUserByID", func() {
		It("successfully get user data", func() {
			user, err := FindUserByID(conn, id)
			Expect(err).NotTo(HaveOccurred())
			Expect(user).To(Equal(&newUser))
		})
	})
})
