package data

import (
	"errors"
	"fmt"
	"github.com/Madou-Shinni/gin-quickstart/constants"
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
		groupIds         []uint
	)
	// db
	db := global.DB
	// page
	offset, limit := pagelimit.OffsetLimit(page.PageNum, page.PageSize)

	if page.OwnerId != 0 {
		// 查询所在的群ids
		err = db.Model(&domain.Relationship{}).
			Where("target = ?", page.OwnerId).
			Where("type = ?", constants.RelationshipTypeGroup).
			Pluck("target", &groupIds).Error
		if err != nil {
			return nil, err
		}
		// 查询私人会话和群会话
		db = db.Where("owner_id = ?", page.OwnerId) // 私人会话
		if len(groupIds) > 0 {
			db = db.Or( // 或者 群会话
				db.Where("type = ?", constants.ConversationTypeGroup).
					Where("owner_id in ?", groupIds),
			)
		}
	}

	// TODO：条件过滤

	err = db.Debug().Offset(offset).Limit(limit).Order("updated_at DESC").Preload("LastMessage").Find(&conversationList).Error

	// 查询会话名称，将用户name或者群name做会话名称
	var userIds []uint
	for _, v := range conversationList {
		if v.Type == constants.ConversationTypePrivate {
			userIds = append(userIds, v.TargetId)
		}
	}

	db = global.DB
	// 查询用户和群
	users := make([]domain.User, 0, len(userIds))
	groups := make([]domain.Group, 0, len(groupIds))
	err = db.Model(&domain.User{}).Select("id", "name").Where("id in ?", userIds).Find(&users).Error
	if err != nil {
		return nil, err
	}
	err = db.Model(&domain.Group{}).Select("id", "name").Where("id in ?", groupIds).Find(&groups).Error
	if err != nil {
		return nil, err
	}

	// 转换为map
	userMap := make(map[uint]domain.User, len(users))
	groupMap := make(map[uint]domain.Group, len(users))
	for _, user := range users {
		userMap[user.ID] = user
	}
	for _, group := range groups {
		groupMap[group.ID] = group
	}

	// 设置会话名称
	for i, v := range conversationList {
		if v.Type == constants.ConversationTypePrivate {
			conversationList[i].Name = userMap[v.TargetId].Name
		}
		if v.Type == constants.ConversationTypeGroup {
			conversationList[i].Name = groupMap[v.OwnerId].Name
		}
	}

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
