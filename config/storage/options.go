package storage

import (
	"github.com/arthurjames/kit/config"
)

type StorageOption func(*StorageConfig)

func WithDatabase(s string) StorageOption {
	return func(sc *StorageConfig) {
		sc.Dbname = config.StringOption{"dbname", s}
	}
}

func WithDriver(s string) StorageOption {
	return func(sc *StorageConfig) {
		sc.Driver = config.StringOption{"driver", s}
	}
}

func WithHost(s string) StorageOption {
	return func(sc *StorageConfig) {
		sc.Host = config.StringOption{"host", s}
	}
}

func WithPassword(s string) StorageOption {
	return func(sc *StorageConfig) {
		sc.Password = config.StringOption{"password", s}
	}
}

func WithSSLMode(b bool) StorageOption {
	return func(sc *StorageConfig) {
		sc.SSLMode = config.EnabledOption{"sslmode", b}
	}
}

func WithUser(s string) StorageOption {
	return func(sc *StorageConfig) {
		sc.User = config.StringOption{"user", s}
	}
}
