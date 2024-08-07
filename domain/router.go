package domain

import "net/http"

type Router interface {
	Routes(*http.ServeMux)
}
