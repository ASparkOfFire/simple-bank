package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/ASparkOfFire/simple-bank/vars"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		logrus.Fatal("an argument is required: (\"up\", \"down\", \"reset\")")
		return
	}
	cmd := os.Args[1]

	// Connect to the database
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		vars.DB_USER,
		vars.DB_PASS,
		vars.DB_HOST,
		vars.DB_PORT,
		vars.DB_NAME,
	))
	if err != nil {
		logrus.Fatalf("Error while opening connection to database: %v", err)
		return
	}
	defer db.Close()

	// Create a migration driver instance
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logrus.Fatalf("Error creating migration driver: %v", err)
		return
	}

	// Create a new migrate instance
	m, err := migrate.NewWithDatabaseInstance("file://db/migration", "postgres", driver)
	if err != nil {
		logrus.Fatalf("Error creating migrate instance: %v", err)
		return
	}

	// Perform the migration
	switch strings.ToLower(cmd) {
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	case "reset":
		_, err := db.Exec(`DROP TABLE IF EXISTS schema_migrations;`)
		if err != nil {
			logrus.Fatalf("Error while resetting database: %v", err)
			return
		}
	default:
		logrus.Fatalf("Unsupported migration command: %s", cmd)
	}

	// Handle migration errors
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			logrus.Info("No migration changes detected")
		} else {
			logrus.Fatalf("Migration error: %v", err)
		}
	} else {
		logrus.Info("Migration applied successfully")
	}
}
