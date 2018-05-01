package storage

import (
	"database/sql"

	config "github.com/arthurjames/kit/config/storage"
	"github.com/davecgh/go-spew/spew"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

// Storage holds a pointer that represents a pool of 0 or more
// connections.
type Storage struct {
	*sql.DB
}

// Transaction models the standard transaction in database/sql.
type Transaction interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

// A Txfn is a function that will be called with an initialized `Transaction` object
// that can be used for executing statements and queries against a database.
type TxFn func(Transaction) error

// Open new storage.
func NewStorage(cfg *config.StorageConfig) (*Storage, error) {
	db, err := sql.Open(cfg.Driver, cfg.ConnectString())
	if err != nil {
		err = errors.Wrapf(err,
			"Couldn't open connection to database (%s)",
			spew.Sdump(cfg))
		return nil, err
	}

	return &Storage{db}, nil
}

// Get underlying `*sql.DB` from current connection and try to ping it.
func (st *Storage) IsOpen() (bool, error) {
	if err := st.DB.Ping(); err != nil {
		return false, err
	}
	return true, nil
}

// Close open database connection.
func (st *Storage) Close() {
	st.DB.Close()
}

// Begin starts and returns a new transaction.
//func (st *Storage) Begin() (*sql.Tx, error) {
//	tx, err := st.DB.Begin()
//	if err != nil {
//		return nil, err
//	}
//	return tx, nil
//}
//

// WithTransaction creates a new transaction and handles rollback/commit based on the
// error object returned by the `TxFn`
func (st *Storage) WithTransaction(fn TxFn) (err error) {
	tx, err := st.DB.Begin()
	if err != nil {
		return
	}

	defer func() {
		if p := recover(); p != nil {
			// a panic occurred, rollback
			tx.Rollback()
			// rethrow panic
			panic(p)
		} else if err != nil {
			// something went wrong, rollback
			tx.Rollback()
		} else {
			// all good, commit
			err = tx.Commit()
		}
	}()

	err = fn(tx)
	return err
}

// SetMaxOpenConns sets the maximum number of open connections to the db.
//func MaxOpenConns(v int) Option {
//	return func(p *Storage) {
//		p.Db.SetMaxOpenConns(v)
//	}
//}
//
//func (ds *Datastore) Migrate(values ...interface{}) {
//	ds.Db.AutoMigrate(values)
//}
//func MaxIdleConns(v int) Option {
//	return func(p *Storage) {
//		p.Db.SetMaxIdleConns(v)
//	}
//}
//

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
