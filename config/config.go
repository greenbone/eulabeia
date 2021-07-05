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
func findConfigFile(path string) string {
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
		path += "/eulabeia/config.toml"
		if fileExists(path) {
			// file exists
			return path
		}
	}
	panic(errors.New("no config file found"))
}

func Overwrite(config_map *toml.Tree, key string, value string) {
	//TODO overwrite at runtime given arguments
	config_map.Set(key, value)
}

// Load the config file after startup
// In the returned config map config values can
// be accessed with Get("section.key").(string)
// The values can also be overwritten at runtime
// Returns the config map
func Load(path string) *toml.Tree {
	config_path := findConfigFile(path)

	// Read the config file
	byte, err := ioutil.ReadFile(config_path)
	if err != nil {
		panic(err)
	}

	tree, err := toml.Load(string(byte))
	if err != nil {
		panic(err)
	}
	return tree
}
