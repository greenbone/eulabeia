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
	ScanInfoStoreTime   int64  // Time (h) before a scan is considere forgotten
	MaxScan             int64  // Maxi number of parallel scans
	MaxQueuedScans      int64  // Maxi number of scans that can be queued
	Niceness            int64  // Niceness of the openvas Process
	MinFreeMemScanQueue uint64 // Min Memory necessary for a Scan to start
}

type Preferences struct {
	LogLevel string // Loglevel (Debug, Info ...)
	LogFile  string // Path to logfile
}

type Sensor struct {
	Id string // The Id (a uuid) of this sensor
}

type Director struct {
	Id          string // The Id (a uuid) of this director
	StoragePath string // The path to store the json into
}

type Configuration struct {
	Context            string
	Certificate        Certificate
	Connection         Connection
	ScannerPreferences ScannerPreferences
	Preferences        Preferences
	Sensor             Sensor
	Director           Director
	path               string
}
