package storage

type StorageOption func(*StorageConfig)

func WithHost(s string) StorageOption {
	return func(sc *StorageConfig) {
		sc.Host = s
	}
}

func WithUser(s string) StorageOption {
	return func(sc *StorageConfig) {
		sc.User = s
	}
}

func WithPassword(s string) StorageOption {
	return func(sc *StorageConfig) {
		sc.Password = s
	}
}

func WithDatabase(s string) StorageOption {
	return func(sc *StorageConfig) {
		sc.Database = s
	}
}

func WithDriver(s string) StorageOption {
	return func(sc *StorageConfig) {
		sc.Driver = s
	}
}
