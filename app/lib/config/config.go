package config

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
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

func init() {
	viper.SetConfigType("toml")
	viper.SetConfigName("swerver.config")
	viper.AddConfigPath("$HOME/.config/swerver")
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error cannot find config file: %s", err))
	}
}

// LoadConfig loads the Config struct via configPath
func LoadConfig() Configuration {

	loadDefaults()

	port := viper.Get("port").(string)
	scriptsPath := viper.Get("scripts_path").(string)
	servicesToMonitor := viper.Get("services_to_monitor").(string)
	iplookupURL := viper.Get("ip_lookup_url").(string)
	webPath := viper.Get("web_path").(string)

	services := strings.Split(servicesToMonitor, ",")

	return Configuration{
		Port:          port,
		StaticPath:    path.Join(webPath, "static"),
		ScriptsPath:   scriptsPath,
		Services:      services,
		TemplatesPath: path.Join(webPath, "templates"),
		IPLookupURL:   iplookupURL,
	}
}

func loadDefaults() {

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	viper.SetDefault("port", "8080")
	viper.SetDefault("path", path)
	viper.SetDefault("scripts_path", filepath.Join(path, "scripts"))
	viper.SetDefault("web_path", filepath.Join(path, "web"))
}
