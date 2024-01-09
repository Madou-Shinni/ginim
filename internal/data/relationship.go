package data

import (
    "errors"
    "fmt"
	"github.com/Madou-Shinni/gin-quickstart/internal/domain"
	"github.com/Madou-Shinni/gin-quickstart/pkg/global"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
	"github.com/Madou-Shinni/gin-quickstart/pkg/tools/pagelimit"
)

type RelationshipRepo struct {
}

func (s *RelationshipRepo) Create(relationship domain.Relationship) error {
	return global.DB.Create(&relationship).Error
}

func (s *RelationshipRepo) Delete(relationship domain.Relationship) error {
	return global.DB.Delete(&relationship).Error
}

func (s *RelationshipRepo) DeleteByIds(ids request.Ids) error {
	return global.DB.Delete(&[]domain.Relationship{}, ids.Ids).Error
}

func (s *RelationshipRepo) Update(relationship map[string]interface{}) error {
    var columns []string
	for key := range relationship {
		columns = append(columns, key)
	}
	if _,ok := relationship["id"];!ok {
        // 不存在id
        return errors.New(fmt.Sprintf("missing %s.id","relationship"))
    }
	model := domain.Relationship{}
	model.ID = uint(relationship["id"].(float64))
	return global.DB.Model(&model).Select(columns).Updates(&relationship).Error
}

func (s *RelationshipRepo) Find(relationship domain.Relationship) (domain.Relationship, error) {
	db := global.DB.Model(&domain.Relationship{})
	// TODO：条件过滤

	res := db.First(&relationship)

	return relationship, res.Error
}

func (s *RelationshipRepo) List(page domain.PageRelationshipSearch) ([]domain.Relationship, error) {
	var (
		relationshipList []domain.Relationship
		err      error
	)
	// db
	db := global.DB.Model(&domain.Relationship{})
	// page
	offset, limit := pagelimit.OffsetLimit(page.PageNum, page.PageSize)

	// TODO：条件过滤

	err = db.Offset(offset).Limit(limit).Find(&relationshipList).Error

	return relationshipList, err
}

func (s *RelationshipRepo) Count() (int64, error) {
	var (
		count int64
		err   error
	)

	err = global.DB.Model(&domain.Relationship{}).Count(&count).Error

	return count, err
}
