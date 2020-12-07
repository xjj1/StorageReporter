package connector

import (
	"errors"

	"github.com/xjj1/StorageReporter/connector/sshcon"
	"github.com/xjj1/StorageReporter/devices"
)

type Connector interface {
	Connect(a *devices.Device) error
	ExecCmd(string) (string, error)
	Close()
}

func Connect(a devices.Device) error {
	if a.Type >= devices.UNKNOWN && a.Type <= devices.PURESTORAGE {
		return sshcon.NewSSH().Connect(&a)
	}

	return errors.New("Can't connect, unknown device type")
}
