package main

import (
	"fmt"
	"math/rand"
	"os"
	"remoterenters/app"
	"time"

	"github.com/andrewarrow/feedback/persist"
	"github.com/andrewarrow/feedback/router"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	if len(os.Args) == 1 {
		return
	}
	arg := os.Args[1]

	if arg == "init" {
		router.InitNewApp()
	} else if arg == "run" {
		r := router.NewRouter("DATABASE_URL")
		r.Paths["/"] = app.HandleWelcome
		r.Paths["buildings"] = app.HandleBuildings
		r.Paths["fresh"] = app.HandleFresh
		r.Paths["vote"] = app.HandleVote
		r.Paths["sites"] = app.HandleSites
		r.Paths["comments"] = app.HandleComments
		r.Paths["stories"] = app.HandleStories
		r.Paths["cookies"] = app.HandleCookies
		r.Paths["rr"] = app.HandleSubs
		r.Paths["subs"] = app.HandleAdminSubs
		r.Paths["api"] = app.HandleApi
		r.BeforeCreate["story"] = app.PrepStory

		r.ListenAndServe(":3000")
	} else if arg == "export" {
		db := persist.PostgresConnection("DATABASE_URL")
		jsonString := persist.SchemaJson(db)
		fmt.Println(jsonString)
	}
}
