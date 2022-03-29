package main

import (
	"net/http"
)

func main() {
	api := &API{}

	_ = http.ListenAndServe(":5050", api.GetMux())
}
