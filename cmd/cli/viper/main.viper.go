package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct{
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct{
		User string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host string `mapstructure:"host"`
	} `mapstructure:"databases"`
}

func main(){
	viper := viper.New()
	viper.AddConfigPath("./config/") // path to config file
	viper.SetConfigName("local") // tên file
	viper.SetConfigType("yaml") 

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read configuration: %w \n", err))
	}

	// read server configuration
	
	fmt.Println("Server port:: ", viper.GetInt("server.port"))
	fmt.Println("Server port:: ", viper.GetString("security.jwt.key"))
}