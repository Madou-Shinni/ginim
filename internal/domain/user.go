package domain

import (
	"github.com/Madou-Shinni/gin-quickstart/pkg/model"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
	"github.com/Madou-Shinni/gin-quickstart/pkg/tools/snowflake"
	"gorm.io/gorm"
)

type User struct {
	ID        uint             `gorm:"primarykey" json:"id,string" form:"id"`           // 主键
	CreatedAt *model.LocalTime `json:"createdAt" form:"createdAt" swaggerignore:"true"` // 创建时间
	UpdatedAt *model.LocalTime `json:"updatedAt" form:"updatedAt" swaggerignore:"true"` // 修改时间
	DeletedAt gorm.DeletedAt   `gorm:"index" json:"deletedAt" form:"deletedAt" swaggerignore:"true"`
	Name      string           `json:"name" gorm:"name;type:varchar(255)" form:"name"`
	Avatar    string           `json:"avatar" gorm:"column:avatar" form:"avatar"`
	GithubId  uint             `json:"githubId" gorm:"github_id"` // github用户id
}

type PageUserSearch struct {
	User
	request.PageSearch
}

func (User) TableName() string {
	return "user"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	id := snowflake.GenerateID()
	u.ID = uint(id)
	return
}
