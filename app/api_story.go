package app

import "github.com/andrewarrow/feedback/router"

func handleApiCreateStory(c *router.Context) {
	message := c.ValidateJsonForModel("story")
	if message != "" {
		c.SendContentAsJsonMessage(message, 422)
		return
	}

	message = c.CreateRowFromJson("story")
	if message != "" {
		c.SendContentAsJsonMessage(message, 422)
		return
	}
	guid := c.Params["guid"]
	row := c.SelectOne("story", "where guid=$1", []any{guid})
	c.SendContentAsJson(row, 200)
}
