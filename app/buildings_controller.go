package app

import "github.com/andrewarrow/feedback/router"

func HandleBuildings(c *router.Context, second, third string) {
	if second == "" {
		handleBuildingsIndex(c)
	} else if second != "" && third == "" {
		handleBuildingsShow(c)
	} else {
		c.NotFound = true
	}
}

func handleBuildingsIndex(c *router.Context) {
	rows := c.SelectAll("building", "", []any{})
	c.SendContentInLayout("buildings_index.html", rows, 200)
}

func handleBuildingsShow(c *router.Context) {
	c.SendContentInLayout("buildings_show.html", nil, 200)
}
