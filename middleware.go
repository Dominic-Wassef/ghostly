package ghostly

import "net/http"

func (c *Ghostly) SessionLoad(next http.Handler) http.Handler {
	return c.Session.LoadAndSave(next)
}
