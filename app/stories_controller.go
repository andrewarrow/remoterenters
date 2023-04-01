package app

import (
	"html"
	"html/template"
	"net/http"
	"strings"

	"github.com/andrewarrow/feedback/models"
	"github.com/andrewarrow/feedback/router"
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
		sub := router.GetCookie(c, "sub")

		c.Params = map[string]any{"title": title}
		if body != "" {
			c.Params["body"] = body
		}
		if url != "" {
			c.Params["url"] = url
		}
		if sub != "" {
			c.Params["sub"] = sub
		}

		returnPath := "/stories/new"
		message := c.ValidateCreate("story")
		if message != "" {
			router.SetFlash(c, message)
			http.Redirect(c.Writer, c.Request, returnPath, 302)
			return
		}

		message = c.Insert("story")
		if message != "" {
			router.SetFlash(c, message)
			http.Redirect(c.Writer, c.Request, returnPath, 302)
			return
		}

		if sub == "" {
			http.Redirect(c.Writer, c.Request, "/", 302)
		} else {
			http.Redirect(c.Writer, c.Request, "/rr/"+sub, 302)
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
		storyShow.Comments = c.SelectAll("comment", "where story_id=$1", params, "")
		c.Title = storyShow.Story["title"].(string)
		c.SendContentInLayout("stories_show.html", storyShow, 200)
	}
}
