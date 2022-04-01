package lib

import (
	"encoding/json"
	"net/http"
)

type API struct {
	DB DBInterface
}

func (a *API) getDevices(w http.ResponseWriter, r *http.Request) {
	devices, err := a.DB.GetDevices()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to request devices from the database"))
		return
	}

	encoded, err := json.Marshal(devices)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to prepare the list of devices"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(encoded)
}

func (a *API) GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", a.getDevices)
	return mux
}
