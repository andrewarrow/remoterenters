package app

import (
	"github.com/andrewarrow/feedback/router"
	"github.com/jmoiron/sqlx"
)

func HandleWelcome(c *router.Context, second, third string) {
	c.SendContentInLayout("welcome.html", WelcomeIndexVars(c.Db, "points desc", ""), 200)
}

type WelcomeVars struct {
	Rows []*Story
}

func WelcomeIndexVars(db *sqlx.DB, order, domain string) *WelcomeVars {
	vars := WelcomeVars{}
	vars.Rows = FetchStories(db, order, domain)
	return &vars
}
