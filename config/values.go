package config

import (
	
	"github.com/spf13/viper"
	"log"
	
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
	}
}

func GetClusterPort() int {
	return viper.GetInt("DHTPORT")
}

func GetBootstrapNode() string {
	return viper.GetString("BootStrapNode")
}
