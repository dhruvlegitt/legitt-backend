package main

import (
	"net/http"
)

// Register all routes here
func RegisterRoutes() http.Handler {
	// Main Root router for the app
	return c.Handler(router)
}
