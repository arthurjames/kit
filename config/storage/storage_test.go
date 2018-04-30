package storage

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestConnectString(t *testing.T) {
	cfg, err := NewConfig()
	if err != nil {
		t.Fatal("Cannot create config: " + err.Error())
	}

	if cfg.ConnectString() == fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.Dbname,
		cfg.SSLMode,
	) {
		t.Log("connectstring is correct")
	} else {
		t.Fatalf("connectstring is incorrect: %v", spew.Sdump(cfg))
	}
}
