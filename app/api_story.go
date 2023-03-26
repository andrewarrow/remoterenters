package app

import (
	"html"
	"html/template"
	"strings"

	"github.com/andrewarrow/feedback/models"
	"github.com/andrewarrow/feedback/router"
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

	c.Params["title"] = models.RemoveMostNonAlphanumeric(c.Params["title"].(string))
	body := strings.Replace(html.EscapeString(c.Params["body"].(string)), "\n", "<br/>", -1)
	c.Params["body"] = template.HTML(body + "<br/><br/>")
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
