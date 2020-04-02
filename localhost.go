package middleware

import (
	"net"
	"net/http"
)

// Localhost middleware stops the http serve chain if caller is not calling from
// 127.0.0.1, localhost or ::1
// Returns http status 500 if SplitHostPort fails
// Return http status 404 if caller remoteAddr does not match
func Localhost(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		remoteAddr, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !(remoteAddr == "127.0.0.1" || remoteAddr == "localhost" || remoteAddr == "::1") {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		h.ServeHTTP(w, r)
	})
}

