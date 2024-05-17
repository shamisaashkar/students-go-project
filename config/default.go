package config

import (
	"time"

	"main.go/db"
)

func Default() Config {
	return Config{
		Debug: true,
		Database: db.Config{
			URL:                "mongodb://127.0.0.1:27017",
			Name:               "students",
			ConnecttionTimeout: time.Second,
		},
	}
}
