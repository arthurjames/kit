package storage

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Storage struct {
	*gorm.DB
}

func New() (*Storage, error) {
	dbhost := os.Getenv("DBHOST")
	if dbhost != "" {
		dbhost = fmt.Sprintf("host=%v ", dbhost)
	}
	db, err := gorm.Open("postgres",
		fmt.Sprintf("%vuser=admin password=changeme dbname=925 sslmode=disable", dbhost))
	if err != nil {
		log.Fatalf("Error connecting to database: '%v'", err)
		return nil, err
	}

	db.DB().SetMaxIdleConns(0) // See https://github.com/jinzhu/gorm/issues/246
	db.DB().SetMaxOpenConns(100)
	//	db.SetLogger(gorm.Logger{level.Trace})
	db.LogMode(true)

	// Configure any package-level settings
	return &Storage{db}, nil
}

func (db *Storage) Close() {
	db.Close()
}

func (db *Storage) Migrate(values ...interface{}) {
	db.AutoMigrate(values)
}
