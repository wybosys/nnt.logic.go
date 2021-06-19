package kernel

type ISerializableObject interface {
	// 序列化
	Serialize() []byte

	// 反序列化
	Unserialize(str []byte) bool
}
