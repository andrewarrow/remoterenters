package app

import (
	"github.com/andrewarrow/feedback/router"
)

func FetchStory(c *router.Context, guid string) map[string]any {
	model := c.FindModel("story")
	params := []any{guid}
	return c.SelectOneFrom(model, "where guid=$1", params)
}
