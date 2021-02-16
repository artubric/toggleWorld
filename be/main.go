package main

import (
	"net/http"

	featuretoggle "github.com/artubric/toggleWorld/featureToggle"
)

func main() {
	// serve all the "/api/*" paths
	featuretoggle.SetupRouters("/api")

	// serve static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.ListenAndServe(":5000", nil)
}
