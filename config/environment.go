package config

import (
	"os"
	"strings"
)

func server(server string, c *Configuration) {
	c.Connection.Server = server
}

// lookup table that binds an environment variable to a function that overrides the configuration variable in the config file struct
var lookup = map[string]func(string, *Configuration){
	"MQTT_SERVER": server,
}

// OverrideViaENV overrides configuration settings with environment variables, if they are set.
//
// Uses a defined lookup table to identify and get environment variables to override.
func OverrideViaENV(c *Configuration) {
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if f, ok := lookup[pair[0]]; ok {
			f(pair[1], c)
		}
	}
}
