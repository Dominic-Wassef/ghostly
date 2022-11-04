package ghostly

import (
	"net/http"
	"strconv"

	"github.com/justinas/nosurf"
)

func (g *Ghostly) SessionLoad(next http.Handler) http.Handler {
	return g.Session.LoadAndSave(next)
}

func (g *Ghostly) NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	secure, _ := strconv.ParseBool(g.config.cookie.secure)

	csrfHandler.ExemptGlob("/api/*")

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   secure,
		SameSite: http.SameSiteStrictMode,
		Domain:   g.config.cookie.domain,
	})

	return csrfHandler
}
