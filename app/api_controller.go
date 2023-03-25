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
		c.SendContentAsJsonMessage(message, 422)
		return
	}

	message = c.CreateRowFromJson("user")
	if message != "" {
		c.SendContentAsJsonMessage(message, 422)
		return
	}
	guid := c.Params["guid"]
	row := c.SelectOne("user", "where guid=$1", []any{guid})
	c.SendContentAsJson(row, 200)
}
