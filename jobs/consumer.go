package jobs

import (
	"github.com/Madou-Shinni/gin-quickstart/constants"
	"github.com/Madou-Shinni/gin-quickstart/pkg/global"
	"github.com/Madou-Shinni/gin-quickstart/pkg/tools/message_queue"
	"github.com/Madou-Shinni/go-logger"
	"go.uber.org/zap"
)

func Consumer() {
	if global.Rdb == nil {
		logger.Error("redis连接失败")
		return
	}
	// 从消息队列中获取消息
	rdb := global.Rdb
	subscribe := message_queue.RedisMessagePSubscribeChannels(rdb, string(constants.ChannelMessage))

	// 处理消息
	for {
		message := <-subscribe
		switch constants.Channel(message.Channel) {
		case constants.ChannelMessage:
			logger.Info("message channel message: ", zap.Any("message", message))
		}

		// 将消息发送给客户端
		return
	}
}
