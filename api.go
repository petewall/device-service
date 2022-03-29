package main

import "net/http"

type API struct {
}

func (a *API) GetDevices(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("[]"))
}

func (a *API) GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", a.GetDevices)
	return mux
}
