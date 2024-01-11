package handle

import (
	"github.com/Madou-Shinni/gin-quickstart/internal/domain"
	"github.com/Madou-Shinni/gin-quickstart/internal/service"
	"github.com/Madou-Shinni/gin-quickstart/pkg/constant"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
	"github.com/Madou-Shinni/gin-quickstart/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserHandle struct {
	s *service.UserService
}

func NewUserHandle() *UserHandle {
	return &UserHandle{s: service.NewUserService()}
}

// Add 创建User
// @Tags     User
// @Summary  创建User
// @accept   application/json
// @Produce  application/json
// @Param    data body     domain.User true "创建User"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /user [post]
func (cl *UserHandle) Add(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if err := cl.s.Add(user); err != nil {
		response.Error(c, constant.CODE_ADD_FAILED, constant.CODE_ADD_FAILED.Msg())
		return
	}

	response.Success(c)
}

// Delete 删除User
// @Tags     User
// @Summary  删除User
// @accept   application/json
// @Produce  application/json
// @Param    data body     domain.User true "删除User"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /user [delete]
func (cl *UserHandle) Delete(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if err := cl.s.Delete(user); err != nil {
		response.Error(c, constant.CODE_DELETE_FAILED, constant.CODE_DELETE_FAILED.Msg())
		return
	}

	response.Success(c)
}

// DeleteByIds 批量删除User
// @Tags     User
// @Summary  批量删除User
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.Ids true "批量删除User"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /user/delete-batch [delete]
func (cl *UserHandle) DeleteByIds(c *gin.Context) {
	var ids request.Ids
	if err := c.ShouldBindJSON(&ids); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if err := cl.s.DeleteByIds(ids); err != nil {
		response.Error(c, constant.CODE_DELETE_FAILED, constant.CODE_DELETE_FAILED.Msg())
		return
	}

	response.Success(c)
}

// Update 修改User
// @Tags     User
// @Summary  修改User
// @accept   application/json
// @Produce  application/json
// @Param    data body     domain.User true "修改User"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /user [put]
func (cl *UserHandle) Update(c *gin.Context) {
	var user map[string]interface{}
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if err := cl.s.Update(user); err != nil {
		response.Error(c, constant.CODE_UPDATE_FAILED, constant.CODE_UPDATE_FAILED.Msg())
		return
	}

	response.Success(c)
}

// Find 查询User
// @Tags     User
// @Summary  查询User
// @accept   application/json
// @Produce  application/json
// @Param    data query     domain.User true "查询User"
// @Success  200  {string} string            "{"code":200,"msg":"查询成功","data":{}"}"
// @Router   /user [get]
func (cl *UserHandle) Find(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindQuery(&user); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	res, err := cl.s.Find(user)

	if err != nil {
		response.Error(c, constant.CODE_FIND_FAILED, constant.CODE_FIND_FAILED.Msg())
		return
	}

	response.Success(c, res)
}

// List 查询User列表
// @Tags     User
// @Summary  查询User列表
// @accept   application/json
// @Produce  application/json
// @Param    data query     domain.User true "查询User列表"
// @Success  200  {string} string            "{"code":200,"msg":"查询成功","data":{}"}"
// @Router   /user/list [get]
func (cl *UserHandle) List(c *gin.Context) {
	var user domain.PageUserSearch
	if err := c.ShouldBindQuery(&user); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	res, err := cl.s.List(user)

	if err != nil {
		response.Error(c, constant.CODE_FIND_FAILED, constant.CODE_FIND_FAILED.Msg())
		return
	}

	response.Success(c, res)
}
