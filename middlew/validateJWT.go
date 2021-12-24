package middlew

import (
	"net/http"
	"twister/app/routers"
)

// Check if we have a valid jwt
func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcessToken(r.Header.Get("Authorization"))

		if err != nil {
			http.Error(rw, "Error on token"+err.Error(), http.StatusBadRequest)
		}

		next.ServeHTTP(rw, r)
	}
}
