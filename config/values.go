package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net"
	"os"
	"tribes/tools"
)

var (
	def_TLSPORT       string = "21000"
	def_TribeID       string = "AdzfNdsMAajMMuPpVsNXvWWxIDohwppz"
	def_MyPublicHost  string = "whatever.example.com"
	def_BootstrapHost string = "127.0.0.1"
	def_BootstrapPort string = "21000"
)

func init() {

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(tools.ConfigPath)

	err := viper.ReadInConfig() // Find and read the config file

	if err != nil { // Handle errors reading the config file
		log.Printf("[OMG] Cannot read config file : %s", err)
		viper.SetDefault("TLSPORT", def_TLSPORT)
		viper.SetDefault("MyPublicHost", def_MyPublicHost)
		viper.SetDefault("MyTribeID", def_TribeID)
		viper.SetDefault("BootStrapHost", def_BootstrapHost)
		viper.SetDefault("BootStrapPort", def_BootstrapPort)

		os.MkdirAll(tools.ConfigPath, 0755)
		f, _ := os.Create(tools.ConfigFile)
		_, _ = fmt.Fprintf(f, "TLSPORT = %q\r\n", def_TLSPORT)
		_, _ = fmt.Fprintf(f, "MyTribeID = %q\r\n", def_TribeID)
		_, _ = fmt.Fprintf(f, "MyPublicHost = %q\r\n", def_MyPublicHost)
		_, _ = fmt.Fprintf(f, "MyBootStrapHost = %q\r\n", def_BootstrapHost)
		_, _ = fmt.Fprintf(f, "MyBootStrapPort = %q\r\n", def_BootstrapPort)
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

func GetBootStrapHost() string {
	return viper.GetString("MyBootStrapHost")
}

func GetBootStrapPort() int {
	return viper.GetInt("MyBootStrapPort")
}

func GetIPFromHostname(hostname string) string {

	address, err := net.LookupIP(hostname)
	if err != nil {
		log.Printf("[OMG] cant' resolve %s : %s", hostname, err)
	}

	return address[0].String()

}
