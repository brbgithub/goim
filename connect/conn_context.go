package connect

import (
	"github.com/brbgithub/goim/public/logger"
	"net"
	"time"
)

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
	IsSignIn bool // 是否
	DeviceId int64 // 设备id
	UserId int64 // 用户id
}

// Package 消息包
type Package struct {
	Code int // 消息类型
	Content []byte // 消息体
}

func NewConnContext(conn *net.TCPConn) *ConnContext {
	codec := NewCodec(conn)
	return &ConnContext{Codec:codec}
}

// DoConn 处理TCP连接
func (c *ConnContext) DoConn() {
	defer RecoverPanic()

	c.HandleConnect()

	for {
		err := c.Codec.Conn.SetReadDeadline(time.Now().Add(ReadDeadline))
		if err != nil{

		}
	}
}

// HandleConnect 建立连接
func (c *ConnContext) HandleConnect() {
	logger.Logger.Info("tcp connect")
}

// Realease 释放TCP连接
func (c *ConnContext) Realease() {
	delete(c.DeviceId)
	err := c.Codec.Conn.Close()
	if err != nil{
		logger.Sugar.Error(err)
	}


}
