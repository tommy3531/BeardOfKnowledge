package routes

import (
	"net/http"
)

type Routes []Route

type Route struct {
	Name 		string
	Method 		string
	Pattern 	string
	HandleFunc 	http.HandlerFunc
}

type SubRoutePackage struct {
	Routes Routes
	Middelware func (next http.Handler) http.Handler
}