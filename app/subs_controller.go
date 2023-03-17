package app

import (
	"net/http"

	"github.com/andrewarrow/feedback/router"
	"github.com/andrewarrow/feedback/util"
)

func HandleSubs(c *router.Context, second, third string) {
	if second == "" {
		handleSubsIndex(c)
	} else if second != "" && third == "" {
		handleSubsShow(c, second)
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
		model := c.FindModel("story")
		params := []any{slug}
		rows := c.SelectAllFrom(model, "where sub=$1 order by points desc", params)
		c.SendContentInLayout("stories_index.html", rows, 200)
	}
}

func handleCreateSub(c *router.Context, slug string) {
	if c.User == nil {
		c.UserRequired = true
		return
	}
	guid := util.PseudoUuid()
	c.Db.Exec("insert into subs (guid, slug, user_id) values ($1, $2, $3)", guid, slug, c.User.Id)
	http.Redirect(c.Writer, c.Request, "/rr/"+slug+"/", 302)
}
