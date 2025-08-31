package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"sync"
)

type Configuration struct {
	AssetsDir        string `yaml:"assets"`
	Port             string `yaml:"port"`
	DBSys            string `yaml:"dbsys"`
	DBPath           string `yaml:"dbpath"`
	MigrationsPath   string `yaml:"migrations_path"`
	MigCustomScript  string `yaml:"mig_custom_scripts"`
	CurrencyAffinity int32  `yaml:"currency_affinity"`
}

var (
	instance *Configuration
	once     sync.Once
	confPath string
)

func GetInstance() *Configuration {
	once.Do(func() {
		if instance != nil {
			return
		}
		instance = &Configuration{}
		instance.loadConfiguration()
	})

	return instance
}

func SetConfPath(path string) {
	confPath = path
}

func SetConfiguration(config *Configuration) {
	instance = config
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
