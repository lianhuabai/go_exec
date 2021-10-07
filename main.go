package main

import (
	"homework/logger"
	"homework/server"
	"net/http"
)

func main() {
	logger.Debugf("hhhh")
	http.HandleFunc("/healthz", server.Healthz)
	http.ListenAndServe("127.0.0.1:80", nil)
}
