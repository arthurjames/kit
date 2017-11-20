package storage

import (
	"github.com/arthurjames/kit/config"
)

type StorageConfig struct {
	Host     config.StringOption
	User     config.StringOption
	Password config.StringOption
	Dbname   config.StringOption
	Driver   config.StringOption
	SSLMode  config.EnabledOption
}
