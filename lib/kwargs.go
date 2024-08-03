package lib

import (
	utils "berttejeda/cheater/utils"
)

type kwargs struct {
    // Required
    topics []string
    searchPaths  []string
    fileExtensions []string
    // Optional
    pauseBetweenTopics bool
}

// Each optional attribute will have its own public method
func (c *kwargs) WithPause(pauseBetweenTopics bool) *kwargs {
    c.pauseBetweenTopics = pauseBetweenTopics
    return c
}

func (c *kwargs) WithFileExtensions(r []string) *kwargs {
    c.fileExtensions = r
    return c
}

func (c *kwargs) WithSearchPaths(default_paths []string, paths []string) *kwargs {
    
		var searchPaths []string

		if len(default_paths) > 0 {
			searchPaths = default_paths
		} else {
			searchPaths = paths
		}

    for i := range searchPaths {
        (searchPaths)[i], _ = utils.FSOExpandUser((searchPaths)[i])
    }

    c.searchPaths = searchPaths
    return c
}

// This only accepts the required options as params
func InitKwargs(topics []string) *kwargs {
    // First fill in the options with default values
    return &kwargs{topics, []string{"."}, []string{".*"}, true}
}

func Do(c *kwargs) {}