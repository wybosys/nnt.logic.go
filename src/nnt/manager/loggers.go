package manager

import (
	"nnt/logger"
	"nnt/core/kernel"
	"nnt/config"
	"nnt/core/entry"
	"log"
)

type prvLoggers struct {
	servers map[string]logger.ILogger
}

var Loggers *prvLoggers = &prvLoggers{
	make(map[string]logger.ILogger),
}

func (self *prvLoggers) Start(cfg []*kernel.JsonObject) {
	for _, e := range cfg {
		if !config.NodeIsEnable(e) {
			continue;
		}

		entrykey, _ := e.GetString("entry")
		if hdl := entry.Instance(entrykey); hdl != nil {
			nodeid, _ := e.GetString("id")
			loggerhdl := hdl.(logger.ILogger)
			if loggerhdl.Config(e) {
				self.servers[nodeid] = loggerhdl
				loggerhdl.Start()
			} else {
				log.Fatal(nodeid + " 配置失败")
			}
		} else {
			log.Fatal("无法实例化 " + entrykey)
		}
	}
}

func init() {
	logger.Log = func(s string) {
		for _, e := range Loggers.servers {
			e.Log(s)
		}
	}

	logger.Warn = func(s string) {
		for _, e := range Loggers.servers {
			e.Warn(s)
		}
	}

	logger.Info = func(s string) {
		for _, e := range Loggers.servers {
			e.Info(s)
		}
	}

	logger.Fatal = func(err error, s string) {
		for _, e := range Loggers.servers {
			e.Fatal(err, s)
		}
	}

	logger.Exception = func(err error) {
		for _, e := range Loggers.servers {
			e.Exception(err)
		}
	}

	logger.Error = func(err error) {
		for _, e := range Loggers.servers {
			e.Error(err)
		}
	}
}
