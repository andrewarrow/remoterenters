package app

import (
	"net/http"
	"strings"

	"github.com/andrewarrow/feedback/router"
	"github.com/andrewarrow/feedback/util"
)

func HandleComments(c *router.Context, second, third string) {
	if second == "" {
		c.NotFound = true
	} else if second != "" && third == "" {
		if c.Method == "POST" {
			postComment(c, second)
		} else {
			showComment(c, second)
		}
	} else {
		c.NotFound = true
	}
}

func showComment(c *router.Context, second string) {
	comment := FetchComment(c, second)
	if len(comment) == 0 {
		c.NotFound = true
		return
	}
	c.Title = comment["body"].(string)
	if len(c.Title) > 80 {
		c.Title = c.Title[0:80] + "..."
	}
	story := FetchStory(c, comment["story_guid"].(string))
	if len(story) == 0 {
		c.NotFound = true
		return
	}
	title := story["title"].(string)
	if len(title) > 40 {
		title = title[0:40] + "..."
	}
	comment["story_title"] = title
	c.SendContentInLayout("comments_show.html", comment, 200)
	return
}

func postComment(c *router.Context, second string) {
	if c.User == nil {
		c.UserRequired = true
		return
	}
	body := strings.TrimSpace(c.Request.FormValue("body"))
	returnPath := "/stories/" + second + "/"
	if len(body) < 10 {
		router.SetFlash(c, "body too short.")
		http.Redirect(c.Writer, c.Request, returnPath, 302)
		return
	}

	guid := util.PseudoUuid()
	story := FetchStory(c, second)
	if len(story) == 0 {
		c.NotFound = true
		return
	}

	tx := c.Db.MustBegin()
	tx.Exec("insert into comments (body, guid, username, story_id, story_guid) values ($1, $2, $3, $4, $5)", body, guid, c.User.Username, story["id"], story["guid"])
	tx.Exec("update stories set comments=comments+1 where id=$1", story["id"])
	tx.Commit()
	http.Redirect(c.Writer, c.Request, returnPath, 302)
}
