package db

import (
	"log"

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

func (r *repo) ListArrays() ([]devices.Device, error) {
	var A []devices.Device
	rows, err := r.db.Query("select ArrayType, Cluster, Name, Friendlyname, Username, Password from Arrays")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var at devices.DeviceType
		var Cluster, Name, Friendlyname, Username, Password string
		err = rows.Scan(&at, &Cluster, &Name, &Friendlyname, &Username, &Password)
		if err != nil {
			log.Fatalf("(listedevices) Error in DB %v", err)
		}
		A = append(A, devices.Device{
			Type:         at,
			Cluster:      Cluster,
			Name:         Name,
			Friendlyname: Friendlyname,
			Username:     Username,
			Password:     Password})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return A, nil
}

func (r *repo) ListArraysNames() ([]string, error) {
	A, err := r.ListArrays()
	if err != nil {
		return []string{}, err
	}
	var ret []string
	for _, v := range A {
		ret = append(ret, v.Type.String()+` `+v.Name+` `+v.Friendlyname+` `+v.Cluster)
	}
	return ret, nil
}
