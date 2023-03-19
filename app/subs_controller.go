package app

import (
	"net/http"

	"github.com/andrewarrow/feedback/models"
	"github.com/andrewarrow/feedback/router"
	"github.com/andrewarrow/feedback/util"
)

func HandleSubs(c *router.Context, second, third string) {
	if second == "" {
		handleSubsIndex(c)
	} else if second != "" && third == "" {
		handleSubsShow(c, models.MakeSlug(second))
	} else {
		c.NotFound = true
	}
}

func handleSubsIndex(c *router.Context) {
	c.SendContentInLayout("subs_index.html", nil, 200)
}

func handleSubsShow(c *router.Context, slug string) {
	if c.Method == "POST" {
		handleCreateSub(c, slug)
		return
	}
	sub := FetchSub(c, slug)
	if len(sub) == 0 {
		c.SendContentInLayout("subs_new.html", slug, 200)
	} else {
		router.SetCookie(c, "sub", slug)
		params := []any{slug}
		rows := c.SelectAll("story", "where sub=$1 order by points desc", params)
		c.SendContentInLayout("stories_index.html", rows, 200)
	}
}

func handleCreateSub(c *router.Context, slug string) {
	if len(c.User) == 0 {
		c.UserRequired = true
		return
	}
	guid := util.PseudoUuid()
	c.Db.Exec("insert into subs (username, guid, slug, user_id) values ($1, $2, $3, $4)",
		c.User["username"], guid, slug, c.User["id"])
	http.Redirect(c.Writer, c.Request, "/rr/"+slug+"/", 302)
}
