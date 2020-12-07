package connector

import (
	"github.com/xjj1/StorageReporter/devices"
)

type Connector interface {
	Connect(a *devices.Device) error
	ExecCmd(string) (string, error)
	Close()
}
