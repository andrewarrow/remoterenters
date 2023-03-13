package app

import (
	"net/http"
	"strings"

	"github.com/andrewarrow/feedback/router"
	"github.com/andrewarrow/feedback/util"
)

type StoryShow struct {
	Story    *Story
	Comments []*Comment
}

func HandleStories(c *router.Context, second, third string) {
	if second == "" {
		handleStoriesIndex(c)
	} else if third != "" {
		c.NotFound = true
	} else {
		if second == "new" {
			c.UserRequired = true
			if c.User != nil {
				c.SendContentInLayout("stories_new.html", nil, 200)
			}
			return
		} else if second != "" {
			story := FetchStory(c.Db, second)
			if story == nil {
				c.NotFound = true
				return
			}
			if story.HasUrl && story.Domain == "" {
				story.Domain = util.ExtractDomain(story.Url)
				c.Db.Exec("update stories set domain=$1 where guid=$2", story.Domain, second)
			}
			storyShow := StoryShow{}
			storyShow.Story = story
			storyShow.Comments = FetchComments(c.Db, story.Id)
			c.Title = story.Title
			c.SendContentInLayout("stories_show.html", storyShow, 200)
			return
		}
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
