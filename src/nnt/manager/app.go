package manager

import (
	"log"
	"nnt/config"
	"nnt/core/fs"
	"nnt/core/kernel"
	"nnt/core/url"
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

	if config.DEBUG = kernel.ArgsContains("--debug"); config.DEBUG {
		log.Print("debug模式启动")
	} else if config.DEVELOP = kernel.ArgsContains("--develop"); config.DEVELOP {
		log.Print("develop模式启动")
	} else if config.PUBLISH = kernel.ArgsContains("--publish"); config.PUBLISH {
		log.Print("publish模式启动")
	}
	if config.DISTRIBUTION = !config.IsDebug(); config.DISTRIBUTION {
		log.Print("distribution模式启动")
	}
	if config.LOCAL = config.IsLocal(); config.LOCAL {
		log.Print("LOCAL 环境")
	}
	if config.DEVOPS = config.IsDevops(); config.DEVOPS {
		log.Print("DEVOPS 环境")
	}
	if config.DEVOPS_DEVELOP = config.IsDevopsDevelop(); config.DEVOPS_DEVELOP {
		log.Print("DEVOPS DEVELOP 环境")
	}
	if config.DEVOPS_RELEASE = config.IsDevopsRelease(); config.DEVOPS_RELEASE {
		log.Print("DEVOPS RELEASE 环境")
	}

	// 读取配置
	content, err := fs.FileGetContents(appcfg)
	if err != nil {
		log.Print(err, "读取APP配置失败")
	}
	cfg := kernel.ToJsonObject(content)
	self.appcfgobj = cfg

	// 读取系统配置
	c := cfg.Get("config")
	if v, ok := c.CheckGet("sidexpire"); ok {
		t, _ := v.Int()
		config.SID_EXPIRE = t
	}
	if v, ok := c.CheckGet("cidexpire"); ok {
		t, _ := v.Int()
		config.CID_EXPIRE = t
	}
	if v, ok := c.CheckGet("cache"); ok {
		t, _ := v.String()
		config.CACHE = url.Expand(t)
	}

	// 读取开发配置
	if devcfg != "" {
		content, _ := fs.FileGetContents(devcfg)
		cfg := kernel.ToJsonObject(content)
		if v, ok := cfg.CheckGet("client"); ok {
			t, _ := v.Bool()
			config.CLIENT_ALLOW = t
		}
		if v, ok := cfg.CheckGet("server"); ok {
			t, _ := v.Bool()
			config.SERVER_ALLOW = t
		}
		if v, ok := cfg.CheckGet("allow"); ok {
			t, _ := v.StringArray()
			config.ACCESS_ALLOW = t
		}
		if v, ok := cfg.CheckGet("deny"); ok {
			t, _ := v.StringArray()
			config.ACCESS_DENY = t
		}
	}

	// 缓存目录
	if !fs.IsDir(config.CACHE) {
		fs.MkDir(config.CACHE)
	}
}

func (self *prvApp) Start() {
	cfg := self.appcfgobj
	if v, ok := cfg.CheckGet("logger"); ok {
		t, _ := v.JsArray()
		Loggers.Start(t)
	}
	if v, ok := cfg.CheckGet("dbms"); ok {
		t, _ := v.JsArray()
		Dbmss.Start(t)
	}
	if v, ok := cfg.CheckGet("server"); ok {
		t, _ := v.JsArray()
		Servers.Start(t)
	}
}
