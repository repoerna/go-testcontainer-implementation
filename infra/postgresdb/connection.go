package postgresdb

import (
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Conn(dsn string) (*gorm.DB, error) {
	if dsn == "" {
		return nil, errors.New("database source name is empty")
	}

	conn, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)

	if err != nil {
		return nil, err
	}

	return conn, nil
}
