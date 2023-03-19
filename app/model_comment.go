package app

import (
	"github.com/andrewarrow/feedback/router"
)

func FetchComment(c *router.Context, guid string) map[string]any {
	params := []any{guid}
	return c.SelectOne("comment", "where guid=$1", params)
}
