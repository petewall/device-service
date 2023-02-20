package lib

import "github.com/coreos/go-semver/semver"

type Device struct {
	MAC              string `json:"mac" redis:"mac"`
	Name             string `json:"name" redis:"name"`
	Firmware         string `json:"firmware" redis:"firmware"`
	Version          string `json:"version" redis:"version"`
	AssignedFirmware string `json:"assignedFirmware" redis:"assignedFirmware"`
	AssignedVersion  string `json:"assignedVersion" redis:"assignedVersion"`
	LastUpdate       int64  `json:"lastUpdate" redis:"lastUpdate"`
}

func (d *Device) IsDifferent(firmwareType, firmwareVersion string) bool {
	return d.Firmware != firmwareType || d.Version != firmwareVersion
}

func (d *Device) IsOlderThan(version string) bool {
	myVersion, _ := semver.NewVersion(d.Version)
	otherVersion, _ := semver.NewVersion(version)
	return myVersion.LessThan(*otherVersion)
}

type UpdateDevicePayload struct {
	Firmware string `json:"firmware"`
	Version  string `json:"version"`
}
