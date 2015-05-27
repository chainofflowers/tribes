package config

import (
	"../tools/"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

var (
	def_TLSPORT      string = "21000"
	def_TribeID      string = "AdzfNdsMAajMMuPpVsNXvWWxIDohwppz"
	def_MyPublicHost string = "whatever.example.com"
)

func init() {

	var user_home = tools.GetHomeDir()
	config_path := filepath.Join(user_home, "News")
	config_file := filepath.Join(user_home, "News", "config.toml")

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(config_path)

	err := viper.ReadInConfig() // Find and read the config file

	if err != nil { // Handle errors reading the config file
		log.Printf("[OMG] Cannot read config file : %s", err)
		viper.SetDefault("TLSPORT", def_TLSPORT)
		viper.SetDefault("MyPublicHost", def_MyPublicHost)
		viper.SetDefault("MyTribeID", def_TribeID)

		os.MkdirAll(config_path, 0755)
		f, _ := os.Create(config_file)
		_, _ = f.WriteString("TLSPORT = \"" + def_TLSPORT + "\"\n")
		_, _ = f.WriteString("MyTribeID = \"" + def_TribeID + "\"\n")
		_, _ = f.WriteString("MyPublicHost = \"" + def_MyPublicHost + "\"\n")
		f.Close()
	}
}

func GetClusterPort() int {
	return viper.GetInt("TLSPORT")
}

func GetTribeID() string {
	return viper.GetString("MyTribeID")
}

func GetPublicHost() string {
	return viper.GetString("MyPublicHost")
}
