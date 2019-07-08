package config

import (
	"log"
	"os"
	"path"
	"strings"

	"github.com/joho/godotenv"
)

// Configuration struct contains the json decoded options
type Configuration struct {
	Port          string
	StaticPath    string
	Services      []string
	ScriptsPath   string
	TemplatesPath string
	IPLookupURL   string
}

// LoadConfig loads the Config struct via configPath
func LoadConfig() Configuration {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
	portEnv := getEnv("SWERVER_PORT")
	pathEnv := getEnv("SWERVER_PATH")
	servicesEnv := getEnv("SWERVER_SERVICES")
	iplookupEnv := getEnv("SWERVER_IPLOOKUP_URL")

	services := strings.Split(servicesEnv, ",")

	return Configuration{
		Port:          portEnv,
		StaticPath:    path.Join(pathEnv, "web", "static"),
		ScriptsPath:   path.Join(pathEnv, "scripts"),
		Services:      services,
		TemplatesPath: path.Join(pathEnv, "web", "templates"),
		IPLookupURL:   iplookupEnv,
	}
}

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Fatal("No .env value found for " + key)
	return ""
}
