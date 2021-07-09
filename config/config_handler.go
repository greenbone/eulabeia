package config

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/pelletier/go-toml"
)

type ConfigurationHandler struct {
	Configuration Configuration // The configuration object
	module        string        // The module name
	path          string        // Path to the configuration file
}

// Check if file exists
// Returns true if file exists, false else
func fileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		// file exists
		return true
	}
	return false
}

// Look up the config file in the order
// 1. given by parameter --config
// 2. custom user config file in home
// 3. in /usr/etc or /etc/ ...
// Returns the path to the first found file
func findConfigFile(path string, module string) string {
	if path != "" && fileExists((path)) {
		return path
	}
	// look in the default paths
	var defaultPaths = [...]string{
		os.Getenv("HOME") + "/.config",
		"/usr/etc",
		"/etc",
	}
	for _, path := range defaultPaths {
		path += "/" + module + "/config.toml"
		if fileExists(path) {
			// file exists
			return path
		}
	}
	panic(errors.New("no config file found"))
}

// Load the config file after startup
func (c *ConfigurationHandler) Load(path string, module string) {
	c.module = module
	c.Configuration = Configuration{}
	c.path = findConfigFile(path, module)

	// Read the config file
	bytes, err := ioutil.ReadFile(c.path)
	if err != nil {
		panic(err)
	}

	toml.Unmarshal(bytes, c.Configuration)
}
