package domain

import (
	"github.com/Madou-Shinni/gin-quickstart/pkg/model"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
)

type User struct {
	model.Model
	Name string `json:"name"`
}

type PageUserSearch struct {
	User
	request.PageSearch
}

func (User) TableName() string {
	return "user"
}
