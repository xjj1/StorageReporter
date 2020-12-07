package devices

type DeviceType int

const (
	UNKNOWN DeviceType = iota
	HP3PAR
	HPMSA
	HPNIMBLE
	PURESTORAGE
)

type Device struct {
	Type         DeviceType
	Cluster      string
	Name         string
	Friendlyname string
	Username     string
	Password     string
}

type Disk struct {
	Size        int64
	Used        int64
	Snapshot    int64
	Servers     []string
	Usercpg     string
	Wwn         string
	Rcopystatus string
}

type FreeSpaceData struct {
	StorageType      string
	TotalSpace       int64
	FreeSpace        int64
	EstFreeSpace     int64
	UsedPerc         int64
	Snapshot         int64
	Presented_size   int64
	Oversubscription float64
}

/*
type HistEntry struct {
	datetime       string
	array          string
	disktype       string
	allsize        string
	freesize       string
	est_free_size  string
	used_perc      string
	snapshots      string
	presented_size string
}*/

type DeviceData struct {
	DeviceName      string
	Friendlyname    string
	Luns            map[string]Disk
	FreeSpaceArray  []FreeSpaceData
	Hosts           []string
	Error           error
	Checkhealth     []string
	ShowhostPathsum []string
	ShowhostPersona []string
}
