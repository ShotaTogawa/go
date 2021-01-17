package controllers

import (
	"net/http"

	"../../config"
)

func StartMainServer() error {
	// URLの登録
	http.HandleFunc("/", top)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
