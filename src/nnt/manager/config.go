package manager

import (
	"os"
)

var (
	// DEBUG模式
	DEBUG = false

	// DEVELOP模式，和debug的区别，develop用来部署开发服务器，debug用来做本地开发，会影响到app.json中对服务器的启动处理
	DEVELOP = false

	// PUBLISH模式，和release类似，除了会使用线上配置外，其他又和develop一致
	PUBLISH = false

	// 正式版模式
	DISTRIBUTION = false

	// 本地模式
	LOCAL = false

	// 容器部署
	DEVOPS = false

	// 内网测试容器部署
	DEVOPS_DEVELOP = false

	// 外网容器部署
	DEVOPS_RELEASE = true

	// sid过期时间，此框架中时间最小单位为秒
	SID_EXPIRE = 86400

	// clientid 过期时间
	CID_EXPIRE = 600

	// model含有最大fields的个数
	MODEL_FIELDS_MAX = 100

	// transaction超时时间
	TRANSACTION_TIMEOUT = 20

	// 是否允许客户端访问
	CLIENT_ALLOW = false

	// 是否允许服务端访问
	SERVER_ALLOW = true

	// 白名单
	ACCESS_ALLOW = make([]string, 0)

	// 黑名单
	ACCESS_DENY = make([]string, 0)

	// 服务端缓存目录
	CACHE = "cache"

	// 最大下载文件的大小
	FILESIZE_LIMIT = 10485760 //10M
)

func IsDebug() bool {
	return DEBUG || DEVELOP || PUBLISH
}

func IsRelease() bool {
	return DISTRIBUTION
}

func DebugValue(d interface{}, r interface{}) interface{} {
	if DISTRIBUTION {
		return r
	}
	return d
}

// 支持DEVOPS的架构判断
func IsDevops() bool {
	return os.Getenv("DEVOPS") != ""
}

func IsDevopsDevelop() bool {
	return os.Getenv("DEVOPS") != "" && os.Getenv("DEVOPS_RELEASE") == ""
}

func IsDevopsRelease() bool {
	return os.Getenv("DEVOPS_RELEASE") != ""
}

func IsLocal() bool {
	return os.Getenv("DEVOPS") == ""
}
