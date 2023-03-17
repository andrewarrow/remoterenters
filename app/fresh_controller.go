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
	model := c.FindModel("story")
	rows := c.SelectAllFrom(model, "order by created_at desc", c.EmptyParams())
	c.SendContentInLayout("stories_index.html", rows, 200)
}
