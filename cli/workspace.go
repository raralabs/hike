package cli

import (
	"log"
	"strings"

	"github.com/raralabs/go-wm/wm"
)

func WSMode(f func(string) string, project func(*wm.Project, string), mode string) {
	ws := wm.OpenWorkSpace("hike")

	fn := func(p *wm.Project) {
		project(p, mode)
	}

	option := strings.TrimSpace(f("> "))
	switch option {
	case "create":
		projName := f("Project Name> ")
		projDir := f("Project Directory> ")
		pr, _ := ws.AddNewProject(projName, projDir)
		pr.Run(fn)
		pr.Save()

	case "open":
		projName := f("Project Name> ")
		pr, err := ws.OpenProject(projName)
		if err != nil {
			log.Panic(err)
		}
		pr.Run(fn)
		pr.Save()

	case "remove":
		log.Panicln("No support")
	}
}
