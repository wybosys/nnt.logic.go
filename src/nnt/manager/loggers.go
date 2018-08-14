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

	}

	logger.Warn = func(s string) {

	}

	logger.Info = func(s string) {

	}

	logger.Fatal = func(e error, s string) {

	}

	logger.Exception = func(e error) {

	}

	logger.Error = func(e error) {

	}
}
