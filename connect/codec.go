package connect

import (
	"errors"
	"net"
)

const (
	TypeLen = 2 // 消息类型字节数组长度
	LenLen = 2 // 消息长度字节数组长度
	HeadLen = 4 // 消息头部字节数组长度（消息类型字节数组长度+消息长度字节数组长度）
	ContentMaxLen = 4092 // 消息体最大长度
	BufLen = ContentMaxLen + 4 // 缓冲buffer字节数组长度
)

var ErrOutOfSize = errors.New("package content out of size") // package的content字节数组过大

type Codec struct {
	Conn net.Conn
	ReadBuf buffer // 读缓冲
	WriteBuf []byte // 写缓冲
}
