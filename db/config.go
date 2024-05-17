package db

import "time"

type Config struct {
	URL                string        `koanf:"url"`
	Name               string        `koanf:"name"`
	ConnecttionTimeout time.Duration `koanf:"connecttion_timeout"`
}
