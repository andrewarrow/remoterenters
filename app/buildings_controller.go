package app

import "github.com/andrewarrow/feedback/router"

func HandleBuildings(c *router.Context, second, third string) {
	if second == "" {
		handleBuildingsIndex(c)
	} else if third != "" {
		c.NotFound = true
	} else {
		handleBuildingsShow(c)
	}
}

type BuildingVars struct {
	Rows []*Building
}

func handleBuildingsIndex(c *router.Context) {
	vars := BuildingVars{}
	vars.Rows = FetchBuildings(c.Db)
	c.SendContentInLayout("buildings_index.html", &vars, 200)
}

func handleBuildingsShow(c *router.Context) {
	c.SendContentInLayout("buildings_show.html", nil, 200)
}
