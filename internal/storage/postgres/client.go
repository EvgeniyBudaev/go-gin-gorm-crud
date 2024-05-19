package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"strconv"
	"time"
)

// Client for PostgreSQL database
type PGStorage struct {
	//Configure connection
	cfg ConfigDB
	//Gorm databese
	db *gorm.DB
}

// Create new PGStorage client
func New(cfg ConfigDB) (dbs *gorm.DB, err error) {
	cfgStr := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=%s`,
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.SSLMode)
	logLevel, err := strconv.Atoi(cfg.LogLevel)
	if err != nil {
		return nil, err
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.LogLevel(logLevel),
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
	gc := gorm.Config{
		Logger: newLogger,
	}
	if len(cfg.SchemeName) != 0 {
		gc.NamingStrategy = schema.NamingStrategy{
			TablePrefix: cfg.SchemeName,
		}
	}
	db, err := gorm.Open(postgres.Open(cfgStr), &gc)
	return db, err
}
