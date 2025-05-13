package utils

import (
	"bufio"
	"log"
	"os"
	"reflect"
	"strings"
)

type Config struct {
	Port  string `env:"APP_PORT" default:":8080"`
	DBUrl string `env:"DATABASE_URL" default:""`
}

func loadEnvFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error opening .env file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || line[0] == '#' {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			log.Fatalf("Invalid line in .env file: %s", line)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		os.Setenv(key, value)
	}
}

func LoadConfig() Config {
	loadEnvFile(".env")
	cfg := Config{}
	v := reflect.ValueOf(&cfg).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		envKey := fieldType.Tag.Get("env")
		defaultValue := fieldType.Tag.Get("default")

		if envValue, exists := os.LookupEnv(envKey); exists {
			field.SetString(envValue)
		} else {
			field.SetString(defaultValue)
		}
	}
	return cfg
}
