package cli

import (
	"log"
	"strings"

	"github.com/raralabs/go-wm/wm"
)

func WSMode(f func(string) string, project func(*wm.Project)) {
	ws := wm.OpenWorkSpace("hike")

	mode := strings.TrimSpace(f("> "))
	switch mode {
	case "create":
		projName := f("Project Name> ")
		projDir := f("Project Directory> ")
		pr, _ := ws.AddNewProject(projName, projDir)
		pr.Run(project)
		pr.Save()

	case "open":
		projName := f("Project Name> ")
		pr, err := ws.OpenProject(projName)
		if err != nil {
			log.Panic(err)
		}
		pr.Run(project)
		pr.Save()

	case "remove":
		log.Panicln("No support")
	}
}
