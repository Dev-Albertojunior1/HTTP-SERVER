package router

import (
	"context"
	"net/http"
	"strings"

	"honnef.co/go/tools/pattern"
)

type Router struct {
	routes map[string]http.HandlerFunc
}

func NewRouter() *Router {

	return &Router{
		routes: make(map[string]http.HandlerFunc),
	}
}


func (r *Router) HandleFunc(method, pattern string, handler http.HandlerFunc) {
	if r.routes[method] == nil {
		r.routes[method] = make(map[string]http.HandlerFunc)
	}
	r.routes[method][pattern] = handler
}

func (r *Router) GET(pattern string, handler http.HandlerFunc){
  r.HandleFunc(http.MethodGet, pattern, handler)
}


func (r *Router) POST(pattern string, handler http.HandlerFunc){
  r.HandleFunc((http.MethodPost, pattern, handler)
}

func (r *Router) DELETE(pattern string, handler http.HandlerFunc){

}

