package devices

func (d *Device) String() string {
	if d.Cluster != "" {
		return d.Cluster
	} else {
		return d.Name
	}
}
