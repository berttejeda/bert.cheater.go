package config

import (
  "fmt"
	logger "github.com/sirupsen/logrus"
  "github.com/spf13/viper"
)

type Search struct {
 Paths []string `mapstructure:"paths"`
 Filters []string `mapstructure:"filters"`
}

type Config struct {
 Search Search  `mapstructure:"search"`
 PauseBetweenTopics bool `mapstructure:"pause"`
}

func Init() (Config, error) {

	viper.SetConfigName("cheater") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./config/") // config file path
	viper.AutomaticEnv()             // read value ENV variable

	err := viper.ReadInConfig()

	logger.Debug("Using config file: ", viper.ConfigFileUsed())


	if err != nil {
	  logger.Warning("Failed to parse config file,error was ", err, " Using defaults")
	}

	var config Config

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(err)
		return config, err
	}

	return config, err

}

