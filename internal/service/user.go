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
type UserRepo interface {
	Create(user domain.User) error
	Delete(user domain.User) error
	Update(user map[string]interface{}) error
	Find(user domain.User) (domain.User, error)
	List(page domain.PageUserSearch) ([]domain.User, error)
	Count() (int64, error)
	DeleteByIds(ids request.Ids) error
}

type UserService struct {
	repo UserRepo
}

func NewUserService() *UserService {
	return &UserService{repo: &data.UserRepo{}}
}

func (s *UserService) Add(user domain.User) error {
	// 3.持久化入库
	if err := s.repo.Create(user); err != nil {
		// 4.记录日志
		logger.Error("s.repo.Create(user)", zap.Error(err), zap.Any("domain.User", user))
		return err
	}

	return nil
}

func (s *UserService) Delete(user domain.User) error {
	if err := s.repo.Delete(user); err != nil {
		logger.Error("s.repo.Delete(user)", zap.Error(err), zap.Any("domain.User", user))
		return err
	}

	return nil
}

func (s *UserService) Update(user map[string]interface{}) error {
	if err := s.repo.Update(user); err != nil {
		logger.Error("s.repo.Update(user)", zap.Error(err), zap.Any("domain.User", user))
		return err
	}

	return nil
}

func (s *UserService) Find(user domain.User) (domain.User, error) {
	res, err := s.repo.Find(user)

	if err != nil {
		logger.Error("s.repo.Find(user)", zap.Error(err), zap.Any("domain.User", user))
		return res, err
	}

	return res, nil
}

func (s *UserService) List(page domain.PageUserSearch) (response.PageResponse, error) {
	var (
		pageRes response.PageResponse
	)

	data, err := s.repo.List(page)
	if err != nil {
		logger.Error("s.repo.List(page)", zap.Error(err), zap.Any("domain.PageUserSearch", page))
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

func (s *UserService) DeleteByIds(ids request.Ids) error {
	if err := s.repo.DeleteByIds(ids); err != nil {
		logger.Error("s.DeleteByIds(ids)", zap.Error(err), zap.Any("ids request.Ids", ids))
		return err
	}

	return nil
}
