package app

import (
	"html"
	"html/template"
	"net/http"
	"strings"

	"github.com/andrewarrow/feedback/models"
	"github.com/andrewarrow/feedback/router"
	"github.com/andrewarrow/feedback/util"
)

func HandleStories(c *router.Context, second, third string) {
	if second == "" {
		handleStoriesIndex(c)
	} else if second != "" && third == "" {
		handleStoryShow(c, second)
	} else {
		c.NotFound = true
	}
}

func handleStoriesIndex(c *router.Context) {
	if c.Method == "POST" {
		title := strings.TrimSpace(c.Request.FormValue("title"))
		url := strings.TrimSpace(c.Request.FormValue("url"))
		body := strings.TrimSpace(c.Request.FormValue("body"))
		returnPath := "/stories/new/"
		if len(title) < 10 {
			router.SetFlash(c, "title too short.")
			http.Redirect(c.Writer, c.Request, returnPath, 302)
			return
		}
		if len(title) > 140 {
			router.SetFlash(c, "title too long.")
			http.Redirect(c.Writer, c.Request, returnPath, 302)
			return
		}
		if body == "" && url == "" {
			router.SetFlash(c, "body or url required.")
			http.Redirect(c.Writer, c.Request, returnPath, 302)
			return
		}
		if url != "" {
			if len(url) < 13 ||
				!(strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")) {
				router.SetFlash(c, "url too short.")
				http.Redirect(c.Writer, c.Request, returnPath, 302)
				return
			}
		}
		if body != "" && len(body) < 10 {
			router.SetFlash(c, "body too short.")
			http.Redirect(c.Writer, c.Request, returnPath, 302)
			return
		}
		guid := util.PseudoUuid()
		sub := router.GetCookie(c, "sub")
		if url != "" {
			domain := util.ExtractDomain(url)
			c.Db.Exec("insert into stories (points, sub, title, url, guid, username, domain) values ($1, $2, $3, $4, $5, $6, $7)", 1, sub, title, url, guid, c.User["username"], domain)
		} else {
			c.Db.Exec("insert into stories (points, sub, title, body, guid, username) values ($1, $2, $3, $4, $5, $6)", 1, sub, title, body, guid, c.User["username"])
		}
		if sub == "" {
			http.Redirect(c.Writer, c.Request, "/", 302)
		} else {
			http.Redirect(c.Writer, c.Request, "/rr/"+sub+"/", 302)
		}
		return
	}
	c.NotFound = true
}

type StoryShow struct {
	Story    map[string]any
	Comments []map[string]any
}

func handleStoryShow(c *router.Context, second string) {
	if second == "new" {
		if len(c.User) == 0 {
			c.UserRequired = true
			return
		}
		sub := router.GetCookie(c, "sub")
		c.SendContentInLayout("stories_new.html", sub, 200)
	} else if second != "" {

		story := FetchStory(c, second)

		if len(story) == 0 {
			c.NotFound = true
			return
		}
		story["title"] = models.RemoveMostNonAlphanumeric(story["title"].(string))
		body := strings.Replace(html.EscapeString(story["body"].(string)), "\n", "<br/>", -1)
		story["body"] = template.HTML(body + "<br/><br/>")
		params := []any{story["id"]}
		storyShow := StoryShow{}
		storyShow.Story = story
		storyShow.Comments = c.SelectAll("comment", "where story_id=$1", params)
		c.Title = storyShow.Story["title"].(string)
		c.SendContentInLayout("stories_show.html", storyShow, 200)
	}
}
