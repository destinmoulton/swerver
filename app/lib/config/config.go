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

	"github.com/destinmoulton/swerver/app/lib/rando"
)

var configPath = ""
var configFileName = "swerver.config"
var configFileType = "toml"

// Configuration struct contains the json decoded options
type Configuration struct {
	Port             string
	Services         []string
	ScriptsPath      string
	WebPath          string
	WebStaticPath    string
	WebTemplatesPath string
	IPLookupURL      string
	Username         string
	PasswordHash     string
	CryptoSecret     string
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
	cryptoSecret := viper.GetString("crypto_secret")

	services := strings.Split(servicesToMonitor, ",")

	return Configuration{
		Port:             port,
		ScriptsPath:      scriptsPath,
		Services:         services,
		WebPath:          webPath,
		WebTemplatesPath: path.Join(webPath, "templates"),
		WebStaticPath:    path.Join(webPath, "static"),
		IPLookupURL:      iplookupURL,
		Username:         username,
		PasswordHash:     passwordHash,
		CryptoSecret:     cryptoSecret,
	}
}

// GetSingle returns the config value at <key>
func GetSingle(key string) string {
	return viper.GetString(key)
}

// Save the config
func Save(options map[string]string) {
	for k, v := range options {
		viper.Set(k, v)
	}

	viper.WriteConfig()
}

func loadDefaults() {

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	viper.SetDefault("port", "8080")
	viper.SetDefault("scripts_path", filepath.Join(path, "scripts"))
	viper.SetDefault("web_path", filepath.Join(path, "web"))
	viper.SetDefault("services_to_monitor", "")
	viper.SetDefault("ip_lookup_url", "https://ipecho.net/plain")
	viper.SetDefault("crypto_secret", rando.GenerateRandomString(20))
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
