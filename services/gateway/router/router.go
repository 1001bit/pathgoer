package router

import (
	"time"

	"github.com/1001bit/pathgoer/services/gateway/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func New() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(time.Second * 10))
	r.Use(middleware.RedirectSlashes)

	r.Get("/", routes.Home)

	return r
}
