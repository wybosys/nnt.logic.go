package config

type Node struct {
	// 节点id
	Id string

	// 实体对应的类
	Entry string

	// 开发模式配置, 如果不配置，则代表任何模式都启用，否则只有命中的模式才启用
	Enable string
}
