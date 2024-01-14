package domain

import (
	"github.com/Madou-Shinni/gin-quickstart/constants"
	"github.com/Madou-Shinni/gin-quickstart/pkg/model"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
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
