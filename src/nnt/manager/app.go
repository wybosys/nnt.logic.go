package manager

import (
	"nnt/core/fs"
	"nnt/core/url"
	"nnt/core/kernel"
	"nnt/logger"
)

type prvApp struct {
	Appcfg string
	Devcfg string

	// 保存解析好的配置
	appcfgobj *kernel.JsonObject
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
	appcfg, devcfg := url.Expand(self.Appcfg), ""
	if self.Devcfg != "" {
		devcfg = url.Expand(self.Devcfg)
	}

	if Config.DEBUG = kernel.ArgsContains("--debug"); Config.DEBUG {
		logger.Log("debug模式启动")
	} else if Config.DEVELOP = kernel.ArgsContains("--develop"); Config.DEVELOP {
		logger.Log("develop模式启动")
	} else if Config.PUBLISH = kernel.ArgsContains("--publish"); Config.PUBLISH {
		logger.Log("publish模式启动")
	}
	if Config.DISTRIBUTION = !Config.IsDebug(); Config.DISTRIBUTION {
		logger.Log("distribution模式启动")
	}
	if Config.LOCAL = Config.IsLocal(); Config.LOCAL {
		logger.Log("LOCAL 环境")
	}
	if Config.DEVOPS = Config.IsDevops(); Config.DEVOPS {
		logger.Log("DEVOPS 环境")
	}
	if Config.DEVOPS_DEVELOP = Config.IsDevopsDevelop(); Config.DEVOPS_DEVELOP {
		logger.Log("DEVOPS DEVELOP 环境")
	}
	if Config.DEVOPS_RELEASE = Config.IsDevopsRelease(); Config.DEVOPS_RELEASE {
		logger.Log("DEVOPS RELEASE 环境")
	}

	// 读取配置
	content, err := fs.FileGetContents(appcfg)
	if err != nil {
		logger.Fatal(err, "读取APP配置失败")
	}
	cfg := kernel.ToJsonObject(content)
	self.appcfgobj = cfg

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
		Config.CACHE = url.Expand(t)
	}

	// 读取开发配置
	if devcfg != "" {
		content, _ := fs.FileGetContents(devcfg)
		cfg := kernel.ToJsonObject(content)
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

func (self *prvApp) Start() {
	cfg := self.appcfgobj
	if v, ok := cfg.CheckGet("logger"); ok {
		t, _ := v.Array()
		Loggers.Start(t)
	}
	if v, ok := cfg.CheckGet("dbms"); ok {
		t, _ := v.Array()
		Dbmss.Start(t)
	}
	if v, ok := cfg.CheckGet("server"); ok {
		t, _ := v.Array()
		Servers.Start(t)
	}
}
