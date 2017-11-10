package storage

import (
	"testing"

	config "github.com/arthurjames/kit/config/storage"
)

func TestNew(t *testing.T) {

	db, err := NewStorage(
		config.WithHost("localhost"),
		config.WithUser("postgres"),
		config.WithPassword("changeme"),
		config.WithDatabase("postgres"),
		config.WithDriver("postgres"),
	)
	if err != nil {
		t.Fatal(err)
	} else if db == nil {
		t.Fatal("expected db")
	}
}
