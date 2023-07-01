package user

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"primaryKey"`
	Email    string    `gorm:"unique,not null"`
	Password string    `gorm:"not null"`
}
