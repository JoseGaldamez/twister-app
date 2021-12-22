package middlew

import (
	"net/http"
	"twister/app/db"
)

// CheckDataBase is a middleware that let know if DB is Connected
func CheckDataBase(next http.HandlerFunc) http.HandlerFunc {

	return func(rw http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(rw, "Connection lost", 500)
			return
		}

		next.ServeHTTP(rw, r)
	}

}
