package conf

import "os"

// connect和logic公用配置
var (
	MySQL = "root:datatom2019@tcp(192.168.72.131:3306)/im?charset=utf8&parseTime=true"
	NSQIP = "127.0.0.1:4150"
	RedisIP = "192.168.72.131:6379"
	DeviceIdPre = "connect:device_id:"
)

// connect配置
var (
	ConnectRPCClientIPs = []string{
		"127.0.0.1:60000",
	}
	ConnectTCPListenIP = "127.0.0.1"
	ConnectTCPListenPort = "50000"
)

// logic配置
var (
	LogicRPCServerIP = "127.0.0.1:60000"
	LogicHTTPListenIP = "127.0.0.1:8000"
)

func init() {
	env := os.Getenv("im_env")
	if env == "dev"{
		initDevelopConf()
	}

	if env == "pro"{
		initProductConf()
	}
}

func initDevelopConf() {

}

func initProductConf() {

}