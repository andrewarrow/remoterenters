package app

import "github.com/andrewarrow/feedback/router"

func HandleFresh(c *router.Context, second, third string) {
	if second == "" {
		handleFreshIndex(c)
	} else if third != "" {
		c.NotFound = true
	} else {
		c.NotFound = true
	}
}
func handleFreshIndex(c *router.Context) {
	c.SendContentInLayout("welcome.html", WelcomeIndexVars(c.Db, "created_at desc", ""), 200)
}
