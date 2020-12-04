package db

import (
	"github.com/pkg/errors"
	"github.com/xjj1/StorageReporter/arrays"
	"github.com/xjj1/StorageReporter/devices"
)

func (r *repo) AddArray(a devices.Device) error {
	var err error
	a.Type, err = arrays.AutoDetect(a)
	if err != nil {
		return errors.New("Cannot autodetect / unknown array")
	}
	_, err = r.db.Exec(`insert into Arrays(
		ArrayType,
		Cluster,
		Name,
		Friendlyname,
		Username,
		Password
	)
	values (?, ?, ?, ?, ?, ?)`,
		a.Type, a.Cluster, a.Name, a.Friendlyname, a.Username, a.Password)
	if err != nil {
		return errors.Wrap(err, "add array")
	}
	return nil
}
