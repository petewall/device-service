package lib

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type API struct {
	DB DBInterface
}

func sendJSON(object interface{}, w http.ResponseWriter) {
	encoded, err := json.Marshal(object)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("failed to prepare the list of devices"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(encoded)
}

func (a *API) getDevices(w http.ResponseWriter, r *http.Request) {
	devices, err := a.DB.GetDevices()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("failed to request devices from the database"))
		return
	}

	sendJSON(devices, w)
}

func (a *API) getDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	device, err := a.DB.GetDevice(vars["mac"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("failed to request device from the database"))
		return
	}

	if device == nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(fmt.Sprintf("no device found with MAC %s", vars["mac"])))
		return
	}

	sendJSON(device, w)
}

func (a *API) GetMux() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", a.getDevices).Methods("GET")
	r.HandleFunc("/{mac}", a.getDevice).Methods("GET")
	return r
}
