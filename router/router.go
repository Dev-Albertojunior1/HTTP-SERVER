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

func (r *Router) GET(pattern string, handler http.HandlerFunc) {
	r.HandleFunc(http.MethodGet, pattern, handler)
}

func (r *Router) POST(pattern string, handler http.HandlerFunc) {
	r.HandleFunc(http.MethodPost, pattern, handler)
}

func (r *Router) DELETE(pattern string, handler http.HandlerFunc) {
	r.HandleFunc(http.MethodDelete, pattern, handler)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if handler, params, found := r.findHandler(req.Method, req.URL.Path); found {
		ctx := req.Context()
		for k, v := range params {
			ctx = context.WithValue(ctx, k, v)
		}

		handler.ServeHTTP(w, req.WithContext(ctx))
		return
	}
	http.NotFound(w, req)
}

func (r *Router) findHandler(method, path string) (http.HandlerFunc, map[string]string, bool) {

	for pattern, handler := range r.routes[method] {
		params, found := match(pattern, path)
		if found {
			return handler, params, true
		}
	}
	return nil, nil, false
}

func match(pattern) {

}
