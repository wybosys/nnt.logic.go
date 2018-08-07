package manager

import (
	"nnt"
	"os"
)

type _Config struct {
	// DEBUG模式
	DEBUG bool

	// DEVELOP模式，和debug的区别，develop用来部署开发服务器，debug用来做本地开发，会影响到app.json中对服务器的启动处理
	DEVELOP bool

	// PUBLISH模式，和release类似，除了会使用线上配置外，其他又和develop一致
	PUBLISH bool

	// 正式版模式
	DISTRIBUTION bool

	// 本地模式
	LOCAL bool

	// 容器部署
	DEVOPS bool

	// 内网测试容器部署
	DEVOPS_DEVELOP bool

	// 外网容器部署
	DEVOPS_RELEASE bool

	// sid过期时间，此框架中时间最小单位为秒
	SID_EXPIRE uint

	// clientid 过期时间
	CID_EXPIRE uint

	// model含有最大fields的个数
	MODEL_FIELDS_MAX uint

	// transaction超时时间
	TRANSACTION_TIMEOUT uint

	// 是否允许客户端访问
	CLIENT_ALLOW bool

	// 是否允许服务端访问
	SERVER_ALLOW bool

	// 白名单
	ACCESS_ALLOW []string

	// 黑名单
	ACCESS_DENY []string

	// 服务端缓存目录
	CACHE string

	// 最大下载文件的大小
	FILESIZE_LIMIT uint
}

var Config = &_Config{
	DEBUG:               false,
	DEVELOP:             false,
	PUBLISH:             false,
	DISTRIBUTION:        false,
	LOCAL:               false,
	DEVOPS:              false,
	DEVOPS_DEVELOP:      false,
	DEVOPS_RELEASE:      true,
	SID_EXPIRE:          86400,
	CID_EXPIRE:          600,
	MODEL_FIELDS_MAX:    100,
	TRANSACTION_TIMEOUT: 20,
	CLIENT_ALLOW:        false,
	SERVER_ALLOW:        true,
	ACCESS_ALLOW:        make([]string, 0),
	ACCESS_DENY:         make([]string, 0),
	CACHE:               "cache",
	FILESIZE_LIMIT:      10485760, //10M
}

func (self *_Config) IsDebug() bool {
	return self.DEBUG || self.DEVELOP || self.PUBLISH
}

func (self *_Config) IsRelease() bool {
	return self.DISTRIBUTION
}

func (self *_Config) DebugValue(d nnt.Any, r nnt.Any) nnt.Any {
	if self.DISTRIBUTION {
		return r
	}
	return d
}

// 支持DEVOPS的架构判断
func (self *_Config) IsDevops() bool {
	return os.Getenv("DEVOPS") != ""
}

func (self *_Config) IsDevopsDevelop() bool {
	return os.Getenv("DEVOPS") != "" && os.Getenv("DEVOPS_RELEASE") == ""
}

func (self *_Config) IsDevopsRelease() bool {
	return os.Getenv("DEVOPS_RELEASE") != ""
}

func (self *_Config) IsLocal() bool {
	return os.Getenv("DEVOPS") == ""
}
