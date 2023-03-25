package app

import (
	"github.com/andrewarrow/feedback/router"
)

func HandleApi(c *router.Context, second, third string) {
	if second == "user" && third == "" && c.Method == "POST" {
		handleApiCreateUser(c)
		return
	}
	c.NotFound = true
}

func handleApiCreateUser(c *router.Context) {
	params := c.ReadBodyIntoJson()
	model := c.FindModel("user")
	for _, field := range model.RequiredFields() {
		if params[field.Name] == nil {
			c.SendContentAsJson(c.JsonInfo("missing "+field.Name), 422)
			return
		}
	}

	c.SendContentAsJson(c.JsonInfo("ok"), 200)
}
