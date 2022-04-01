package lib

type Device struct {
	MAC               string `json:"mac" redis:"mac"`
	CurrentFirmware   string `json:"currentFirmware" redis:"currentFirmware"`
	CurrentVersion    string `json:"currentVersion" redis:"currentVersion"`
	AssignedFirmware  string `json:"assignedFirmware" redis:"assignedFirmware"`
	AssignedVersion   string `json:"assignedVersion" redis:"assignedVersion"`
	AcceptsPrerelease bool   `json:"acceptsPrerelease" redis:"acceptsPrerelease"`
}
