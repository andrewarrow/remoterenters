package app

import (
	"net/http"

	"github.com/andrewarrow/feedback/router"
)

func HandleCookies(c *router.Context, second, third string) {
	if second == "" {
		c.NotFound = true
	} else if second != "" && third == "" {
		handleRemoveSub(c)
	} else {
		c.NotFound = true
	}
}

func handleRemoveSub(c *router.Context) {
	router.SetCookie(c, "sub", "")
	http.Redirect(c.Writer, c.Request, "/stories/new/", 302)
}
