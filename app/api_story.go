package app

import (
	"github.com/andrewarrow/feedback/router"
	"github.com/andrewarrow/feedback/util"
)

func handleApiCreateStory(c *router.Context) {
	if c.User == nil {
		c.SendContentAsJsonMessage("Authorization bad", 401)
		return
	}
	message := c.ValidateJsonForModel("story")
	if message != "" {
		c.SendContentAsJsonMessage(message, 422)
		return
	}

	if c.Params["url"] != nil {
		c.Params["domain"] = util.ExtractDomain(c.Params["url"].(string))
	}
	c.Params["username"] = c.User["username"].(string)
	message = c.CreateRowFromJson("story")
	if message != "" {
		c.SendContentAsJsonMessage(message, 422)
		return
	}

	guid := c.Params["guid"]
	row := c.SelectOne("story", "where guid=$1", []any{guid})
	c.SendContentAsJson(row, 200)
}
