package storage

import (
	"testing"

	config "github.com/arthurjames/kit/config/storage"
)

var store *Storage
var err error

func setup(t *testing.T) func(t *testing.T) {
	t.Log("setup test case")

	store, err = NewStorage(
		config.WithHost("localhost"),
		config.WithUser("postgres"),
		config.WithDriver("postgres"),
		config.WithSSLMode(false),
	)
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

	if store.connectString() == "host=localhost user=postgres sslmode=disable" {
		t.Log("connectstring is correct")
	} else {
		t.Fatal("connectstring is incorrect")
	}
}
