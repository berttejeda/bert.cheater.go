package main

import (
		"fmt"
    "os"
    config "berttejeda/cheater/config"
    lib "berttejeda/cheater/lib"
    // commands "berttejeda/cheater/commands"
    "github.com/alecthomas/kingpin/v2"
    logger "github.com/sirupsen/logrus"
    // "strings"
    // "reflect"
)

var (
  app         = kingpin.New("bt-cheater", "Search through your markdown notes by keyword")
  verbose     = app.Flag("verbose", "Enable verbose mode").Short('v').Bool()
  debug       = app.Flag("debug", "Enable debug mode").Short('d').Bool()
  jsonLogging = app.Flag("json-logging", "Enable json log format").Short('J').Bool()
  find        = app.Command("find", "Retrieve cheat notes and display in terminal")
  filters     = find.Flag("filters", "File extensions to math when searching").Short('f').Default("md", "txt").Strings()
  paths       = find.Flag("paths", "File search paths").Short('p').Default(".").Strings()
  topics      = find.Arg("args", "Topics to match").Strings()
)

func main() {

	cli := kingpin.MustParse(app.Parse(os.Args[1:]))

	// Enable Debug Logging if applicable
	if *debug {
		logger.SetLevel(logger.DebugLevel)
	}

	config, err := config.Init()

	if err != nil {
	    logger.Error(err)
	} 		

	// Enable JSON Logging if applicable
	if *jsonLogging {
		logger.SetFormatter(&logger.JSONFormatter{})
	}

  switch cli {

  // Find cheat notes
  case find.FullCommand():

		if *filters == nil {
			fmt.Println("nil!")
		}
		
  	kwargs := lib.InitKwargs(*topics).WithFileExtensions(*filters).WithSearchPaths(config.Search.Paths, *paths)

    lib.ProcessCheatFiles(kwargs)

	}

}