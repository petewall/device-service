package main

import (
	"fmt"
	"net/http"

	. "github.com/petewall/device-service/v2"
)

func main() {
	config := &Config{
		DBHost: "localhost",
		DBPort: 6379,
		Port:   5050,
	}

	db := Connect(config)
	api := &API{
		DB: db,
	}

	_ = http.ListenAndServe(fmt.Sprintf(":%d", config.Port), api.GetMux())
}
