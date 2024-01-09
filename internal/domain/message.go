package domain

import (
	"github.com/Madou-Shinni/gin-quickstart/constants"
	"github.com/Madou-Shinni/gin-quickstart/pkg/model"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
)

type Message struct {
	model.Model
	From      uint                  // 发送者
	To        uint                  // 接受者
	Type      constants.MessageType // 消息类型
	MediaType constants.MediaType   // 媒体类型
	Content   string                // 内容
}

type PageMessageSearch struct {
	Message
	request.PageSearch
}

func (Message) TableName() string {
	return "message"
}
