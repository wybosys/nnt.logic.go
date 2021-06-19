package app

import (
	"github.com/wybosys/nnt.logic.go/nnt/manager"
)

func Main() {
	hdl := manager.App()
	hdl.LoadConfig()
	hdl.Start()
}
