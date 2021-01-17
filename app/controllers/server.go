package controllers

import (
	"net/http"

	"../../config"
)

func StartMainServer() error {
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
