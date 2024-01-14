package domain

import (
	"github.com/Madou-Shinni/gin-quickstart/constants"
	"github.com/Madou-Shinni/gin-quickstart/pkg/model"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
	"gorm.io/gorm"
)

type Group struct {
	model.Model
	Owner uint   `json:"owner,omitempty"` // 拥有者
	Name  string `json:"name,omitempty"`  // 群名称
}

type PageGroupSearch struct {
	Group
	request.PageSearch
}

func (Group) TableName() string {
	return "group"
}

func (g *Group) AfterCreate(tx *gorm.DB) (err error) {
	// 添加关系
	err = tx.Model(&Relationship{}).Create(&Relationship{
		Owner:  g.Owner,
		Target: g.ID,
		Type:   constants.RelationshipGroup,
	}).Error
	if err != nil {
		return err
	}
	// 添加会话
	err = tx.Model(&Conversation{}).Create(&Conversation{
		OwnerId:       g.Owner,
		TargetId:      g.ID,
		Type:          constants.ConversationTypeGroup,
		LastMessageId: 0,
	}).Error
	if err != nil {
		return err
	}

	return
}
