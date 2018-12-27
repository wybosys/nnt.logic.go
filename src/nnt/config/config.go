package config

import (
	"nnt/core/kernel"
	"nnt/core/stl"
	"nnt/manager"
	"strings"
)

func NodeIsEnable(cfg *kernel.JsonObject) bool {
	v, ok := cfg.CheckGet("enable")
	if !ok {
		return true;
	}

	enable, _ := v.String()
	conds := stl.StringArray(strings.Split(enable, ","))
	_, fnd := conds.Query(func(idx int, e string) bool {
		if e == "" {
			return false
		}
		if e == "debug" {
			return manager.DEBUG
		}
		if e == "develop" {
			return manager.DEVELOP
		}
		if e == "publish" {
			return manager.PUBLISH
		}
		if e == "distribution" {
			return manager.DISTRIBUTION
		}
		if e == "release" {
			return manager.PUBLISH || manager.DISTRIBUTION
		}
		if e == "devops" {
			return manager.DEVOPS
		}
		if e == "devops-develop" {
			return manager.DEVOPS_DEVELOP
		}
		if e == "devops-release" {
			return manager.DEVOPS_RELEASE
		}
		if e == "local" {
			return manager.LOCAL
		}
		if kernel.ArgsContains("--" + e) {
			return true
		}
		return false
	})

	return fnd
}
