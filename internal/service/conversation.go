package service

import (
	"github.com/Madou-Shinni/gin-quickstart/internal/data"
	"github.com/Madou-Shinni/gin-quickstart/internal/domain"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
	"github.com/Madou-Shinni/gin-quickstart/pkg/response"
	"github.com/Madou-Shinni/go-logger"
	"go.uber.org/zap"
)

// 定义接口
type ConversationRepo interface {
	Create(conversation domain.Conversation) error
	Delete(conversation domain.Conversation) error
	Update(conversation map[string]interface{}) error
	Find(conversation domain.Conversation) (domain.Conversation, error)
	List(page domain.PageConversationSearch) ([]domain.Conversation, error)
	Count() (int64, error)
	DeleteByIds(ids request.Ids) error
}

type ConversationService struct {
	repo ConversationRepo
}

func NewConversationService() *ConversationService {
	return &ConversationService{repo: &data.ConversationRepo{}}
}

func (s *ConversationService) Add(conversation domain.Conversation) error {
	// 3.持久化入库
	if err := s.repo.Create(conversation); err != nil {
		// 4.记录日志
		logger.Error("s.repo.Create(conversation)", zap.Error(err), zap.Any("domain.Conversation", conversation))
		return err
	}

	return nil
}

func (s *ConversationService) Delete(conversation domain.Conversation) error {
	if err := s.repo.Delete(conversation); err != nil {
		logger.Error("s.repo.Delete(conversation)", zap.Error(err), zap.Any("domain.Conversation", conversation))
		return err
	}

	return nil
}

func (s *ConversationService) Update(conversation map[string]interface{}) error {
	if err := s.repo.Update(conversation); err != nil {
		logger.Error("s.repo.Update(conversation)", zap.Error(err), zap.Any("domain.Conversation", conversation))
		return err
	}

	return nil
}

func (s *ConversationService) Find(conversation domain.Conversation) (domain.Conversation, error) {
	res, err := s.repo.Find(conversation)

	if err != nil {
		logger.Error("s.repo.Find(conversation)", zap.Error(err), zap.Any("domain.Conversation", conversation))
		return res, err
	}

	return res, nil
}

func (s *ConversationService) List(page domain.PageConversationSearch) (response.PageResponse, error) {
	var (
		pageRes response.PageResponse
	)

	data, err := s.repo.List(page)
	if err != nil {
		logger.Error("s.repo.List(page)", zap.Error(err), zap.Any("domain.PageConversationSearch", page))
		return pageRes, err
	}

	count, err := s.repo.Count()
	if err != nil {
		logger.Error("s.repo.Count()", zap.Error(err))
		return pageRes, err
	}

	pageRes.Data = data
	pageRes.Total = count

	return pageRes, nil
}

func (s *ConversationService) DeleteByIds(ids request.Ids) error {
	if err := s.repo.DeleteByIds(ids); err != nil {
		logger.Error("s.DeleteByIds(ids)", zap.Error(err), zap.Any("ids request.Ids", ids))
		return err
	}

	return nil
}
