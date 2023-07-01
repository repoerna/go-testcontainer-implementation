package main

import (
	"context"
	"go-testcontainer-implementation/infra/postgresdb"
	"go-testcontainer-implementation/user"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	cfg := loadConfiguration("config.yml")

	dburl := cfg.DB.getDatabaseURL()
	os.Setenv("DATABASE_URL", dburl)

	conn, err := postgresdb.Conn(ctx)
	if err != nil {
		log.Fatal("can't establish connection to database")
	}

	conn.AutoMigrate(user.User{})

}
