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
	model := c.FindModel("story")
	rows := c.SelectAllFrom(model, "order by points desc", c.EmptyParams())
	c.SendContentInLayout("stories_index.html", rows, 200)
}
