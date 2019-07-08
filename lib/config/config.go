package config

import (
	"encoding/json"
	"log"
	"os"
)

// Configuration struct contains the json decoded options
type Configuration struct {
	Port          string
	AssetsPath    string
	TemplatesGlob string
	Services      []string
	ScriptsPath   string
	Scripts       []string
}

// LoadConfig loads the Config struct via configPath
func LoadConfig(configPath string) Configuration {
	file, err := os.Open(configPath)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(file)
	config := Configuration{}
	derr := decoder.Decode(&config)
	if derr != nil {
		log.Fatal(derr)
	}
	return config
}
