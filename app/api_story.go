package app

import (
	"github.com/andrewarrow/feedback/router"
)

func handleApiCreateStory(c *router.Context) {
	if c.User == nil {
		c.SendContentAsJsonMessage("Authorization bad", 401)
		return
	}
	c.ReadJsonBodyIntoParams()
	message := c.Validate("story")
	if message != "" {
		c.SendContentAsJsonMessage(message, 422)
		return
	}

	message = c.Insert("story")
	if message != "" {
		c.SendContentAsJsonMessage(message, 422)
		return
	}

	guid := c.Params["guid"]
	row := c.SelectOne("story", "where guid=$1", []any{guid})
	c.SendContentAsJson(row, 200)
}
