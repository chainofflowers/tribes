package config

import (
	"../tools/"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

func init() {

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("$HOME/News/")

	err := viper.ReadInConfig() // Find and read the config file

	if err != nil { // Handle errors reading the config file
		log.Printf("[OMG] Cannot read config file : %s", err)
		viper.SetDefault("DHTPORT", "30000")
		viper.SetDefault("BootstrapNode", "boseburo.ddns.net:30000")

		var user_home = tools.GetHomeDir()
		config_file := filepath.Join(user_home, "News", "config.toml")
		os.MkdirAll(filepath.Join(user_home, "News"), 0755)
		f, _ := os.Create(config_file)
		_, _ = f.WriteString("DHTPORT = \"21000\"\n")
		_, _ = f.WriteString("BootstrapNode = \"boseburo.ddns.net:30000\"\n")
		_, _ = f.WriteString("MyPublicHost = \"whatever.ddns.net\"\n")
		f.Close()
	}
}

func GetClusterPort() int {
	return viper.GetInt("DHTPORT")
}

func GetBootstrapNode() string {
	return viper.GetString("BootStrapNode")
}

func GetPublicHost() string {
	return viper.GetString("MyPublicHost")
}
