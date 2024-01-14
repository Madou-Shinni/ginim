package domain

import (
	"github.com/Madou-Shinni/gin-quickstart/constants"
	"github.com/Madou-Shinni/gin-quickstart/pkg/model"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
)

// Conversation 会话
type Conversation struct {
	model.Model
	OwnerId       uint                       `json:"ownerId" gorm:"owner_id"`
	TargetId      uint                       `json:"targetId" gorm:"target_id"`
	Type          constants.ConversationType `json:"type" gorm:"column:type;type:int"`     // 会话类型
	LastMessageId uint                       `json:"lastMessageId" gorm:"last_message_id"` // 最后一条消息id
}

type PageConversationSearch struct {
	Conversation
	request.PageSearch
}

func (Conversation) TableName() string {
	return "conversation"
}
