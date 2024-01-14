package domain

import (
	"github.com/Madou-Shinni/gin-quickstart/constants"
	"github.com/Madou-Shinni/gin-quickstart/pkg/model"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
	"gorm.io/gorm"
)

type Relationship struct {
	model.Model
	Owner  uint                       `json:"owner,omitempty"`                            // 拥有者
	Target uint                       `json:"target,omitempty"`                           // 目标
	Type   constants.RelationshipType `json:"type,omitempty" gorm:"column:type;type:int"` // 关系类型
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

	return
}
