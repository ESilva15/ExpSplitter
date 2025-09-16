// Package config defines the configuration of our app
package config

import (
	"log"
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

// PostgresConfig defines the data required to connect to Postgres.
type PostgresConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	DB   string `yaml:"db"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
}

// Configuration defines the base configuration of the expenses app
type Configuration struct {
	AssetsDir       string          `yaml:"assets"`
	Port            string          `yaml:"port"`
	PgCfg           *PostgresConfig `yaml:"postgres"`
	MigrationsPath  string          `yaml:"migrations_path"`
	MigCustomScript string          `yaml:"mig_custom_scripts"`
}

var (
	instance *Configuration
	once     sync.Once
	confPath string
)

// GetInstance returns the instance of the configuration of the app.
func GetInstance() *Configuration {
	once.Do(func() {
		instance = &Configuration{}
		instance.loadConfiguration()
	})

	return instance
}

// SetConfig sets the path for the configuration file.
func SetConfig(path string) {
	confPath = path
}

func (c *Configuration) loadConfiguration() {
	file, err := os.ReadFile(confPath)
	if err != nil {
		log.Fatalf("Unable to open configuration file [%s]: %s", confPath,
			err.Error())
	}

	err = yaml.Unmarshal(file, &instance)
	if err != nil {
		log.Fatalf("Error parsing JSON: %s", err.Error())
	}
}
