package ghostly

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (c *Ghostly) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	if c.Debug {
		mux.Use(middleware.Logger)
	}
	mux.Use(middleware.Recoverer)
	mux.Use(c.SessionLoad)

	return mux
}
