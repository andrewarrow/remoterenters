package main

import (
	"embed"
	_ "embed"
	"fmt"
	"math/rand"
	"os"
	"remoterenters/app"
	"time"

	"github.com/andrewarrow/feedback/router"
)

//go:embed app/feedback.json
var embeddedFile []byte

//go:embed views/*.html
var embeddedTemplates embed.FS

//go:embed assets/**/*.*
var embeddedAssets embed.FS

func main() {
	rand.Seed(time.Now().UnixNano())
	if len(os.Args) == 1 {
		return
	}
	arg := os.Args[1]

	if arg == "init" {
		router.InitNewApp()
	} else if arg == "run" {
		fmt.Println(len(embeddedFile))
		router.EmbeddedTemplates = embeddedTemplates
		router.EmbeddedAssets = embeddedAssets
		r := router.NewRouter("DATABASE_URL", embeddedFile)
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
	}
}
