package storage

import (
	"fmt"
	"testing"

	config "github.com/arthurjames/kit/config/storage"
	"github.com/kelseyhightower/envconfig"
)

var store *Storage
var err error
var storageCfg config.StorageConfig

func setup(t *testing.T) func(t *testing.T) {
	t.Log("setup test case")

	err := envconfig.Process("storage", &storageCfg)

	if err != nil {
		t.Fatal("reading config failed: " + err.Error())
	}

	store, err = NewStorage(storageCfg)

	defer store.Close()

	if err != nil {
		t.Fatal(err)
	} else if store == nil {
		t.Fatal("expected db")
	}

	return func(t *testing.T) {
		store.Close()
		t.Log("teardown test case")
	}
}

func TestNewStorage(t *testing.T) {
	f := setup(t)
	defer f(t)
}

func TestIsOpen(t *testing.T) {

	f := setup(t)
	defer f(t)

	if store.IsOpen() {
		t.Fatal("created db is not accessible")
	} else {
		t.Log("created db is accessible")
	}

}

func TestClose(t *testing.T) {
	f := setup(t)
	defer f(t)

	store.Close()
	db := store.Db.DB()
	if err := db.Ping(); err == nil {
		t.Fatal("after closing db should not be accessible")
	} else {
		t.Log("after closing db is not accessible")
	}
}

func TestConnectString(t *testing.T) {
	f := setup(t)
	defer f(t)

	if err != nil {
		t.Fatal("fail")
	}

	if store.connectString() == fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s",
		storageCfg.Host,
		storageCfg.User,
		storageCfg.Password,
		storageCfg.Dbname,
		storageCfg.SSLMode,
	) {
		t.Log("connectstring is correct")
	} else {
		t.Fatal("connectstring is incorrect")
	}
}
