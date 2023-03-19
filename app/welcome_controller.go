package app

import "github.com/andrewarrow/feedback/router"

func HandleWelcome(c *router.Context, second, third string) {
	if second == "" {
		handleWelcomeIndex(c)
	} else if second != "" && third == "" {
		c.NotFound = true
	} else {
		c.NotFound = true
	}
}

func handleWelcomeIndex(c *router.Context) {
	rows := c.SelectAll("story", "order by points desc", []any{})
	c.SendContentInLayout("stories_index.html", rows, 200)
}
