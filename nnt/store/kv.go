package store

import "github.com/wybosys/nnt.logic.go/nnt/core/kernel"

type IKv interface {
	IDbms

	Get(key string) *kernel.Variant
	Set(key string, val *kernel.Variant)
	GetSet(key string, val *kernel.Variant) *kernel.Variant
	Del(key string) DbExecuteStat
	AutoInc(key string, delta int) int
	Inc(key string, delta int) int
}

type AbstractKv struct {
	AbstractDbms
}
