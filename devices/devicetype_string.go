// Code generated by "stringer -type DeviceType types.go"; DO NOT EDIT.

package devices

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[UNKNOWN-0]
	_ = x[HP3PAR-1]
	_ = x[HPMSA-2]
	_ = x[HPNIMBLE-3]
	_ = x[PURESTORAGE-4]
}

const _DeviceType_name = "UNKNOWNHP3PARHPMSAHPNIMBLEPURESTORAGE"

var _DeviceType_index = [...]uint8{0, 7, 13, 18, 26, 37}

func (i DeviceType) String() string {
	if i < 0 || i >= DeviceType(len(_DeviceType_index)-1) {
		return "DeviceType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DeviceType_name[_DeviceType_index[i]:_DeviceType_index[i+1]]
}
