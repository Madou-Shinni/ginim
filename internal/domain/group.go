package domain

import (
	"github.com/Madou-Shinni/gin-quickstart/pkg/model"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
)

type Group struct {
	model.Model
}

type PageGroupSearch struct {
	Group
	request.PageSearch
}

func (Group) TableName() string {
	return "group"
}