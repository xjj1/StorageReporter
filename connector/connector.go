package connector

import (
	"github.com/xjj1/StorageReporter/connector/sshcon"
	"github.com/xjj1/StorageReporter/devices"
)

type Connector interface {
	ExecCmd(string) (string, error)
}

func Connect(a devices.Device) (Connector, error) {
	switch a.Type {
	case devices.HP3PAR:
		return sshcon.NewSSH(a)
	}
}
