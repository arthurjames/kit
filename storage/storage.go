package storage

import (
	"fmt"

	config "github.com/arthurjames/kit/config/storage"
	"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
)

type Storage struct {
	// Db holds a pointer that represents a pool of 0 or more
	// connections.
	Db  *gorm.DB
	cfg config.StorageConfig
}

// Create new storage
func NewStorage(options ...func(*config.StorageConfig)) (*Storage, error) {

	var mgr Storage
	for _, option := range options {
		option(&mgr.cfg)
	}

	url := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		mgr.cfg.Host, mgr.cfg.User, mgr.cfg.Password, mgr.cfg.Database)
	db, err := gorm.Open(mgr.cfg.Driver, url)

	if err != nil {
		err = errors.Wrapf(err,
			"Couldn't open connection to database (%s)",
			spew.Sdump(mgr.cfg))
		return nil, err
	}

	mgr.Db = db
	return &mgr, nil
}

// Set max idle connections of connectionpool
func (m *Storage) SetMaxIdleConns(v int) {
	m.Db.DB().SetMaxIdleConns(v)
}

// Set max open connections of connectionpool
func (m *Storage) SetMaxOpenConns(v int) {
	m.Db.DB().SetMaxOpenConns(v)
}

// Use gorm automigrate to create new tables and columns
func (m *Storage) AutoMigrate(interfaces ...interface{}) {
	m.Db.AutoMigrate(interfaces)
}
