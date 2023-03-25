package app

import "github.com/andrewarrow/feedback/router"

func HandleApi(c *router.Context, second, third string) {
	if second != "" && third == "" {
		handleApiCall(c)
		return
	}
	c.NotFound = true
}

func handleApiCall(c *router.Context) {
	m := map[string]any{}
	m["test"] = []string{"hi", "there"}
	c.SendContentAsJson(m, 200)
}
