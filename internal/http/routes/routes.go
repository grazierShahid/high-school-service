package routes

import (
	"net/http"

	"github.com/grazierShahid/high-school-management-system/internal/http/handlers"
)

type Router struct {
	handler *handlers.UserHandler
	mux     *http.ServeMux
}

func NewRouter(handler *handlers.UserHandler, mux *http.ServeMux) *Router {
	return &Router{
		handler: handler,
		mux:     mux,
	}
}

func (router *Router) RegisterRoutes() {
	router.mux.HandleFunc("/health", router.handler.CheckHealth)
}
