package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler //eta ekta type jeta onno ekta type ke support kore

type Manager struct {
	globalMiddlewares []Middleware //serve.go te manager.Use theke middleware gulo ekhane ashbe
}

func NewManager() *Manager {

	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

func (mngr *Manager) Use(middlewares ...Middleware) {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...)
}

func (mngr *Manager) With(handler http.Handler, middlewares ...Middleware) http.Handler {

	h := handler

	for _, middleware := range middlewares { //amra jodi extra middleware With() er moddhe dei routes er vitor theke tahole ekhane ashbe

		h = middleware(h)
	}

	return h
}

func (mngr *Manager) WrapMux(handler http.Handler) http.Handler {

	h := handler

	for _, middleware := range mngr.globalMiddlewares {

		h = middleware(h)
	}

	return h
}
