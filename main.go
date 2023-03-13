package main

import (
	"math/rand"
	"os"
	"remoterenters/app"
	"time"

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
		r := router.NewRouter()
		r.Paths["/"] = app.HandleWelcome
		r.Paths["buildings"] = app.HandleBuildings
		r.Paths["fresh"] = app.HandleFresh
		r.Paths["vote"] = app.HandleVote
		r.Paths["sites"] = app.HandleSites
		r.Paths["comments"] = app.HandleComments
		r.Paths["stories"] = app.HandleStories

		r.ListenAndServe(":3000")
	} else if arg == "help" {
	}
}
