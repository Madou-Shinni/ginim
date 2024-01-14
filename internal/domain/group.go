package domain

import (
	"github.com/Madou-Shinni/gin-quickstart/constants"
	"github.com/Madou-Shinni/gin-quickstart/pkg/model"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
	"gorm.io/gorm"
)

type Group struct {
	model.Model
	Owner uint   `json:"owner,string,omitempty"` // 拥有者
	Name  string `json:"name,omitempty"`         // 群名称
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
	// 群拥有群主这个成员
	err = tx.Model(&Relationship{}).Create(&Relationship{
		Owner:  g.ID,    // 群id
		Target: g.Owner, // 群主id
		Type:   constants.RelationshipTypeGroup,
	}).Error
	if err != nil {
		return err
	}
	// 添加会话
	// 群拥有一个会话
	err = tx.Model(&Conversation{}).Create(&Conversation{
		OwnerId:       g.ID, // 群id,这个会话是属于群的,群成员通过群id来索引会话
		TargetId:      0,
		Type:          constants.ConversationTypeGroup,
		LastMessageId: 0,
	}).Error
	if err != nil {
		return err
	}

	return
}
