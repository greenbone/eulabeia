package config

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/google/uuid"
)

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

	err := ioutil.WriteFile(path, content, 0644)
	if err != nil {
		t.Errorf("File write should have worked.")
	}

	config := New(path, "eulabeia")

	// Check some Config fields
	if config.Connection.Server != server {
		t.Errorf("Connection.Server should be %s", server)
	}

	if config.Connection.Timeout != timeout {
		t.Errorf("Connection.Timeout should be %d", timeout)
	}

	if config.Sensor.Id != "" {
		t.Errorf("Connection.Sensor.Id should not be set.")
	}

	// Set and check sensor ID in TOML strcut
	config.Sensor.Id = uuid.NewString()
	if config.Sensor.Id == "" {
		t.Errorf("Connection.Sensor.Id should be set.")
	}
	_, err = uuid.Parse(config.Sensor.Id)
	if err != nil {
		t.Errorf("Connection.Sensor.Id should be an uuid.")
	}

	// Save TOML struct back to file
	Save(config)

	// Reload file
	config = New(path, "eulabeia")

	if config.Sensor.Id == "" {
		t.Errorf("Connection.Sensor.Id should be set.")
	}
	_, err = uuid.Parse(config.Sensor.Id)
	if err != nil {
		t.Errorf("Connection.Sensor.Id should be an uuid.")
	}

	// Set director ID in TOML strcut
	config.Director.Id = uuid.NewString()
	if config.Director.Id == "" {
		t.Errorf("Connection.Director.Id should be set.")
	}
	_, err = uuid.Parse(config.Director.Id)
	if err != nil {
		t.Errorf("Connection.Director.Id should be an uuid.")
	}

	os.Remove(path)
}
