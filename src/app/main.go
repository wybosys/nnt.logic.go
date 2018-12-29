package app

import (
	"nnt/manager"
)

func Main() {
	hdl := manager.App()
	hdl.LoadConfig();
	hdl.Start();
}
