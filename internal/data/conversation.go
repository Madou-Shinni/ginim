package data

import (
	"errors"
	"fmt"
	"github.com/Madou-Shinni/gin-quickstart/internal/domain"
	"github.com/Madou-Shinni/gin-quickstart/pkg/global"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
	"github.com/Madou-Shinni/gin-quickstart/pkg/tools/pagelimit"
)

type ConversationRepo struct {
}

func (s *ConversationRepo) Create(conversation domain.Conversation) error {
	return global.DB.Create(&conversation).Error
}

func (s *ConversationRepo) Delete(conversation domain.Conversation) error {
	return global.DB.Delete(&conversation).Error
}

func (s *ConversationRepo) DeleteByIds(ids request.Ids) error {
	return global.DB.Delete(&[]domain.Conversation{}, ids.Ids).Error
}

func (s *ConversationRepo) Update(conversation map[string]interface{}) error {
	var columns []string
	for key := range conversation {
		columns = append(columns, key)
	}
	if _, ok := conversation["id"]; !ok {
		// 不存在id
		return errors.New(fmt.Sprintf("missing %s.id", "conversation"))
	}
	model := domain.Conversation{}
	model.ID = uint(conversation["id"].(float64))
	return global.DB.Model(&model).Select(columns).Updates(&conversation).Error
}

func (s *ConversationRepo) Find(conversation domain.Conversation) (domain.Conversation, error) {
	db := global.DB.Model(&domain.Conversation{})
	// TODO：条件过滤

	res := db.First(&conversation)

	return conversation, res.Error
}

func (s *ConversationRepo) List(page domain.PageConversationSearch) ([]domain.Conversation, error) {
	var (
		conversationList []domain.Conversation
		err              error
	)
	// db
	db := global.DB.Model(&domain.Conversation{})
	// page
	offset, limit := pagelimit.OffsetLimit(page.PageNum, page.PageSize)

	// TODO：条件过滤

	err = db.Offset(offset).Limit(limit).Find(&conversationList).Error

	return conversationList, err
}

func (s *ConversationRepo) Count(page domain.PageConversationSearch) (int64, error) {
	var (
		count int64
		err   error
	)

	err = global.DB.Model(&domain.Conversation{}).Count(&count).Error

	return count, err
}
