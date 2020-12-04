package arrays

import (
	"errors"

	"github.com/xjj1/StorageReporter/connector"
	"github.com/xjj1/StorageReporter/devices"
)

func AutoDetect(a devices.Device) (devices.DeviceType, error) {
	for _, t := range []devices.DeviceType{devices.HP3PAR, devices.HPMSA, devices.HPNIMBLE, devices.PURESTORAGE} {
		a.Type = t
		_, err := connector.Connect(a)
		if err == nil {
			return t, nil
		}
	}

	return devices.UNKNOWN, errors.New("Can't connect / unsupported device")
}
