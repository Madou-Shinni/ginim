package domain

import (
	"github.com/Madou-Shinni/gin-quickstart/constants"
	"github.com/Madou-Shinni/gin-quickstart/pkg/model"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
)

// Conversation 会话
type Conversation struct {
	model.Model
	ConversationID uint                       `json:"conversationId,string" gorm:"column:conversation_id"` // 会话id
	OwnerId        uint                       `json:"ownerId,string" gorm:"owner_id"`
	TargetId       uint                       `json:"targetId,string" gorm:"target_id"`
	Type           constants.ConversationType `json:"type" gorm:"column:type;type:int"`     // 会话类型
	LastMessageId  uint                       `json:"lastMessageId" gorm:"last_message_id"` // 最后一条消息id
	LastMessage    Message                    `json:"lastMessage" gorm:"foreignKey:ID;references:LastMessageId"`
	UnreadCount    int                        `json:"unreadCount" gorm:"-"` // 未读消息数
	Name           string                     `json:"name" gorm:"-"`        // 会话名称
}

type PageConversationSearch struct {
	Conversation
	request.PageSearch
}

func (Conversation) TableName() string {
	return "conversation"
}
