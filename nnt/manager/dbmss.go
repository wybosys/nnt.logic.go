package manager

import "github.com/wybosys/nnt.logic.go/nnt/core/kernel"

type prvDbmss struct {
}

var Dbmss *prvDbmss = &prvDbmss{}

func (*prvDbmss) Start(cfg []*kernel.JsonObject) {

}
