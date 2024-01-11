package domain

import (
	"github.com/Madou-Shinni/gin-quickstart/constants"
	"github.com/Madou-Shinni/gin-quickstart/pkg/model"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
)

type Message struct {
	model.Model
	From      uint                  `json:"from,omitempty"`      // 发送者
	To        uint                  `json:"to,omitempty"`        // 接受者
	Type      constants.MessageType `json:"type,omitempty"`      // 消息类型
	MediaType constants.MediaType   `json:"mediaType,omitempty"` // 媒体类型
	Content   string                `json:"content,omitempty"`   // 内容
	Sender    bool                  `json:"sender"`              // 是否是自己发送
}

type PageMessageSearch struct {
	Message
	request.PageSearch
}

func (Message) TableName() string {
	return "message"
}
