package app

import "github.com/andrewarrow/feedback/router"

func HandleFresh(c *router.Context, second, third string) {
	if second == "" {
		handleFreshIndex(c)
	} else if second != "" && third == "" {
		c.NotFound = true
	} else {
		c.NotFound = true
	}
}

func handleFreshIndex(c *router.Context) {
	rows := c.SelectAll("story", "order by created_at desc", []any{})
	c.SendContentInLayout("stories_index.html", rows, 200)
}
