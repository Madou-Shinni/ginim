package data

import (
    "errors"
    "fmt"
	"github.com/Madou-Shinni/gin-quickstart/internal/domain"
	"github.com/Madou-Shinni/gin-quickstart/pkg/global"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
	"github.com/Madou-Shinni/gin-quickstart/pkg/tools/pagelimit"
)

type GroupRepo struct {
}

func (s *GroupRepo) Create(group domain.Group) error {
	return global.DB.Create(&group).Error
}

func (s *GroupRepo) Delete(group domain.Group) error {
	return global.DB.Delete(&group).Error
}

func (s *GroupRepo) DeleteByIds(ids request.Ids) error {
	return global.DB.Delete(&[]domain.Group{}, ids.Ids).Error
}

func (s *GroupRepo) Update(group map[string]interface{}) error {
    var columns []string
	for key := range group {
		columns = append(columns, key)
	}
	if _,ok := group["id"];!ok {
        // 不存在id
        return errors.New(fmt.Sprintf("missing %s.id","group"))
    }
	model := domain.Group{}
	model.ID = uint(group["id"].(float64))
	return global.DB.Model(&model).Select(columns).Updates(&group).Error
}

func (s *GroupRepo) Find(group domain.Group) (domain.Group, error) {
	db := global.DB.Model(&domain.Group{})
	// TODO：条件过滤

	res := db.First(&group)

	return group, res.Error
}

func (s *GroupRepo) List(page domain.PageGroupSearch) ([]domain.Group, error) {
	var (
		groupList []domain.Group
		err      error
	)
	// db
	db := global.DB.Model(&domain.Group{})
	// page
	offset, limit := pagelimit.OffsetLimit(page.PageNum, page.PageSize)

	// TODO：条件过滤

	err = db.Offset(offset).Limit(limit).Find(&groupList).Error

	return groupList, err
}

func (s *GroupRepo) Count() (int64, error) {
	var (
		count int64
		err   error
	)

	err = global.DB.Model(&domain.Group{}).Count(&count).Error

	return count, err
}
