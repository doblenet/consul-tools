package main

import (
	"github.com/doblenet/go-doblenet/cliapp"
	cmd "./commands"
)

const (
	gitCommit = "%GIT_COMMIT%"
)

var	appInfo = cliapp.AppInfo{Name: "consul-sc",
		Version: "0.4.1",
		GitTag: "",
		Copyright: "(C)2016 DOBLENET Soluciones Tecnológicas, S.L.",
	}
	

func main() {

	app := cliapp.New(appInfo)
	app.SetCommander(cmd.RootCmd)
	
	app.Run()
}
