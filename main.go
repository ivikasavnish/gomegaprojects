package main

import (
	"github.com/spf13/viper"
)

func init() {
	// godotenv.Load(".env")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
}
func main() {

}
