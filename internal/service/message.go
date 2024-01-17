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
type MessageRepo interface {
	Create(message domain.Message) error
	Delete(message domain.Message) error
	Update(message map[string]interface{}) error
	Find(message domain.Message) (domain.Message, error)
	List(page domain.PageMessageSearch) ([]domain.Message, error)
	Count() (int64, error)
	DeleteByIds(ids request.Ids) error
}

type MessageService struct {
	repo MessageRepo
}

func NewMessageService() *MessageService {
	return &MessageService{repo: &data.MessageRepo{}}
}

func (s *MessageService) Add(message domain.Message) error {
	// 3.持久化入库
	if err := s.repo.Create(message); err != nil {
		// 4.记录日志
		logger.Error("s.repo.Create(message)", zap.Error(err), zap.Any("domain.Message", message))
		return err
	}

	return nil
}

func (s *MessageService) Delete(message domain.Message) error {
	if err := s.repo.Delete(message); err != nil {
		logger.Error("s.repo.Delete(message)", zap.Error(err), zap.Any("domain.Message", message))
		return err
	}

	return nil
}

func (s *MessageService) Update(message map[string]interface{}) error {
	if err := s.repo.Update(message); err != nil {
		logger.Error("s.repo.Update(message)", zap.Error(err), zap.Any("domain.Message", message))
		return err
	}

	return nil
}

func (s *MessageService) Find(message domain.Message) (domain.Message, error) {
	res, err := s.repo.Find(message)

	if err != nil {
		logger.Error("s.repo.Find(message)", zap.Error(err), zap.Any("domain.Message", message))
		return res, err
	}

	return res, nil
}

func (s *MessageService) List(page domain.PageMessageSearch) (response.PageResponse, error) {
	var (
		pageRes response.PageResponse
	)

	data, err := s.repo.List(page)
	if err != nil {
		logger.Error("s.repo.List(page)", zap.Error(err), zap.Any("domain.PageMessageSearch", page))
		return pageRes, err
	}

	count, err := s.repo.Count()
	if err != nil {
		logger.Error("s.repo.Count()", zap.Error(err))
		return pageRes, err
	}

	pageRes.List = data
	pageRes.Total = count

	return pageRes, nil
}

func (s *MessageService) DeleteByIds(ids request.Ids) error {
	if err := s.repo.DeleteByIds(ids); err != nil {
		logger.Error("s.DeleteByIds(ids)", zap.Error(err), zap.Any("ids request.Ids", ids))
		return err
	}

	return nil
}
