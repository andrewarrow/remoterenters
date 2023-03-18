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
	model := c.FindModel("sub")
	rows := c.SelectAllFrom(model, "order by created_at desc", c.EmptyParams())
	c.SendContentInLayout("subs_index.html", rows, 200)
}

func handleAdminSubsByUsername(c *router.Context, username string) {
	u := c.LookupUsername(username)
	if u == nil {
		c.NotFound = true
		return
	}

	model := c.FindModel("sub")
	params := []any{u.Id}
	rows := c.SelectAllFrom(model, "where user_id=$1 order by created_at desc",
		params)
	c.SendContentInLayout("subs_index.html", rows, 200)
}
