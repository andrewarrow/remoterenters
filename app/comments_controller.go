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
	} else if third != "" {
		c.NotFound = true
	} else {
		if c.Method == "POST" {
			postComment(c, second)
		} else {
			showComment(c, second)
		}
	}
}

func showComment(c *router.Context, second string) {
	comment := FetchComment(c.Db, second)
	if comment == nil {
		c.NotFound = true
		return
	}
	c.Title = comment.RawBody
	if len(c.Title) > 80 {
		c.Title = c.Title[0:80] + "..."
	}
	story := FetchStory(c.Db, comment.StoryGuid)
	if story == nil {
		c.NotFound = true
		return
	}
	comment.StoryTitle = story.Title
	if len(comment.StoryTitle) > 40 {
		comment.StoryTitle = comment.StoryTitle[0:40] + "..."
	}
	c.SendContentInLayout("comments_show.html", comment, 200)
	return
}

func postComment(c *router.Context, second string) {
	c.UserRequired = true
	if c.User == nil {
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
	story := FetchStory(c.Db, second)
	if story == nil {
		c.NotFound = true
		return
	}

	tx := c.Db.MustBegin()
	tx.Exec("insert into comments (body, guid, username, story_id, story_guid) values ($1, $2, $3, $4, $5)", body, guid, c.User.Username, story.Id, story.Guid)
	tx.Exec("update stories set comments=comments+1 where id=$1", story.Id)
	tx.Commit()
	http.Redirect(c.Writer, c.Request, returnPath, 302)
}
