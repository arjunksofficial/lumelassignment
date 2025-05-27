package database

import (
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PG *gorm.DB

func InitPostgres() error {
	dsn, err := ReadPostgresConfigFromEnv()
	if err != nil {
		return errors.Wrap(err, "failed to read Postgres config from environment variables")
	}
	log.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return errors.Wrap(err, "failed to connect to Postgres database")
	}
	PG = db
	return nil
}

func GetPostgresDB() *gorm.DB {
	if PG == nil {
		err := InitPostgres()
		if err != nil {
			panic("failed to initialize Postgres database: " + err.Error())
		}
	}
	return PG
}
func ClosePostgres() error {
	if PG == nil {
		return nil // No connection to close
	}
	sqlDB, err := PG.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func ReadPostgresConfigFromEnv() (string, error) {
	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDB := os.Getenv("POSTGRES_DB")
	postgresPort := os.Getenv("POSTGRES_PORT")
	if postgresHost == "" || postgresUser == "" || postgresPassword == "" || postgresDB == "" || postgresPort == "" {
		return "", os.ErrInvalid
	}
	dsn := "host=" + postgresHost + " user=" + postgresUser + " password=" + postgresPassword + " dbname=" + postgresDB + " port=" + postgresPort + " sslmode=disable"
	return dsn, nil
}
