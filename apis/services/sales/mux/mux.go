package mux

import (
	"net/http"
)

// WebAPIAuth constructs a http.Handler with all application routes bound.
func WebAPI() http.Handler {
	mux := http.NewServeMux()

	h := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	}

	mux.HandleFunc("GET /test", h)

	return mux
}
