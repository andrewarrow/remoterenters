package app

import "github.com/andrewarrow/feedback/router"

func HandleSites(c *router.Context, second, third string) {
	if second == "" {
		c.NotFound = true
	} else if second != "" && third == "" {
		handleSitesShow(c, second)
	} else {
		c.NotFound = true
	}
}

func handleSitesShow(c *router.Context, second string) {
	model := c.FindModel("story")
	params := []any{second}
	rows := c.SelectAllFrom(model, "where sub='' and domain=$1 order by created_at desc", params)
	c.SendContentInLayout("stories_index.html", rows, 200)
}
