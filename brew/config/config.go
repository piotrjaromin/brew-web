package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Temperature struct {
	RefreshIntervalSeconds int `yaml:"refreshIntervalSeconds"`
}

type Heater struct {
	Pin int `yaml:"pin"`
}

type Keg struct {
	Heaters []Heater `yaml:"heaters"`
}

// Config contains configuration data for modules in this project
type Config struct {
	Keg         Keg         `yaml:"keg"`
	Port        int         `yaml:"port"`
	Temperature Temperature `yaml:"temperature"`
}

// GetConfig creates Config struct and fills it fields
// from yaml found under configPath
func GetConfig(configPath string) (*Config, error) {
	file, errFile := os.Open(configPath)
	if errFile != nil {
		return nil, fmt.Errorf("error while reading config file: %s", errFile.Error())
	}

	decoder := yaml.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		return nil, fmt.Errorf("error while un marshaling config file: %s", errFile.Error())
	}

	return &configuration, nil
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
