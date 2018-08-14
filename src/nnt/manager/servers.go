package manager

import (
	"nnt/core/kernel"
	"nnt/server"
	"nnt/config"
	"nnt/core/entry"
	"log"
)

type prvServers struct {
	servers map[string]server.IServer
}

var Servers *prvServers = &prvServers{
	make(map[string]server.IServer),
}

func (self *prvServers) Start(cfg []*kernel.JsonObject) {
	for _, e := range cfg {
		if !config.NodeIsEnable(e) {
			continue;
		}

		entrykey, _ := e.GetString("entry")
		if hdl := entry.Instance(entrykey); hdl != nil {
			nodeid, _ := e.GetString("id")
			srvhdl := hdl.(server.IServer)
			if srvhdl.Config(e) {
				self.servers[nodeid] = srvhdl
				srvhdl.Start()
			} else {
				log.Fatal(nodeid + " 配置失败")
			}
		} else {
			log.Fatal("无法实例化 " + entrykey)
		}
	}
}
