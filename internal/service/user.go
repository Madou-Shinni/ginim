package service

import (
	"errors"
	"github.com/Madou-Shinni/gin-quickstart/internal/conf"
	"github.com/Madou-Shinni/gin-quickstart/internal/data"
	"github.com/Madou-Shinni/gin-quickstart/internal/domain"
	"github.com/Madou-Shinni/gin-quickstart/pkg/global"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
	"github.com/Madou-Shinni/gin-quickstart/pkg/response"
	"github.com/Madou-Shinni/gin-quickstart/pkg/tools"
	"github.com/Madou-Shinni/go-logger"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"strconv"
	"time"
)

var (
	ErrorLoginByGithub = errors.New("github第三方登录失败")
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

func (s *UserService) LoginByGithub(code string) (userInfo *domain.User, token string, err error) {
	clientId := conf.Conf.Github.ClientId
	clientSecret := conf.Conf.Github.ClientSecret
	result, err := tools.LoginGithub(tools.LoginGithubReq{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		Code:         code,
	})
	if err != nil {
		logger.Error("tools.LoginGithub(tools.LoginGithubReq{})", zap.Error(err))
		return nil, "", ErrorLoginByGithub
	}

	info, err := result.GetUserInfo()
	if err != nil {
		logger.Error("result.GetUserInfo()", zap.Error(err))
		return nil, "", ErrorLoginByGithub
	}

	err = global.DB.Where("github_id = ?", info.ID).Attrs(domain.User{
		Name:     info.Name,
		Avatar:   info.Avatar,
		GithubId: info.ID,
	}).FirstOrCreate(&userInfo).Error
	if err != nil {
		return nil, "", ErrorLoginByGithub
	}

	// 生成token
	token, err = tools.GenToken(jwt.MapClaims{
		tools.UserIdKey: strconv.FormatUint(uint64(userInfo.ID), 10),
		tools.ExpKey:    time.Now().AddDate(0, 0, 30).Unix(),
	}, conf.Conf.JwtConfig.Secret)
	if err != nil {
		logger.Error("tools.GenToken(jwt.MapClaims{})", zap.Error(err))
		return nil, "", ErrorLoginByGithub
	}

	return userInfo, token, nil
}
