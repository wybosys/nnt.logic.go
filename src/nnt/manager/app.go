package manager

import (
	"nnt/core"
	"os"
)

type prvApp struct {
	Appcfg string
	Devcfg string
}

var (
	gs_app = prvApp{
		Appcfg: "~/app.json",
		Devcfg: "~/devops.json",
	}
)

func App() *prvApp {
	return &gs_app
}

func (self *prvApp) LoadConfig() {
	appcfg, devcfg := core.Urls.Expand(self.Appcfg), ""
	if self.Devcfg != "" {
		devcfg = core.Urls.Expand(self.Devcfg)
	}

	if Config.DEBUG = core.ArrayT.Contains(os.Args, "--debug", nil); Config.DEBUG {

	}
}

func (_ *prvApp) Start() {

}

func (_ *prvApp) Stop() {

}
