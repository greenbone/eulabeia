package config

type Certificate struct {
	DefaultKeyFile  string // Path to the default location of the private Key
	DefaultCertFile string // Path to the default location of the private Cert
}

type Connection struct {
	Server  string // Bind address of server in format 133.713.371.337:1337
	Timeout int64  // TODO
}

type ScannerPreferences struct {
	ScanInfoStoreTime int64 // TODO
	MaxScan           int64 // TODO
	MaxQueuedScans    int64 // Maximum number of scans that can be queued
}

type Preferences struct {
	LogLevel string // Loglevel (Debug, Info ...)
	LogFile  string // Path to logfile
	Niceness int64  // TODO
}

type Configuration struct {
	Certificate        Certificate
	Connection         Connection
	ScannerPreferences ScannerPreferences
	Preferences        Preferences
}
