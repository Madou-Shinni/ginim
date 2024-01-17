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
type GroupRepo interface {
	Create(group domain.Group) error
	Delete(group domain.Group) error
	Update(group map[string]interface{}) error
	Find(group domain.Group) (domain.Group, error)
	List(page domain.PageGroupSearch) ([]domain.Group, error)
	Count() (int64, error)
	DeleteByIds(ids request.Ids) error
}

type GroupService struct {
	repo GroupRepo
}

func NewGroupService() *GroupService {
	return &GroupService{repo: &data.GroupRepo{}}
}

func (s *GroupService) Add(group domain.Group) error {
	// 3.持久化入库
	if err := s.repo.Create(group); err != nil {
		// 4.记录日志
		logger.Error("s.repo.Create(group)", zap.Error(err), zap.Any("domain.Group", group))
		return err
	}

	return nil
}

func (s *GroupService) Delete(group domain.Group) error {
	if err := s.repo.Delete(group); err != nil {
		logger.Error("s.repo.Delete(group)", zap.Error(err), zap.Any("domain.Group", group))
		return err
	}

	return nil
}

func (s *GroupService) Update(group map[string]interface{}) error {
	if err := s.repo.Update(group); err != nil {
		logger.Error("s.repo.Update(group)", zap.Error(err), zap.Any("domain.Group", group))
		return err
	}

	return nil
}

func (s *GroupService) Find(group domain.Group) (domain.Group, error) {
	res, err := s.repo.Find(group)

	if err != nil {
		logger.Error("s.repo.Find(group)", zap.Error(err), zap.Any("domain.Group", group))
		return res, err
	}

	return res, nil
}

func (s *GroupService) List(page domain.PageGroupSearch) (response.PageResponse, error) {
	var (
		pageRes response.PageResponse
	)

	data, err := s.repo.List(page)
	if err != nil {
		logger.Error("s.repo.List(page)", zap.Error(err), zap.Any("domain.PageGroupSearch", page))
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

func (s *GroupService) DeleteByIds(ids request.Ids) error {
	if err := s.repo.DeleteByIds(ids); err != nil {
		logger.Error("s.DeleteByIds(ids)", zap.Error(err), zap.Any("ids request.Ids", ids))
		return err
	}

	return nil
}
