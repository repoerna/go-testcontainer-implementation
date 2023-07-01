package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type postgresConfig struct {
	Host     string `yaml:"host"`
	Port     uint   `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbName"`
	SSLMode  string `yaml:"sslMode"`
}

func (p *postgresConfig) GetDatabaseURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		p.Username,
		p.Password,
		p.Host,
		p.Port,
		p.DBName,
		p.SSLMode,
	)
}

type config struct {
	DB postgresConfig `yaml:"database"`
}

func loadConfiguration(file string) config {
	// Read the YAML file
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Failed to read YAML file %s: %v", file, err)
	}

	// Parse the YAML data into AppConfig struct
	var cfg config
	err = yaml.Unmarshal(data, &cfg)
	log.Println(cfg)
	if err != nil {
		log.Fatalf("Failed to parse YAML %s: %v", file, err)
	}

	return cfg
}
