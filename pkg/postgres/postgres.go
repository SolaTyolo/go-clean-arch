// Package postgres implements postgres connection.
package postgres

import (
	"log"
	"os"
	"time"

	postgres "go.elastic.co/apm/module/apmgormv2/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Postgres -.
type Postgres struct {
	DB *gorm.DB
}

// New -.
func New(url string) (*Postgres, error) {
	pg := &Postgres{}

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		PrepareStmt: true,
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             5 * time.Second,
				LogLevel:                  logger.Error,
				IgnoreRecordNotFoundError: true,
			},
		),
		CreateBatchSize:          88,
		DisableNestedTransaction: true,
	})

	if err != nil {
		return nil, err
	}
	pg.DB = db
	return pg, nil
}
