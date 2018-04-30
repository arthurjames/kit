package storage

import (
	"fmt"
	"strings"

	"github.com/fatih/structs"
	"github.com/kelseyhightower/envconfig"
)

type StorageConfig struct {
	Host     string `default:"localhost"`
	User     string `default:"postgres"`
	Password string `default:"changeme"`
	Dbname   string `default:"postgres"`
	Driver   string `default:"postgres"`
	SSLMode  string `default:"disable"`
}

// Create new storage config.
func NewConfig() (*StorageConfig, error) {
	cfg := new(StorageConfig)
	if err := envconfig.Process("storage", cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

// Generate PostgreSQL connectstring.
func (cfg *StorageConfig) ConnectString() string {
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
