package utils

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func FSOExpandUser(path string) (string, error) {
    if strings.HasPrefix(path, "~") {

    	  dirname, _ := os.UserHomeDir()
    		path = filepath.Join(dirname, path[2:])	
    }
    return path, nil
}

func FSOGlob(root string, fn func(string)bool) []string {
   var files []string
   filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
      if fn(s) {
         files = append(files, s)
      }
      return nil
   })
   return files
}