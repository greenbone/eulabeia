package scanner

type Scanner interface {
	StartScan(scan string, niceness int) error
	StopScan(scan string) error
	ScanFinished(scan string) error
	GetVersion() (string, error)
	GetSettings() (map[string]string, error)
}
