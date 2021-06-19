package config

import "github.com/wybosys/nnt.logic.go/nnt/core/kernel"

type AppNodes struct {
	// 全局配置信息
	Config kernel.JsonObject

	// 服务配置信息
	Server []kernel.JsonObject

	// 数据库配置信息
	Dbms []kernel.JsonObject

	// 日志配置信息
	Logger []kernel.JsonObject
}
