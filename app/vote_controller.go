package app

import "github.com/andrewarrow/feedback/router"

func HandleVote(c *router.Context, second, third string) {
	if second == "" {
		c.NotFound = true
	} else if second != "" && third == "" {
		story := FetchStory(c, second)
		if len(story) == 0 {
			c.Writer.WriteHeader(404)
			return
		}
		c.Db.Exec("update stories set points=points+1 where guid=$1", second)
		c.Writer.WriteHeader(200)
	} else {
		c.NotFound = true
	}
}
