package config

import "main.go/db"

type Config struct {
	Debug    bool      `koanf:"debug"`
	Database db.Config `koanf:"database"`
}
