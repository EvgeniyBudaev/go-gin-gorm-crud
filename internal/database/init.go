// database - module initialization
package database

import (
	"github.com/EvgeniyBudaev/go-gin-gorm-crud/internal/storage/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func NewConnectionToDB() (*gorm.DB, error) {
	dbcon := postgres.ConfigDB{
		Host:       os.Getenv(postgres.HOST),
		Port:       os.Getenv(postgres.PORT),
		User:       os.Getenv(postgres.USER),
		Password:   os.Getenv(postgres.PASSWORD),
		DBName:     os.Getenv(postgres.NAME),
		SchemeName: os.Getenv(postgres.SCHEME_NAME),
		SSLMode:    os.Getenv(postgres.SSL_MODE),
		LogLevel:   os.Getenv(postgres.LOG_LEVEL),
	}
	dbClient, err := postgres.New(dbcon)
	if err != nil {
		log.Fatalf("error running postresql: %s", err.Error())
		return nil, err
	}
	return dbClient, nil
}
