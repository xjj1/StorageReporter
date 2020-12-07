package arrays

import (
	"log"
	"strings"

	"github.com/pkg/errors"
	"github.com/xjj1/StorageReporter/connector"
	"github.com/xjj1/StorageReporter/devices"
)

func AutoDetect(c connector.Connector, a *devices.Device) error {

	if err := c.Connect(a); err == nil {
		// if err != nil {
		// 	return errors.New("AutoDetect:")
		// }

		var res string

		// check if 3PAR :
		res, err := c.ExecCmd("showversion")
		if err == nil && strings.Contains(res, "Release version") {
			a.Type = devices.HP3PAR
			log.Println("Detected 3PAR")
			return nil
		}

		// check if MSA :
		res, err = c.ExecCmd("show version")
		if err == nil && strings.Contains(res, "Controller") {
			a.Type = devices.HPMSA
			return nil
		}

		// check if PURESTORAGE :
		res, err = c.ExecCmd("purevol list")

		if err == nil && strings.Contains(res, "Name") {
			a.Type = devices.PURESTORAGE
			log.Println("Detected PURESTORAGE")
			return nil
		}

		c.Close()
	}
	// check if Nimble
	tmpType := a.Type // save the original type // should be devices.UNKNOWN
	// force it to use the Nimble ssh
	a.Type = devices.HPNIMBLE

	err := c.Connect(a)
	if err != nil {
		a.Type = tmpType
		return errors.Wrap(err, "Checking if Nimble")
	}
	res, err := c.ExecCmd("pool --list")
	if err == nil && strings.Contains(res, "--------------+") {
		return nil
	}
	// if not Nimble return the original type
	a.Type = tmpType

	return errors.New("no valid device found")
}
