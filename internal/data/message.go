package data

import (
    "errors"
    "fmt"
	"github.com/Madou-Shinni/gin-quickstart/internal/domain"
	"github.com/Madou-Shinni/gin-quickstart/pkg/global"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
	"github.com/Madou-Shinni/gin-quickstart/pkg/tools/pagelimit"
)

type MessageRepo struct {
}

func (s *MessageRepo) Create(message domain.Message) error {
	return global.DB.Create(&message).Error
}

func (s *MessageRepo) Delete(message domain.Message) error {
	return global.DB.Delete(&message).Error
}

func (s *MessageRepo) DeleteByIds(ids request.Ids) error {
	return global.DB.Delete(&[]domain.Message{}, ids.Ids).Error
}

func (s *MessageRepo) Update(message map[string]interface{}) error {
    var columns []string
	for key := range message {
		columns = append(columns, key)
	}
	if _,ok := message["id"];!ok {
        // 不存在id
        return errors.New(fmt.Sprintf("missing %s.id","message"))
    }
	model := domain.Message{}
	model.ID = uint(message["id"].(float64))
	return global.DB.Model(&model).Select(columns).Updates(&message).Error
}

func (s *MessageRepo) Find(message domain.Message) (domain.Message, error) {
	db := global.DB.Model(&domain.Message{})
	// TODO：条件过滤

	res := db.First(&message)

	return message, res.Error
}

func (s *MessageRepo) List(page domain.PageMessageSearch) ([]domain.Message, error) {
	var (
		messageList []domain.Message
		err      error
	)
	// db
	db := global.DB.Model(&domain.Message{})
	// page
	offset, limit := pagelimit.OffsetLimit(page.PageNum, page.PageSize)

	// TODO：条件过滤

	err = db.Offset(offset).Limit(limit).Find(&messageList).Error

	return messageList, err
}

func (s *MessageRepo) Count() (int64, error) {
	var (
		count int64
		err   error
	)

	err = global.DB.Model(&domain.Message{}).Count(&count).Error

	return count, err
}
