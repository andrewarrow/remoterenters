package app

import (
	"github.com/andrewarrow/feedback/router"
	"github.com/andrewarrow/feedback/util"
)

func FetchStory(c *router.Context, guid string) map[string]any {
	params := []any{guid}
	return c.SelectOne("story", "where guid=$1", params)
}

func PrepStory(c *router.Context) {
	if c.Params["url"] != nil {
		c.Params["domain"] = util.ExtractDomain(c.Params["url"].(string))
	}
	c.Params["username"] = c.User["username"].(string)
	c.Params["points"] = 1
}
