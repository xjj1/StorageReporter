package db

import (
	"github.com/pkg/errors"
	"github.com/xjj1/StorageReporter/devices"
)

func (r *repo) AddArray(a *devices.Device) error {
	_, err := r.db.Exec(`insert into Arrays(
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
