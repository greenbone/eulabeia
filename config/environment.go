package config

import (
	"os"
	"strings"
)

func server(server string, c *Configuration) {
	c.Connection.Server = server
}

var lookup = map[string]func(string, *Configuration){
	"MQTT_SERVER": server,
}

func OverrideEnvConfiguration(c *Configuration) {
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if f, ok := lookup[pair[0]]; ok {
			f(pair[1], c)
		}
	}
}
