// health.go
//
// Author: Joseph Egan
// Created: 2025-12-30
// Description: health check handler
//

package handlers

import (
	"net/http"
)

//standard health check to determine if server is alive
func HealthHandler (writer http.ResponseWriter, response *http.Request) {
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("ok"))
}