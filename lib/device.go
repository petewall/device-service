package lib

type Device struct {
	MAC              string `json:"mac" redis:"mac"`
	Name             string `json:"name" redis:"name"`
	Firmware         string `json:"firmware" redis:"firmware"`
	Version          string `json:"version" redis:"version"`
	AssignedFirmware string `json:"assignedFirmware" redis:"assignedFirmware"`
	AssignedVersion  string `json:"assignedVersion" redis:"assignedVersion"`
	LastUpdate       int64  `json:"lastUpdate" redis:"lastUpdate"`
}

type UpdateDevicePayload struct {
	Firmware string `json:"firmware"`
	Version  string `json:"version"`
}
