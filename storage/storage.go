package storage

import (
	"database/sql"
	"fmt"
	"strings"

	config "github.com/arthurjames/kit/config/storage"
	"github.com/davecgh/go-spew/spew"
	"github.com/fatih/structs"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

// Storage holds a pointer that represents a pool of 0 or more
// connections.
type Storage struct {
	*sql.DB
}

// Open storage
func NewStorage(cfg config.StorageConfig) (*Storage, error) {

	db, err := sql.Open(cfg.Driver, connectString(cfg))

	if err != nil {
		err = errors.Wrapf(err,
			"Couldn't open connection to database (%s)",
			spew.Sdump(cfg))
		return nil, err
	}

	return &Storage{db}, nil
}

// Get underlying `*sql.DB` from current connection and try to ping it
func (st *Storage) IsOpen() (bool, error) {
	if err := st.Ping(); err != nil {
		return false, err
	}
	return true, nil
}

// Close open database connection
func (st *Storage) Close() {
	st.DB.Close()
}

// Generate PostgreSQL connectstring
func connectString(cfg config.StorageConfig) string {
	s := structs.New(cfg)
	str := []string{}
	for _, name := range s.Names() {
		field := s.Field(name)

		if !field.IsZero() && name != "Driver" {
			str = append(str, fmt.Sprintf("%s=%s", strings.ToLower(name), field.Value()))
		}
	}
	return strings.Join(str, " ")
}

//func (ds *Datastore) Migrate(values ...interface{}) {
//	ds.Db.AutoMigrate(values)
//}
//func MaxIdleConns(v int) Option {
//	return func(p *Storage) {
//		p.Db.SetMaxIdleConns(v)
//	}
//}
//
//func MaxOpenConns(v int) Option {
//	return func(p *Storage) {
//		p.Db.SetMaxOpenConns(v)
//	}
//}
//
//func LogMode(b bool) Option {
//	return func(p *Storage) {
//		p.Db.LogMode(b)
//	}
//}
//
//type StorageDriver interface {
//	New(string) StorageManager
//	RegisterDriver(string)
//}
//
//var drivers = []Storage{{Name: "postgresql",
//	Path: "github.com/jinzhu/gorm/dialects/postgres"}}
//
