package connect

import "time"

const (
	ReadDeadline  = 10 * time.Minute
	WriteDeadline = 10 * time.Second
)

// 消息协议
const (
	CodeSignIn         = 1 // 设备登录
	CodeSignInACK      = 2 // 设备登录回执
	CodeSyncTrigger    = 3 // 消息同步触发
	CodeHeartbeat      = 4 // 心跳
	CodeHeartbeatACK   = 5 // 心跳回执
	CodeMessageSend    = 6 // 消息发送
	CodeMessageSendACK = 7 // 消息发送回执
	CodeMessage        = 8 // 消息投递
	CodeMessageACK     = 9 // 消息投递回执
)

// ConnContext 连接上下文
type ConnContext struct {
	Codec *Codec // 编解码器
}
