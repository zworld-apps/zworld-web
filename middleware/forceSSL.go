package middleware

import (
	"net/http"
	"os"
)

const (
	xForwardedProtoHeader = "x-forwarded-proto"
)

// ForceSSL is a middleware that forces HTTPS requests to the Heroku load
// balancer. It ensures some headers present on Heroku server.
//
// Copied from https://github.com/jonahgeorge/force-ssl-heroku
func ForceSSL(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if os.Getenv("GO_ENV") == "production" {
			if r.Header.Get(xForwardedProtoHeader) != "https" {
				sslUrl := "https://" + r.Host + r.RequestURI
				http.Redirect(w, r, sslUrl, http.StatusTemporaryRedirect)
				return
			}
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
