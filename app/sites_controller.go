package app

import "github.com/andrewarrow/feedback/router"

func HandleSites(c *router.Context, second, third string) {
	if second == "" {
		c.NotFound = true
	} else if third != "" {
		c.NotFound = true
	} else {
		c.SendContentInLayout("welcome.html", WelcomeIndexVars(c.Db, "created_at desc", second), 200)
	}
}
