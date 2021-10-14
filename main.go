package main

import (
	"httpserver/logger"
	"httpserver/server"
	"net/http"
)

func main() {
	logger.Debugf("hhhh")
	http.HandleFunc("/healthz", server.Healthz)
	http.ListenAndServe("0.0.0.0:80", nil)
}
