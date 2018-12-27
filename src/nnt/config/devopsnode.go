package config

type DevopsNode struct {
	// 是否允许客户端直接访问本服务
	Client bool

	// 是否允许服务端访问本服务
	Server bool

	// 白名单
	Allow []string

	// 黑名单
	Deny []string
}
