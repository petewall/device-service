package deviceservice

import (
	"encoding/json"
	"net/http"
)

type API struct {
	DB DBInterface
}

func (a *API) getDevices(w http.ResponseWriter, r *http.Request) {
	devices, _ := a.DB.GetDevices()
	// if err != nil {

	// }

	encoded, _ := json.Marshal(devices)
	// if err != nil {

	// }

	_, _ = w.Write([]byte(encoded))
}

func (a *API) GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", a.getDevices)
	return mux
}
