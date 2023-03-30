package app

import (
	"github.com/andrewarrow/feedback/router"
)

func HandleAdminSubs(c *router.Context, second, third string) {
	if second == "" {
		handleAdminSubsIndex(c)
	} else if second != "" && third == "" {
		handleAdminSubsByUsername(c, second)
	} else {
		c.NotFound = true
	}
}

func handleAdminSubsIndex(c *router.Context) {
	rows := c.SelectAll("sub", "order by created_at desc", []any{}, "")
	c.SendContentInLayout("subs_index.html", rows, 200)
}

func handleAdminSubsByUsername(c *router.Context, username string) {
	u := c.LookupUsername(username)
	if len(u) == 0 {
		c.NotFound = true
		return
	}

	params := []any{u["id"]}
	rows := c.SelectAll("sub", "where user_id=$1 order by created_at desc",
		params, "")
	c.SendContentInLayout("subs_index.html", rows, 200)
}
