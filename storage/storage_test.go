package storage

import (
	"testing"

	config "github.com/arthurjames/kit/config/storage"
)

var store *Storage

func setup(t *testing.T) func(t *testing.T) {
	t.Log("setup test case")

	cfg, err := config.NewConfig()
	if err != nil {
		t.Fatal("Creating config failed: " + err.Error())
	}

	store, err = NewStorage(cfg)
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

	if _, err := store.IsOpen(); err != nil {
		t.Fatalf("created db is not accessible: %s", err)
	} else {
		t.Log("created db is accessible")
	}

}

func TestClose(t *testing.T) {
	f := setup(t)
	defer f(t)

	store.Close()

	if err := store.DB.Ping(); err != nil {
		t.Log(err.Error())
	} else {
		t.Fatal("after closing db should not be accessible")
	}
}
