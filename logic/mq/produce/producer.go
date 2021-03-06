package produce

import (
	"fmt"
	"github.com/brbgithub/goim/conf"
	"github.com/brbgithub/goim/logic/db"
	"github.com/brbgithub/goim/public/logger"
	"github.com/brbgithub/goim/public/transfer"
	jsoniter "github.com/json-iterator/go"
	"github.com/nsqio/go-nsq"
)

var producer *nsq.Producer

func init() {
	var err error
	cfg := nsq.NewConfig()
	producer, err = nsq.NewProducer(conf.NSQIP, cfg)
	if err != nil {
		panic("nsq new panic")
	}

	err = producer.Ping()
	if err != nil {
		panic("nsq ping panic")
	}
}

// PublishMessage 发布消息投递
func PublishMessage(message transfer.Message) {
	// 获取设备连接的连接层服务器
	connectIP, err := db.RedisClient.Get(db.DeviceIdPre + fmt.Sprint(message.DeviceId)).Result()
	if err != nil {
		logger.Sugar.Error(err)
		return
	}
	// 获取服务器消费的topic
	topic := connectIP + ".message"

	body, err := jsoniter.Marshal(message)
	if err != nil {
		logger.Sugar.Error(err)
		return
	}
	err = producer.Publish(topic, body)
	if err != nil {
		logger.Sugar.Error(err)
	}
}

// PublishMessage 发布消息发送回执
func PublishMessageSendACK(ack transfer.MessageSendACK) {
	// 获取设备连接的连接层服务器
	connectIP, err := db.RedisClient.Get(db.DeviceIdPre + fmt.Sprint(ack.DeviceId)).Result()
	if err != nil {
		logger.Sugar.Error(err)
		return
	}
	// 获取服务器消费的topic
	topic := connectIP + ".message_send_ack"

	body, err := jsoniter.Marshal(ack)
	if err != nil {
		logger.Sugar.Error(err)
		return
	}
	err = producer.Publish(topic, body)
	if err != nil {
		logger.Sugar.Error(err)
	}
}
