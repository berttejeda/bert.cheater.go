package options

import (
  "fmt"
	logger "github.com/sirupsen/logrus"
  "github.com/spf13/viper"
  utils "berttejeda/cheater/utils"
)

type Search struct {
 Paths []string `mapstructure:"paths"`
 Filters []string `mapstructure:"filters"`
}

type Options struct {
 Search Search  `mapstructure:"search"`
 PauseBetweenTopics bool `mapstructure:"pause"`
}

func InitOptions() (Options, error) {

	viper.SetConfigName("cheater") // options file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./options/") // options file path
	viper.AutomaticEnv()             // read value ENV variable

	err := viper.ReadInConfig()

	logger.Debug("Using options file: ", viper.ConfigFileUsed())


	if err != nil {
	  logger.Warning("Failed to parse options file,error was ", err, " Using defaults")
	}

	var options Options

	if err := viper.Unmarshal(&options); err != nil {
		fmt.Println(err)
		return options, err
	}

	return options, err

}

type Config struct {
    // Required
    Topics []string
    SearchPaths  []string
    FileExtensions []string
    // Optional
    PauseBetweenTopics bool
}

// Each optional attribute will have its own public method
func (c *Config) WithPause(pauseBetweenTopics bool) *Config {
    c.PauseBetweenTopics = pauseBetweenTopics
    return c
}

func (c *Config) WithFileExtensions(r []string) *Config {
    c.FileExtensions = r
    return c
}

func (c *Config) WithSearchPaths(default_paths []string, paths []string) *Config {
    
    var searchPaths []string

    if len(default_paths) > 0 {
      searchPaths = default_paths
    } else {
      searchPaths = paths
    }

    for i := range searchPaths {
        (searchPaths)[i], _ = utils.FSOExpandUser((searchPaths)[i])
    }

    c.SearchPaths = searchPaths
    return c
}

// This only accepts the required options as params
func InitConfig(topics []string) *Config {
    // First fill in the options with default values
    return &Config{topics, []string{"."}, []string{".*"}, true}
}