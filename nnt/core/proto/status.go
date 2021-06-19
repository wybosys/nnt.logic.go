package proto

type STATUS int

const (
	UNKNOWN             STATUS = -1000
	EXCEPTION                  = -999  // 遇到了未处理的异常
	ROUTER_NOT_FOUND           = -998; // 没有找到路由
	CONTEXT_LOST               = -997  // 上下文丢失
	MODEL_ERROR                = -996  // 恢复模型失败
	PARAMETER_NOT_MATCH        = -995  // 参数不符合要求
	NEED_AUTH                  = -994  // 需要登陆
	TYPE_MISMATCH              = -993  // 参数类型错误
	FILESYSTEM_FAILED          = -992  // 文件系统失败
	FILE_NOT_FOUND             = -991  // 文件不存在
	ARCHITECT_DISMATCH         = -990  // 代码不符合标准架构
	SERVER_NOT_FOUND           = -989  // 没有找到服务器
	LENGTH_OVERFLOW            = -988  // 长度超过限制
	TARGET_NOT_FOUND           = -987  // 目标对象没有找到
	PERMISSIO_FAILED           = -986  // 没有权限
	WAIT_IMPLEMENTION          = -985  // 等待实现
	ACTION_NOT_FOUND           = -984  // 没有找到动作
	TARGET_EXISTS              = -983  // 已经存在
	STATE_FAILED               = -982  // 状态错误
	UPLOAD_FAILED              = -981  // 上传失败
	MASK_WORD                  = -980  // 有敏感词
	SELF_ACTION                = -979  // 针对自己进行操作
	PASS_FAILED                = -978  // 验证码匹配失败
	OVERFLOW                   = -977  // 数据溢出
	AUTH_EXPIRED               = -976  // 授权过期
	SIGNATURE_ERROR            = -975  // 签名错误

	IM_CHECK_FAILED = -899 // IM检查输入的参数失败
	IM_NO_RELEATION = -898 // IM检查双方不存在关系

	SOCK_WRONG_PORTOCOL = -860 // SOCKET请求了错误的通讯协议
	SOCK_AUTH_TIMEOUT   = -859 // 因为连接后长期没有登录，所以服务端主动断开了链接
	SOCK_SERVER_CLOSED  = -858 // 服务器关闭

	THIRD_FAILED = -5 // 第三方出错
	MULTIDEVICE  = -4 // 多端登陆
	HFDENY       = -3 // 高频调用被拒绝（之前的访问还没有结束) high frequency deny
	TIMEOUT      = -2 // 超时
	FAILED       = -1 // 一般失败
	OK           = 0  // 成功
)
