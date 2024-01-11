package data

import (
	"errors"
	"fmt"
	"github.com/Madou-Shinni/gin-quickstart/internal/domain"
	"github.com/Madou-Shinni/gin-quickstart/pkg/global"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
	"github.com/Madou-Shinni/gin-quickstart/pkg/tools/pagelimit"
)

type UserRepo struct {
}

func (s *UserRepo) Create(user domain.User) error {
	return global.DB.Create(&user).Error
}

func (s *UserRepo) Delete(user domain.User) error {
	return global.DB.Delete(&user).Error
}

func (s *UserRepo) DeleteByIds(ids request.Ids) error {
	return global.DB.Delete(&[]domain.User{}, ids.Ids).Error
}

func (s *UserRepo) Update(user map[string]interface{}) error {
	var columns []string
	for key := range user {
		columns = append(columns, key)
	}
	if _, ok := user["id"]; !ok {
		// 不存在id
		return errors.New(fmt.Sprintf("missing %s.id", "user"))
	}
	model := domain.User{}
	model.ID = uint(user["id"].(float64))
	return global.DB.Model(&model).Select(columns).Updates(&user).Error
}

func (s *UserRepo) Find(user domain.User) (domain.User, error) {
	db := global.DB.Model(&domain.User{})
	// TODO：条件过滤

	res := db.First(&user)

	return user, res.Error
}

func (s *UserRepo) List(page domain.PageUserSearch) ([]domain.User, error) {
	var (
		userList []domain.User
		err      error
	)
	// db
	db := global.DB.Model(&domain.User{})
	// page
	offset, limit := pagelimit.OffsetLimit(page.PageNum, page.PageSize)

	// TODO：条件过滤

	err = db.Offset(offset).Limit(limit).Find(&userList).Error

	return userList, err
}

func (s *UserRepo) Count() (int64, error) {
	var (
		count int64
		err   error
	)

	err = global.DB.Model(&domain.User{}).Count(&count).Error

	return count, err
}
