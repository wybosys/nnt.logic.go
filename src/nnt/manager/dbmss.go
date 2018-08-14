package manager

import "nnt/core/kernel"

type prvDbmss struct {
}

var Dbmss *prvDbmss = &prvDbmss{}

func (*prvDbmss) Start(cfg *kernel.JsonObject) {

}
