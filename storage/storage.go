package storage

import (
	"fmt"
	"strings"

	config "github.com/arthurjames/kit/config/storage"
	"github.com/davecgh/go-spew/spew"
	"github.com/fatih/structs"
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

func NewStorage(cfg config.StorageConfig) (*Storage, error) {

	storage := Storage{cfg: cfg}

	db, err := gorm.Open(cfg.Driver, storage.connectString())

	if err != nil {
		err = errors.Wrapf(err,
			"Couldn't open connection to database (%s)",
			spew.Sdump(storage.cfg))
		return nil, err
	}
	defer db.Close()

	storage.Db = db
	return &storage, nil
}

// Get underlying `*sql.DB` from current connection and try to ping it
func (st *Storage) IsOpen() bool {
	db := st.Db.DB()
	if err := db.Ping(); err != nil {
		return false
	}
	return true
}

func (st *Storage) Close() {
	st.Db.Close()
}

func (st *Storage) connectString() string {
	s := structs.New(st.cfg)
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
//
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
