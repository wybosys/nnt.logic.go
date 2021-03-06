package manager

import (
	"github.com/wybosys/nnt.logic.go/nnt/config"
	"github.com/wybosys/nnt.logic.go/nnt/core/fs"
	"github.com/wybosys/nnt.logic.go/nnt/core/json"
	"github.com/wybosys/nnt.logic.go/nnt/core/kernel"
	"github.com/wybosys/nnt.logic.go/nnt/core/url"
	"log"
)

type prvApp struct {
	// app.json的设置
	Appcfg string

	// devops.json的设置
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
	cfg, err := json.ReadFile(appcfg)
	if err != nil {
		log.Fatal(err.Error())
	}
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
		cfg, _ := kernel.ToJsonObject(content)
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
		if v, ok := cfg.CheckGet("path"); ok {
			t, _ := v.String()
			config.DOMAIN = t[16:]
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
