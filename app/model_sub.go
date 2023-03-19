package app

import "github.com/andrewarrow/feedback/router"

func FetchSub(c *router.Context, slug string) map[string]any {
	params := []any{slug}
	return c.SelectOne("sub", "where slug=$1", params)
}
