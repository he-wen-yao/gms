package gms

import (
	"net/http"
)

// HandlerFunc defines the request handler used by gee
type HandlerFunc func(*context)

// Gms implement the interface of ServeHTTP
type gms struct {
	router *Router
}

// New is the constructor of gee.Gms
func New() *gms {
	return &gms{router: NewRouter()}
}

func (g *gms) addRoute(method string, pattern string, handler HandlerFunc) {
	g.router.AddRoute(method, pattern, handler)
}

// GET defines the method to add GET request
func (g *gms) GET(pattern string, handler HandlerFunc) {
	g.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (g *gms) POST(pattern string, handler HandlerFunc) {
	g.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http server
func (g *gms) Run(addr string) (err error) {
	return http.ListenAndServe(addr, g)
}

func (g *gms) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := NewContext(w, req)
	g.router.Handle(c)
}
