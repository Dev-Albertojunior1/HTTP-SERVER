package router

import (
	"context"
	"net/http"
	"strings"
)

type Router struct {
	routes map[string]http.HandlerFunc
}
