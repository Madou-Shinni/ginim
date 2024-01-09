package domain

import (
	"github.com/Madou-Shinni/gin-quickstart/pkg/model"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
)

type Relationship struct {
	model.Model
	Owner  uint // 拥有者
	Target uint // 目标
	Type   uint // 关系类型
}

type PageRelationshipSearch struct {
	Relationship
	request.PageSearch
}

func (Relationship) TableName() string {
	return "relationship"
}
