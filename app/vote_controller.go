package app

import "github.com/andrewarrow/feedback/router"

func HandleVote(c *router.Context, second, third string) {
	if second == "" {
		c.NotFound = true
	} else if third != "" {
		c.NotFound = true
	} else {
		story := FetchStory(c.Db, second)
		if story == nil {
			c.Writer.WriteHeader(404)
			return
		}
		c.Db.Exec("update stories set points=points+1 where guid=$1", second)
		c.Writer.WriteHeader(200)
	}
}
