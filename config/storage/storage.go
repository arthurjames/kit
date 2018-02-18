package storage

type StorageConfig struct {
	Host     string `default:"localhost"`
	User     string `default:"postgres"`
	Password string `default:"changeme"`
	Dbname   string `default:"postgres"`
	Driver   string `default:"postgres"`
	SSLMode  string `default:"disable"`
}
