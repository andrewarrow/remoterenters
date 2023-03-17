package app

import (
	"net/http"
	"strings"

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
		if url != "" {
			domain := util.ExtractDomain(url)
			c.Db.Exec("insert into stories (title, url, guid, username, domain) values ($1, $2, $3, $4, $5)", title, url, guid, c.User.Username, domain)
		} else {
			c.Db.Exec("insert into stories (title, body, guid, username) values ($1, $2, $3, $4)", title, body, guid, c.User.Username)
		}
		http.Redirect(c.Writer, c.Request, "/", 302)
		return
	}
	c.NotFound = true
}

type StoryShow struct {
	Story    map[string]any
	Comments []*map[string]any
}

func handleStoryShow(c *router.Context, second string) {
	if second == "new" {
		if c.User == nil {
			c.UserRequired = true
			return
		}
		c.SendContentInLayout("stories_new.html", nil, 200)
	} else if second != "" {

		storyShow := StoryShow{}

		storyShow.Story = FetchStory(c, second)

		if len(storyShow.Story) == 0 {
			c.NotFound = true
			return
		}
		model := c.FindModel("comment")
		params := []any{storyShow.Story["id"]}
		storyShow.Comments = c.SelectAllFrom(model, "where story_id=$1", params)
		c.Title = storyShow.Story["title"].(string)
		c.SendContentInLayout("stories_show.html", storyShow, 200)
	}
}