package config

import (
	"nnt/core/kernel"
	"nnt/core/stl"
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
			return DEBUG
		}
		if e == "develop" {
			return DEVELOP
		}
		if e == "publish" {
			return PUBLISH
		}
		if e == "distribution" {
			return DISTRIBUTION
		}
		if e == "release" {
			return PUBLISH || DISTRIBUTION
		}
		if e == "devops" {
			return DEVOPS
		}
		if e == "devops-develop" {
			return DEVOPS_DEVELOP
		}
		if e == "devops-release" {
			return DEVOPS_RELEASE
		}
		if e == "local" {
			return LOCAL
		}
		if kernel.ArgsContains("--" + e) {
			return true
		}
		return false
	})

	return fnd
}
