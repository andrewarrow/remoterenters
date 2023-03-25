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
	message := c.ValidateJsonForModel("user")
	if message != "" {
		c.SendContentAsJson(c.JsonInfo(message), 422)
		return
	}

	c.SendContentAsJson(c.JsonInfo("ok"), 200)

}
