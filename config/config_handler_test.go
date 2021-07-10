package config

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/google/uuid"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func writeFile(path string, content []byte) {
	err := ioutil.WriteFile(path, content, 0644)
	check(err)
}

func TestConfigurationHandler(t *testing.T) {
	content := []byte(`[Certificate]
defaultKeyFile = "/usr/var/lib/gvm/private/CA/serverkey.pem"
defaultCertFile = "/usr/var/lib/gvm/CA/servercert.pem"
defaultCaFile = "/usr/var/lib/gvm/CA/cacert.pem"

[Connection]
server = "133.713.371.337:1337"
timeout = 10

[ScannerPreferences]
scanInfoStoreTime = 0
maxScan = 0
maxQueuedScans = 0

[Preferences]
logFile = ""
logLevel = ""
niceness = 10`)

	path := "./config.toml"
	server := "133.713.371.337:1337"
	timeout := int64(10)

	writeFile(path, content)

	confHandler := ConfigurationHandler{}

	confHandler.Load(path, "eulabeia")

	if confHandler.Configuration.Connection.Server != server {
		t.Errorf("Connection.Server should be %s", server)
	}

	if confHandler.Configuration.Connection.Timeout != timeout {
		t.Errorf("Connection.Timeout should be %d", timeout)
	}

	if confHandler.Configuration.Sensor.Id != "" {
		t.Errorf("Connection.Sensor.Id should not be set.")
	}
	confHandler.SetId("sensor")
	if confHandler.Configuration.Sensor.Id == "" {
		t.Errorf("Connection.Sensor.Id should be set.")
	}
	_, err := uuid.Parse(confHandler.Configuration.Sensor.Id)
	if err != nil {
		t.Errorf("Connection.Sensor.Id should be an uuid.")
	}

	os.Remove(path)
}
