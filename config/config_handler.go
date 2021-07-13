package config

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/pelletier/go-toml"
)

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

// Save the Configuration in its current state
// into the stored file path
func Save(c *Configuration) {
	bytes, err := toml.Marshal(c)
	if err != nil {
		// TODO error handling
		panic(err)
	}
	err = ioutil.WriteFile(c.path, bytes, 0644) //TODO what permissions for the file?
	if err != nil {
		// TODO error handling
		panic(err)
	}
}

// Looks for the configuration file and
// returns a filled Configuration struct
func New(path string, module string) *Configuration {
	c := Configuration{}
	c.path = findConfigFile(path, module)

	// Read the config file
	bytes, err := ioutil.ReadFile(c.path)
	if err != nil {
		panic(err)
	}

	toml.Unmarshal(bytes, &c)
	if c.Context == "" {
		c.Context = module
	}

	return &c
}
