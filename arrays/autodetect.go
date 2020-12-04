package arrays

import (
	"errors"

	"github.com/xjj1/StorageReporter/connector"
	"github.com/xjj1/StorageReporter/devices"
)

func AutoDetect(a devices.Device) (devices.DeviceType, error) {
	for i := devices.HP3PAR; i <= devices.PURESTORAGE; i++ {
		a.Type = i
		_, err := connector.Connect(a)
		if err == nil {
			return i, nil
		}
	}
	return devices.UNKNOWN, errors.New("Can't connect / unsupported device")
}
