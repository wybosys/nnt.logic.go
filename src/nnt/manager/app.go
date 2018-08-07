package manager

import (
	"nnt/core"
	"os"
	"log"
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
		core.Logger.Log("debug模式启动")
	} else if Config.DEVELOP = core.ArrayT.Contains(os.Args, "--develop", nil); Config.DEVELOP {
		core.Logger.Log("develop模式启动")
	} else if Config.PUBLISH = core.ArrayT.Contains(os.Args, "--publish", nil); Config.PUBLISH {
		core.Logger.Log("publish模式启动")
	}
	if Config.DISTRIBUTION = !Config.IsDebug(); Config.DISTRIBUTION {
		core.Logger.Log("distribution模式启动")
	}
	if Config.LOCAL = Config.IsLocal(); Config.LOCAL {
		core.Logger.Log("LOCAL 环境")
	}
	if Config.DEVOPS = Config.IsDevops(); Config.DEVOPS {
		core.Logger.Log("DEVOPS 环境")
	}
	if Config.DEVOPS_DEVELOP = Config.IsDevopsDevelop(); Config.DEVOPS_DEVELOP {
		core.Logger.Log("DEVOPS DEVELOP 环境")
	}
	if Config.DEVOPS_RELEASE = Config.IsDevopsRelease(); Config.DEVOPS_RELEASE {
		core.Logger.Log("DEVOPS RELEASE 环境")
	}

	// 读取配置


	// 缓存目录
}

func (_ *prvApp) Start() {

}

func (_ *prvApp) Stop() {

}
