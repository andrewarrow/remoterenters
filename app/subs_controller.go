package app

import "github.com/andrewarrow/feedback/router"

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
		c.SendContentInLayout("subs_show.html", nil, 200)
	}
}

func handleCreateSub(c *router.Context, slug string) {
}
