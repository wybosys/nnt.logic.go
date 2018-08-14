package manager

import "nnt/core/kernel"

type prvServers struct {
}

var Servers *prvServers = &prvServers{}

func (*prvServers) Start(cfg []*kernel.JsonObject) {

}
