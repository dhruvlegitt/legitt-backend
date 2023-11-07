package main

import (
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	// Setup log in log file

	// Register routes
	router := RegisterRoutes()

	print("Server running on port 8080")

	s := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	err = s.ListenAndServe()
	if err != nil {
		logger.Error().Msg(err.Error())
		os.Exit(1)
	}

}
