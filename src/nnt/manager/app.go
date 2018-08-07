package manager

import (
	"nnt/core"
	"os"
)

type _App struct {
	Appcfg string
	Devcfg string
}

var (
	gs_app = _App{
		Appcfg: "~/app.json",
		Devcfg: "~/devops.json",
	}
)

func App() *_App {
	return &gs_app
}

func (self *_App) LoadConfig() {
	appcfg, devcfg := core.Urls.Expand(self.Appcfg), ""
	if self.Devcfg != "" {
		devcfg = core.Urls.Expand(self.Devcfg)
	}


}

func (_ *_App) Start() {

}

func (_ *_App) Stop() {

}
