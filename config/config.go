package config

type Certificate struct {
	DefaultKeyFile  string // Path to the default location of the private Key
	DefaultCertFile string // Path to the default location of the private Cert
}

type Connection struct {
	Server  string	// Bind address of server in format 133.713.371.337:1337
	Timeout int64	// TODO
}

type ScannerPreferences struct {
	scanInfoStoreTime int64 // TODO
	maxScan           int64 // TODO
	maxQueuedScans    int64 // Maximum number of scans that can be queued
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

