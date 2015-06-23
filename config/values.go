package config

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"tribes/tools"

	"github.com/spf13/viper"
)

var (
	def_TLSPORT       string = "21000"
	def_TribeID       string = "AdzfNdsMAajMMuPpVsNXvWWxIDohwppz"
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
		viper.SetDefault("MyTribeID", def_TribeID)
		viper.SetDefault("BootStrapHost", def_BootstrapHost)
		viper.SetDefault("BootStrapPort", def_BootstrapPort)

		os.MkdirAll(tools.ConfigPath, 0755)
		f, _ := os.Create(tools.ConfigFile)
		fmt.Fprintf(f, "TLSPORT = %q\r\n", def_TLSPORT)
		fmt.Fprintf(f, "MyTribeID = %q\r\n", def_TribeID)
		fmt.Fprintf(f, "MyBootStrapHost = %q\r\n", def_BootstrapHost)
		fmt.Fprintf(f, "MyBootStrapPort = %q\r\n", def_BootstrapPort)
		f.Close()
	}
}

func GetClusterPort() int {
	return viper.GetInt("TLSPORT")
}

func GetTribeID() string {
	return viper.GetString("MyTribeID")
}

func GetBootStrapHost() string {
	host := viper.GetString("MyBootStrapHost")

	conn, err := net.Dial("udp", host+":80")
	if err != nil {
		log.Printf("[CONFIG] SYSADMIIIIIN : cannot use UDP")
		return "127.0.0.1"
	}
	conn.Close()
	torn := strings.Split(conn.RemoteAddr().String(), ":")
	return torn[0]

}

func GetBootStrapPort() int {
	return viper.GetInt("MyBootStrapPort")
}
