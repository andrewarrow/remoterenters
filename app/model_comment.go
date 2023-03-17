package app

import (
	"github.com/andrewarrow/feedback/router"
)

func FetchComment(c *router.Context, guid string) map[string]any {
	model := c.FindModel("comment")
	params := []any{guid}
	return c.SelectOneFrom(model, "where guid=$1", params)
}
