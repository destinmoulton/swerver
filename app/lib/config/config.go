package config

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

var configPath = ""
var configFileName = "swerver.config"
var configFileType = "toml"

// Configuration struct contains the json decoded options
type Configuration struct {
	Port          string
	StaticPath    string
	Services      []string
	ScriptsPath   string
	TemplatesPath string
	IPLookupURL   string
	Username      string
	PasswordHash  string
}

func init() {

	usr, uerr := user.Current()
	if uerr != nil {
		panic(uerr)
	}

	configPath = filepath.Join(usr.HomeDir, ".config/swerver")

	if !doesConfigDirExist() {
		fmt.Println("config dir doesn't exist")
		createConfigDirAndFile()
	}

	viper.SetConfigType(configFileType)
	viper.SetConfigName(configFileName)
	viper.AddConfigPath(configPath)
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error cannot find config file: %s", err))
	}

	loadDefaults()

}

// LoadConfig loads the Config struct via configPath
func LoadConfig() Configuration {

	port := viper.GetString("port")
	scriptsPath := viper.GetString("scripts_path")
	servicesToMonitor := viper.GetString("services_to_monitor")
	iplookupURL := viper.GetString("ip_lookup_url")
	webPath := viper.GetString("web_path")
	username := viper.GetString("username")
	passwordHash := viper.GetString("password")

	services := strings.Split(servicesToMonitor, ",")

	return Configuration{
		Port:          port,
		StaticPath:    path.Join(webPath, "static"),
		ScriptsPath:   scriptsPath,
		Services:      services,
		TemplatesPath: path.Join(webPath, "templates"),
		IPLookupURL:   iplookupURL,
		Username:      username,
		PasswordHash:  passwordHash,
	}
}

// Save the config
func Save(options map[string]string) {

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
	viper.SetDefault("services_to_monitor", "")
	viper.SetDefault("ip_lookup_url", "")
	viper.SetDefault("username", "")
	viper.SetDefault("password", "")
}

func doesConfigDirExist() bool {
	info, err := os.Stat(configPath)
	if os.IsNotExist(err) {

		fmt.Println("directory doesn't exist")
		return false
	}
	return info.IsDir()
}

func createConfigDirAndFile() error {
	err := os.MkdirAll(configPath, os.ModePerm)

	if err != nil {

		return err
	}

	f, ferr := os.Create(filepath.Join(configPath, configFileName+"."+configFileType))

	if ferr != nil {
		return ferr
	}
	f.Close()

	return nil
}
