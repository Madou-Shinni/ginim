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
type RelationshipRepo interface {
	Create(relationship domain.Relationship) error
	Delete(relationship domain.Relationship) error
	Update(relationship map[string]interface{}) error
	Find(relationship domain.Relationship) (domain.Relationship, error)
	List(page domain.PageRelationshipSearch) ([]domain.Relationship, error)
	Count(page domain.PageRelationshipSearch) (int64, error)
	DeleteByIds(ids request.Ids) error
}

type RelationshipService struct {
	repo RelationshipRepo
}

func NewRelationshipService() *RelationshipService {
	return &RelationshipService{repo: &data.RelationshipRepo{}}
}

func (s *RelationshipService) Add(relationship domain.Relationship) error {
	// 3.持久化入库
	if err := s.repo.Create(relationship); err != nil {
		// 4.记录日志
		logger.Error("s.repo.Create(relationship)", zap.Error(err), zap.Any("domain.Relationship", relationship))
		return err
	}

	return nil
}

func (s *RelationshipService) Delete(relationship domain.Relationship) error {
	if err := s.repo.Delete(relationship); err != nil {
		logger.Error("s.repo.Delete(relationship)", zap.Error(err), zap.Any("domain.Relationship", relationship))
		return err
	}

	return nil
}

func (s *RelationshipService) Update(relationship map[string]interface{}) error {
	if err := s.repo.Update(relationship); err != nil {
		logger.Error("s.repo.Update(relationship)", zap.Error(err), zap.Any("domain.Relationship", relationship))
		return err
	}

	return nil
}

func (s *RelationshipService) Find(relationship domain.Relationship) (domain.Relationship, error) {
	res, err := s.repo.Find(relationship)

	if err != nil {
		logger.Error("s.repo.Find(relationship)", zap.Error(err), zap.Any("domain.Relationship", relationship))
		return res, err
	}

	return res, nil
}

func (s *RelationshipService) List(page domain.PageRelationshipSearch) (response.PageResponse, error) {
	var (
		pageRes response.PageResponse
	)

	data, err := s.repo.List(page)
	if err != nil {
		logger.Error("s.repo.List(page)", zap.Error(err), zap.Any("domain.PageRelationshipSearch", page))
		return pageRes, err
	}

	count, err := s.repo.Count(page)
	if err != nil {
		logger.Error("s.repo.Count()", zap.Error(err))
		return pageRes, err
	}

	pageRes.List = data
	pageRes.Total = count

	return pageRes, nil
}

func (s *RelationshipService) DeleteByIds(ids request.Ids) error {
	if err := s.repo.DeleteByIds(ids); err != nil {
		logger.Error("s.DeleteByIds(ids)", zap.Error(err), zap.Any("ids request.Ids", ids))
		return err
	}

	return nil
}
