package postgresdb

import (
	"context"
	"fmt"
	"github.com/docker/go-connections/nat"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"os"
	"time"
)

const (
	TestDbUsername = "postgres"
	TestDbPassword = "postgres"
	TestDbHost     = "127.0.0.1"
	TestDbName     = "test_db"
)

func setMappedDBURL(mappedPort nat.Port) {
	err := os.Setenv("DATABASE_URL", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		TestDbUsername,
		TestDbPassword,
		TestDbHost,
		mappedPort.Port(),
		TestDbName,
	))
	if err != nil {
		log.Fatal(err)
	}
}

func GetMappedDBURL() string {
	return os.Getenv("DATABASE_URL")
}

func SpinUpPostgres(ctx context.Context) (testcontainers.Container, error) {

	// prepare parameter
	waitTimeoutDuration := 60 * time.Second

	env := map[string]string{
		"POSTGRES_USER":     TestDbUsername,
		"POSTGRES_PASSWORD": TestDbPassword,
		"POSTGRES_DB":       TestDbName,
	}

	containerPort := "5432/tcp"
	natPort := nat.Port(containerPort)

	dbURL := func(host string, port nat.Port) string {
		return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			TestDbUsername,
			TestDbPassword,
			TestDbHost,
			port.Port(),
			TestDbName,
		)
	}

	waitFunc := wait.ForSQL(natPort, "postgres", dbURL)
	waitFunc.WithStartupTimeout(waitTimeoutDuration)

	// setup container parameter
	containerRequest := testcontainers.ContainerRequest{
		Image:        "postgres:15.3",
		ExposedPorts: []string{containerPort},
		WaitingFor:   waitFunc,
		Env:          env,
	}

	// create container and start container
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: containerRequest,
		Started:          true,
	})

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer func() {
		if r := recover(); r != nil {
			err := container.Terminate(ctx)
			if err != nil {
				log.Fatal(err)
			}
		}
	}()

	// get mapped port
	mappedExposedPort, err := container.MappedPort(ctx, natPort)
	if err != nil {
		log.Fatal(err)
	}
	setMappedDBURL(mappedExposedPort)

	return container, nil
}

func RunMigration(filepath string) error {
	dsn := GetMappedDBURL()
	conn, err := Conn(dsn)
	if err != nil {
		return err
	}

	sqlDB, err := conn.DB()
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})

	migrationFIlePath := fmt.Sprintf("file://%s", filepath)

	migrator, err := migrate.NewWithDatabaseInstance(
		migrationFIlePath,
		"postgres",
		driver,
	)
	if err != nil {
		return err
	}

	err = migrator.Up()
	if err != nil {
		return err
	}

	return nil

}

func TruncateTable(tableName string) error {
	dsn := GetMappedDBURL()
	db, err := Conn(dsn)
	if err != nil {
		return err
	}

	err = db.Raw(
		"TRUNCATE TABLE RETURN ? IDENTITY CASCADE",
		tableName,
	).Error
	if err != nil {
		return err
	}

	return nil
}
