package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func FindUserByID(db *gorm.DB, id uuid.UUID) (*User, error) {
	var user User

	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func SaveUser(db *gorm.DB, user User) error {
	return db.Create(user).Error
}
