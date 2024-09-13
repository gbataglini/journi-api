package middleware

import (
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Func func(HandlerFunc) HandlerFunc

func Cors(next HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, PUT, DELETE, GET, OPTIONS")
		w.Header().Set("Access-Control-Request-Method", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
		next(w, r)
	}
}
