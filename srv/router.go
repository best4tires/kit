package srv

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/best4tires/kit/log"
	"github.com/gorilla/mux"
)

type IgnoreTrailingSlashes struct{}

// Var extract a variable formerly registered by {} from the request
func Var(r *http.Request, key string) string {
	return mux.Vars(r)[key]
}

// Router encapsulates a http router
type Router struct {
	mux *mux.Router
}

// NewRouter creates a new router
func NewRouter() *Router {
	mr := mux.NewRouter()
	mr.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Warnf("not-found: %q", r.URL.String())
		http.Error(w, fmt.Sprintf("not-found %q", r.URL.Path), http.StatusNotFound)
	})
	r := &Router{
		mux: mr,
	}
	return r
}

// WithPrefix return a prefix-router for the given prefix
func (r *Router) WithPrefix(prefix string) *PrefixRouter {
	return &PrefixRouter{
		router: r,
		prefix: prefix,
	}
}

// Handler returns the http handler of the router
func (r *Router) Handler(mwares ...mux.MiddlewareFunc) http.Handler {
	r.mux.Use(mwares...)
	return r.mux
}

func (r *Router) containsOption(options []interface{}, opt interface{}) bool {
	for _, o := range options {
		if reflect.TypeOf(o) == reflect.TypeOf(opt) {
			return true
		}
	}
	return false
}

func (r *Router) patterns(s string, options ...interface{}) []string {
	var ps []string
	switch {
	case r.containsOption(options, IgnoreTrailingSlashes{}):
		p := strings.TrimSuffix(s, "/")
		ps = append(ps, p, p+"/")
	default:
		ps = append(ps, s)
	}
	return ps
}

// GET registers a GET handler
func (r *Router) GET(pattern string, handle http.HandlerFunc, options ...interface{}) {
	for _, p := range r.patterns(pattern, options...) {
		log.Infof("route GET %q", p)
		r.mux.Handle(p, handle).Methods("GET")
	}
}

// POST registers a POST handler
func (r *Router) POST(pattern string, handle http.HandlerFunc, options ...interface{}) {
	for _, p := range r.patterns(pattern, options...) {
		log.Infof("route POST %q", p)
		r.mux.Handle(p, handle).Methods("POST")
	}
}

func (r *Router) PUT(pattern string, handle http.HandlerFunc, options ...interface{}) {
	for _, p := range r.patterns(pattern, options...) {
		log.Infof("route PUT %q", p)
		r.mux.Handle(p, handle).Methods("PUT")
	}
}

func (r *Router) DELETE(pattern string, handle http.HandlerFunc, options ...interface{}) {
	for _, p := range r.patterns(pattern, options...) {
		log.Infof("route DELETE %q", p)
		r.mux.Handle(p, handle).Methods("DELETE")
	}
}

func (r *Router) HEAD(pattern string, handle http.HandlerFunc, options ...interface{}) {
	for _, p := range r.patterns(pattern, options...) {
		log.Infof("route HEAD %q", p)
		r.mux.Handle(p, handle).Methods("HEAD")
	}
}

// PrefixGET registers a GET handler, which matches all routes with the given prefix
func (r *Router) PrefixGET(prefix string, handle http.HandlerFunc, options ...interface{}) {
	r.mux.PathPrefix(prefix).HandlerFunc(handle).Methods("GET")
}

// PrefixPUT registers a PUT handler, which matches all routes with the given prefix
func (r *Router) PrefixPUT(prefix string, handle http.HandlerFunc, options ...interface{}) {
	r.mux.PathPrefix(prefix).HandlerFunc(handle).Methods("PUT")
}

// PrefixPOST registers a POST handler, which matches all routes with the given prefix
func (r *Router) PrefixPOST(prefix string, handle http.HandlerFunc, options ...interface{}) {
	r.mux.PathPrefix(prefix).HandlerFunc(handle).Methods("POST")
}

// PrefixDELETE registers a DELETE handler, which matches all routes with the given prefix
func (r *Router) PrefixDELETE(prefix string, handle http.HandlerFunc, options ...interface{}) {
	r.mux.PathPrefix(prefix).HandlerFunc(handle).Methods("DELETE")
}

// PrefixHEAD registers a HEAD handler, which matches all routes with the given prefix
func (r *Router) PrefixHEAD(prefix string, handle http.HandlerFunc, options ...interface{}) {
	r.mux.PathPrefix(prefix).HandlerFunc(handle).Methods("HEAD")
}

// PrefixGETHandler registers a GET handler, which matches all routes with the given prefix
func (r *Router) PrefixGETHandler(prefix string, handler http.Handler, options ...interface{}) {
	r.mux.PathPrefix(prefix).Handler(handler).Methods("GET")
}

// PrefixPOSTHandler registers a POST handler, which matches all routes with the given prefix
func (r *Router) PrefixPOSTHandler(prefix string, handler http.Handler, options ...interface{}) {
	r.mux.PathPrefix(prefix).Handler(handler).Methods("POST")
}

// PrefixRouter is a router, which routes prefixed routes
type PrefixRouter struct {
	router *Router
	prefix string
}

func (r *PrefixRouter) Resolve(pattern string) string {
	return r.prefix + pattern
}

// GET registers a GET handler
func (r *PrefixRouter) GET(pattern string, handle http.HandlerFunc, options ...interface{}) {
	r.router.GET(r.prefix+pattern, handle, options...)
}

// POST registers a POST handler
func (r *PrefixRouter) POST(pattern string, handle http.HandlerFunc, options ...interface{}) {
	r.router.POST(r.prefix+pattern, handle, options...)
}

// PrefixGET registers a GET handler, which matches all routes with the given prefix
func (r *PrefixRouter) PrefixGET(prefix string, handle http.HandlerFunc, options ...interface{}) {
	r.router.PrefixGET(r.prefix+prefix, handle, options...)
}

func (r *PrefixRouter) WithPrefix(prefix string) *PrefixRouter {
	return &PrefixRouter{
		router: r.router,
		prefix: r.prefix + prefix,
	}
}

func (r *PrefixRouter) Prefix() string {
	return r.prefix
}

func (r *PrefixRouter) Handler(mwares ...mux.MiddlewareFunc) http.Handler {
	r.router.mux.Use(mwares...)
	return r.router.mux
}
