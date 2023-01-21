package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type API struct {
	DB        DBInterface
	LogOutput io.Writer
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
	mac := chi.URLParam(r, "mac")
	device, err := a.DB.GetDevice(mac)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("failed to request device from the database"))
		return
	}

	if device == nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(fmt.Sprintf("no device found with MAC %s", mac)))
		return
	}

	sendJSON(device, w)
}

type UpdateDevicePayload struct {
	Firmware string `json:"firmware"`
	Version  string `json:"version"`
}

func (a *API) updateDevice(w http.ResponseWriter, r *http.Request) {
	mac := chi.URLParam(r, "mac")

	if r.ContentLength == 0 {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("device payload required"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("failed to read request body"))
		return
	}

	var payload *UpdateDevicePayload
	err = json.Unmarshal(body, &payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("invalid device payload"))
		return
	}

	err = a.DB.UpdateDevice(mac, payload.Firmware, payload.Version)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("failed to update device in the database"))
		return
	}
}

func (a *API) setName(w http.ResponseWriter, r *http.Request) {
	mac := chi.URLParam(r, "mac")
	val := r.URL.Query().Get("val")
	if val == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("missing device name"))
		return
	}

	err := a.DB.SetDeviceField(mac, "name", val)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("failed to set device name in the database"))
		return
	}
}

func (a *API) setFirmwareType(w http.ResponseWriter, r *http.Request) {
	mac := chi.URLParam(r, "mac")
	val := r.URL.Query().Get("val")
	if val == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("missing device firmware type value"))
		return
	}

	err := a.DB.SetDeviceField(mac, "assignedFirmware", val)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("failed to set device assigned firmware type in the database"))
		return
	}
}

func (a *API) setFirmwareVersion(w http.ResponseWriter, r *http.Request) {
	mac := chi.URLParam(r, "mac")
	val := r.URL.Query().Get("val")
	if val == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("missing device firmware version value"))
		return
	}

	err := a.DB.SetDeviceField(mac, "assignedVersion", val)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("failed to set device assigned firmware version in the database"))
		return
	}
}

func (a *API) GetHttpHandler() http.Handler {
	r := chi.NewRouter()
	loggingMiddleware := middleware.RequestLogger(&middleware.DefaultLogFormatter{
		Logger: log.New(a.LogOutput, "", log.LstdFlags),
	})

	r.Use(loggingMiddleware)
	r.Get("/", a.getDevices)
	r.Get("/{mac}", a.getDevice)
	r.Post("/{mac}", a.updateDevice)
	r.Post("/{mac}/name", a.setName)
	r.Post("/{mac}/firmware", a.setFirmwareType)
	r.Post("/{mac}/version", a.setFirmwareVersion)
	return r
}
