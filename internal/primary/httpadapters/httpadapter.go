package httpadapters

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type HTTPAdapter struct {
	router      chi.Router
	Middlewares []func(http.Handler) http.Handler
}

type Route struct {
	Pattern     string
	Endpoints   []Endpoint
	SubRoutes   []Route
	Middlewares []func(http.Handler) http.Handler
}

type Endpoint struct {
	Method      string
	Pattern     string
	Function    http.HandlerFunc
	Middlewares []func(http.Handler) http.Handler
}

// Returns an adapter for run http server
// routes []Route routes from application
// timeout int duration in seconds
func NewHTTPAdapter(routes []Route, timeout int) *HTTPAdapter {
	h := &HTTPAdapter{
		chi.NewRouter(),
		[]func(http.Handler) http.Handler{
			middleware.RequestID,
			middleware.RealIP,
			middleware.Logger,
			middleware.Recoverer,
			middleware.Timeout(time.Duration(timeout) * time.Second),
		},
	}
	h.setMiddlewares(routes)
	h.setRoutes(routes)
	return h
}

// Serves http application
func (h *HTTPAdapter) Serve(port string) {
	err := http.ListenAndServe(":"+port, h.router)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (h *HTTPAdapter) setMiddlewares(routes []Route) {
	for _, route := range routes {
		h.Middlewares = append(h.Middlewares, route.Middlewares...)
	}
	h.router.Use(h.Middlewares...)
}

func (h *HTTPAdapter) setRoutes(routes []Route) {
	for _, route := range routes {
		h.router.Route(route.Pattern, func(r chi.Router) {
			setEndpoints(r, route.Endpoints)
			h.setRoutes(route.SubRoutes)
		})
		h.Middlewares = append(h.Middlewares, route.Middlewares...)
	}
}

func setEndpoints(r chi.Router, endpoints []Endpoint) {
	for _, endpoint := range endpoints {
		switch endpoint.Method {
		case http.MethodGet:
			r.Get(endpoint.Pattern, endpoint.Function)
		case http.MethodPost:
			r.Post(endpoint.Pattern, endpoint.Function)
		case http.MethodPut:
			r.Put(endpoint.Pattern, endpoint.Function)
		case http.MethodDelete:
			r.Delete(endpoint.Pattern, endpoint.Function)
		}
	}
}
