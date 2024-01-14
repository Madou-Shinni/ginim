package domain

import (
	"github.com/Madou-Shinni/gin-quickstart/constants"
	"github.com/Madou-Shinni/gin-quickstart/pkg/model"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
	"gorm.io/gorm"
)

type Message struct {
	model.Model
	From           uint                  `json:"from,string,omitempty"`                 // 发送者
	To             uint                  `json:"to,string,omitempty"`                   // 接受者
	Type           constants.MessageType `json:"type,omitempty"`                        // 消息类型
	MediaType      constants.MediaType   `json:"mediaType,omitempty"`                   // 媒体类型
	Content        string                `json:"content,omitempty"`                     // 内容
	ConversationId uint                  `json:"conversationId" gorm:"conversation_id"` // 会话id
}

type PageMessageSearch struct {
	Message
	request.PageSearch
}

func (Message) TableName() string {
	return "message"
}

func (m *Message) BeforeCreate(tx *gorm.DB) (err error) {
	// 查询会话
	var conversation Conversation
	if m.Type == constants.MessageTypeGroup {
		// 群聊消息，以群的纬度查询会话
		err = tx.Model(&Conversation{}).Where("owner_id = ?", m.To).First(&conversation).Error
		if err != nil {
			return err
		}
	}
	if m.Type == constants.MessageTypePrivate {
		// 私聊消息，以用户的纬度查询会话
		err = tx.Model(&Conversation{}).Where("owner_id = ?", m.From).First(&conversation).Error
		if err != nil {
			return err
		}
	}

	// 设置会话id
	m.ConversationId = conversation.ID

	return
}

func (m *Message) AfterCreate(tx *gorm.DB) (err error) {
	// 更新会话最后一条消息，为最新的消息
	err = tx.Model(&Conversation{}).Where("id = ?", m.ConversationId).Update("last_message_id", m.ID).Error
	if err != nil {
		return err
	}

	return
}
