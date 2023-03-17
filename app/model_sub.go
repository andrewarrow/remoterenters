package app

import "github.com/andrewarrow/feedback/router"

func FetchSub(c *router.Context, slug string) map[string]any {
	model := c.FindModel("sub")
	params := []any{slug}
	return c.SelectOneFrom(model, "where slug=$1", params)
}
