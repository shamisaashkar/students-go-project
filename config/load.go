package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
)

const (
	prefix    = "STUDENTS_"
	delimeter = "."
	seprator  = "__"
)

// STUDENTS_DEBUG --> DEBUG --> debug
// STUDENTS_DATABASE_HOST --> DATABASE_HOST --> database_host --> database.host --> database.

func callbackEnv(source string) string {
	base := strings.ToLower(strings.TrimPrefix(source, prefix))
	return strings.ReplaceAll(base, seprator, delimeter)
}

func New() Config {
	k := koanf.New(".")
	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		log.Fatalf("error loading default: %s", err)
	}
	if err := k.Load(file.Provider("config.toml"), toml.Parser()); err != nil {
		log.Printf("error loading config: %s", err)
	}
	var instance Config
	if err := k.Unmarshal("", &instance); err != nil {
		log.Fatalf("error unmarshaling config: %s", err)
	}
	fmt.Printf("%+v", instance)
	return instance

}
