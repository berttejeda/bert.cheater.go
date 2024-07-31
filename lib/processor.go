package lib

import (
	"bufio"
	"fmt"
	"os"
	logger "github.com/sirupsen/logrus"
	"path/filepath"
	"regexp"
   utils "berttejeda/cheater/utils"
)

// Cheat file processor
func ProcessCheatFiles(kwargs *kwargs) {

	fileExtensions := kwargs.fileExtensions

	searchPaths := kwargs.searchPaths

	// Compile the regular expression
	re, err := regexp.Compile(`^# `)
	if err != nil {
		logger.Fatalf("error compiling regex: %v", err)
	}	

	// Ensure array of file extensions matches format of .{{ extension }}
  for i := range fileExtensions {
      (fileExtensions)[i] = fmt.Sprintf(".%s", (fileExtensions)[i])
  }

	for _, rootDir := range searchPaths {
	  
	  logger.Debug(fmt.Sprintf("Searching %s for cheat files", rootDir))

		err := filepath.WalkDir(rootDir, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			
			isMatch := utils.ArrayContains(fileExtensions, filepath.Ext(path))

			if !d.IsDir() && isMatch {

				file, err := os.Open(path)
				if err != nil {
					logger.Fatalf("error opening file: %v", err)
				}

				defer file.Close()

				scanner := bufio.NewScanner(file)
				lineNumber := 0
				for scanner.Scan() {
					lineNumber++
					line := scanner.Text()
					// Check if the line matches the regular expression
					if re.MatchString(line) {
						fmt.Println(fmt.Sprintf("%v: %s", lineNumber, line))
					}
				}

				// Check for scanning errors
				if err := scanner.Err(); err != nil {
					logger.Fatalf("error reading file: %v", err)
				}			

			}

			return nil
		})

		if err != nil {
			logger.Error(fmt.Sprintf("Error walking the path %q: %v\n", rootDir, err))
		} 
	
	}

	return
}