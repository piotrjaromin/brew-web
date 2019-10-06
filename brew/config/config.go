package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type Temperature struct {
	RefreshInterval time.Duration `yaml:"refreshIntervalSeconds"`
}

type Heater struct {
	Pin string `yaml:"pin"`
}

type Keg struct {
	Heaters     []Heater    `yaml:"heaters"`
	Temperature Temperature `yaml:"temperature"`
}

// Config contains configuration data for modules in this project
type Config struct {
	Keg Keg `yaml:"keg"`
}

// GetConfig creates Config struct and fills it fields
// from yaml found under configPath
func GetConfig(configPath string) Config {
	file, errFile := os.Open(configPath)
	if errFile != nil {
		fmt.Println("error while reading config file:", errFile)
	}

	decoder := yaml.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error while un marshaling config file:", err)
	}

	return configuration
}

// GetEnvOrDefault reads environment variable and returns value
// if there is no environment variable present then def value is returned
func GetEnvOrDefault(key string, def string) string {
	fromEnv := os.Getenv(key)
	if len(fromEnv) == 0 {
		return def
	}
	return fromEnv
}
