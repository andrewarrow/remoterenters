package app

import (
	"github.com/andrewarrow/feedback/router"
)

func FetchStory(c *router.Context, guid string) map[string]any {
	params := []any{guid}
	return c.SelectOne("story", "where guid=$1", params)
}
