package repository

import "github.com/xjj1/StorageReporter/devices"

type Email struct {
	Rcptto, Mailserver, Mailfrom, Subject, Username, Password string
}

type Repository interface {
	AddArray(a devices.Device) error
	ListArrays() ([]devices.Device, error)
	ListArraysNames() ([]string, error)
	AddEmailSettings(param []string) error
	GetEmail() (*Email, error)
	Close()
}
