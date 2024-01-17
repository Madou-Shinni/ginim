package domain

import (
	"errors"
	"github.com/Madou-Shinni/gin-quickstart/constants"
	"github.com/Madou-Shinni/gin-quickstart/pkg/model"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
	"github.com/Madou-Shinni/gin-quickstart/pkg/tools/snowflake"
	"gorm.io/gorm"
)

type Relationship struct {
	model.Model
	Owner  uint                       `json:"owner,string,omitempty"`                     // 拥有者
	Target uint                       `json:"target,string,omitempty"`                    // 目标
	Type   constants.RelationshipType `json:"type,omitempty" gorm:"column:type;type:int"` // 关系类型
	Friend *User                      `json:"friend,omitempty" gorm:"foreignKey:Target"`  // 好友
}

type PageRelationshipSearch struct {
	Relationship
	request.PageSearch
}

func (Relationship) TableName() string {
	return "relationship"
}

func (g *Relationship) AfterCreate(tx *gorm.DB) (err error) {
	// 添加好友关系双向
	// 好友也拥有我这个好友
	var relationship Relationship
	// 查询是否已经添加过了
	err = tx.Model(&Relationship{}).Where("owner = ? AND target = ?", g.Target, g.Owner).First(&relationship).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if relationship.ID != 0 {
		// 已经被添加过了,直接返回避免无限调用hooks
		return nil
	}

	if g.Type == constants.RelationshipTypeFriend {
		err = tx.Model(&Relationship{}).Create(&Relationship{
			Owner:  g.Target,
			Type:   g.Type,
			Target: g.Owner,
		}).Error
		if err != nil {
			return err
		}
	}

	// 添加会话 双向
	var conversation Conversation
	conversation.OwnerId = g.Owner
	conversation.TargetId = g.Target
	conversation.Type = constants.ConversationTypePrivate
	conversation.ConversationID = uint(snowflake.GenerateID())
	if g.Type == constants.RelationshipTypeFriend {
		// 添加自己的会话
		err = tx.Model(&Conversation{}).Create(&conversation).Error
		if err != nil {
			return err
		}
		// 添加好友的会话
		err = tx.Model(&Conversation{}).Create(&Conversation{
			OwnerId:        g.Target,
			Type:           constants.ConversationTypePrivate,
			TargetId:       g.Owner,
			ConversationID: conversation.ConversationID, // 会话id保持一致
		}).Error
		if err != nil {
			return err
		}
	}

	return
}
