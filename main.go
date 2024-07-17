package main

import (
		"fmt"
    "github.com/akamensky/argparse"
    logger "github.com/sirupsen/logrus"
    "github.com/spf13/viper"
    "os"
    // "strings"
    // "reflect"
    "strconv"
)

type Search struct {
 Paths []string `mapstructure:"paths"`
 Filters []string `mapstructure:"filters"`
}

type Config struct {
 Search Search  `mapstructure:"search"`
}

func all(b ...any) bool {

	if len(b) == 0 {
		return false
	} else if len(b) == 1 {
	  switch v := b[0].(type) {
	  	case string:
	  	  bs0, err := strconv.ParseBool(b[0].(string))
	  	  if err != nil {
	  	  	return len(b[0].(string)) > 0 
	  	  } else {
	  	  	return bs0
	  	  }
	  	case int:
	  		return b[0].(int) > 0
	  	default:
	  	  return v.(bool)
	  }	
	}

  switch v := b[0].(type) {
    case string:
      bs0, err := strconv.ParseBool(b[0].(string))
      if err != nil {
      	return len(b[0].(string)) > 0 && all(b[1:]...)
      } else {
      	return bs0 && all(b[1:]...)
      }
    case int:
      return b[0].(int) > 0 && all(b[1:]...)
    default:
      return v.(bool) && all(b[1:]...)
  }		
	
}

func main() {

	// Parse CLI Commands & Arguments
	parser := argparse.NewParser("bt-cheater", "Display your markdown notes according to keywords")
	
	// Add global flag for debug logging
	enableDebugLogging := parser.Flag("d", "debug", &argparse.Options{Required: false, Help: "Enable debug logging", Default: false})	
	
	// Add top level command `find`
	findCmd := parser.NewCommand("find", "Retrieve cheat notes from specified cheatfiles according to keywords")
	topics := findCmd.String("s", "session-path", &argparse.Options{Help: "Provide a session cookie file path"})
	cheatFile := findCmd.String("c", "--cheat-file", &argparse.Options{Required: false, Help: "Manually specify cheat file(s) to search against"})
	cheatFileSearchPaths := findCmd.String("p", "--cheat-file-search-paths", &argparse.Options{Required: false, Help: "Manually specify cheat file paths to search against"})
	explodeTopics := findCmd.Flag("x", "explode-topics", &argparse.Options{Required: false, Help: "Write results to their own cheat files", Default: false})	
	
	// Pass our args to the parser
	err := parser.Parse(os.Args)

	// Handle argparse errors
	if err != nil {
	fmt.Print(parser.Usage(err))
	os.Exit(1)

	}	

	// Enable Debug Logging if applicable
	if *enableDebugLogging {
		logger.SetLevel(logger.DebugLevel)
	}

	// Enable Debug Logging if applicable
	logger.SetFormatter(&logger.JSONFormatter{})

	// Config
	viper.SetConfigName("cheater") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./config/") // config file path
	viper.AutomaticEnv()             // read value ENV variable

	err = viper.ReadInConfig()

	if err != nil {
	  logger.Warning("Config file: default \n", err)
	  // os.Exit(1)
	}

	var config Config

	if err := viper.Unmarshal(&config); err != nil {
  	fmt.Println(err)
  	return
 	}

	if findCmd.Happened() {

		// fmt.Printf("%v", *parser.GetArgs()[0].GetParsed())

		if *explodeTopics {
			logger.Info(config.Search.Paths)
		}

		if len(*cheatFile) > 0 {

		}

		if len(*cheatFileSearchPaths) > 0 {

		}
		// for _, arg := range findCmd.GetArgs() {
		// 	fmt.Println(arg)
		// }		
	}	

}