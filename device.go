package deviceservice

type Device struct {
	MAC               string `json:"mac"`
	CurrentFirmware   string `json:"currentFirmware"`
	CurrentVersion    string `json:"currentVersion"`
	AssignedFirmware  string `json:"assignedFirmware"`
	AssignedVersion   string `json:"assignedVersion"`
	AcceptsPrerelease bool   `json:"acceptsPrerelease"`
}

type DeviceUpdate struct {
	MAC             string
	Date            string
	CurrentFirmware string
	CurrentVersion  string
	UpdateSent      bool
	SentFirmware    string
	SentVersion     string
}
