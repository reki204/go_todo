package controllers

import (
	"net/http"

	"github.com/reki204/go_todo/config"
)

func StartMainServer() error {
	http.HandleFunc("/", top)
	return http.ListenAndServe(":" + config.Config.Port, nil)
}
