package manager

import (
	"nnt/core"
	"os"
	"nnt/core/fs"
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
	content, _ := core.FileGetContents(appcfg)
	cfg := core.ToJsonObject(content)

	// 读取系统配置
	c := cfg.Get("config")
	if v, ok := c.CheckGet("sidexpire"); ok {
		t, _ := v.Int()
		Config.SID_EXPIRE = uint(t)
	}
	if v, ok := c.CheckGet("cidexpire"); ok {
		t, _ := v.Int()
		Config.CID_EXPIRE = uint(t)
	}
	if v, ok := c.CheckGet("cache"); ok {
		t, _ := v.String()
		Config.CACHE = core.Urls.Expand(t)
	}

	// 读取开发配置
	if devcfg != "" {
		content, _ := core.FileGetContents(devcfg)
		cfg := core.ToJsonObject(content)
		if v, ok := cfg.CheckGet("client"); ok {
			t, _ := v.Bool()
			Config.CLIENT_ALLOW = t
		}
		if v, ok := cfg.CheckGet("server"); ok {
			t, _ := v.Bool()
			Config.SERVER_ALLOW = t
		}
		if v, ok := cfg.CheckGet("allow"); ok {
			t, _ := v.StringArray()
			Config.ACCESS_ALLOW = t
		}
		if v, ok := cfg.CheckGet("deny"); ok {
			t, _ := v.StringArray()
			Config.ACCESS_DENY = t
		}
	}

	// 缓存目录
	if !fs.IsDir(Config.CACHE) {
		fs.MkDir(Config.CACHE)
	}
}

func (_ *prvApp) Start() {

}

func (_ *prvApp) Stop() {

}
