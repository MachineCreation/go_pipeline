// server.go
//
// Author: Joseph Egan
// Created: 2025-12-30
// Description: server logic
//

package server

import (
	"log"
	"net/http"
)

func StartHTTPServer (address string, handler http.Handler) error {
	//print address to log
	log.Printf("Server listening on %s", address)
	// start service and return
	return http.ListenAndServe(address, handler)
}