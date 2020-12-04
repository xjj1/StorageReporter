package connector

import (
	"errors"

	"github.com/xjj1/StorageReporter/connector/sshcon"
	"github.com/xjj1/StorageReporter/devices"
)

type Connector interface {
	ExecCmd(string) (string, error)
}

func Connect(a devices.Device) (Connector, error) {
	switch a.Type {
	case devices.HP3PAR:
		return sshcon.NewSSH(&a)
	case devices.HPNIMBLE:
		return sshcon.NewSSHNimble(&a)
	case devices.HPMSA:
		return sshcon.NewSSH(&a)
	case devices.PURESTORAGE:
		return sshcon.NewSSH(&a)
	default:
		return nil, errors.New("Can't connect, unknown device type")
	}
}
