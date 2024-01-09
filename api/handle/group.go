package handle

import (
	"github.com/Madou-Shinni/gin-quickstart/internal/domain"
	"github.com/Madou-Shinni/gin-quickstart/internal/service"
	"github.com/Madou-Shinni/gin-quickstart/pkg/constant"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
	"github.com/Madou-Shinni/gin-quickstart/pkg/response"
	"github.com/gin-gonic/gin"
)

type GroupHandle struct {
	s *service.GroupService
}

func NewGroupHandle() *GroupHandle {
	return &GroupHandle{s: service.NewGroupService()}
}

// Add 创建Group
// @Tags     Group
// @Summary  创建Group
// @accept   application/json
// @Produce  application/json
// @Param    data body     domain.Group true "创建Group"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /group [post]
func (cl *GroupHandle) Add(c *gin.Context) {
	var group domain.Group
	if err := c.ShouldBindJSON(&group); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if err := cl.s.Add(group); err != nil {
		response.Error(c, constant.CODE_ADD_FAILED, constant.CODE_ADD_FAILED.Msg())
		return
	}

	response.Success(c)
}

// Delete 删除Group
// @Tags     Group
// @Summary  删除Group
// @accept   application/json
// @Produce  application/json
// @Param    data body     domain.Group true "删除Group"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /group [delete]
func (cl *GroupHandle) Delete(c *gin.Context) {
	var group domain.Group
	if err := c.ShouldBindJSON(&group); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if err := cl.s.Delete(group); err != nil {
		response.Error(c, constant.CODE_DELETE_FAILED, constant.CODE_DELETE_FAILED.Msg())
		return
	}

	response.Success(c)
}

// DeleteByIds 批量删除Group
// @Tags     Group
// @Summary  批量删除Group
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.Ids true "批量删除Group"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /group/delete-batch [delete]
func (cl *GroupHandle) DeleteByIds(c *gin.Context) {
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

// Update 修改Group
// @Tags     Group
// @Summary  修改Group
// @accept   application/json
// @Produce  application/json
// @Param    data body     domain.Group true "修改Group"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /group [put]
func (cl *GroupHandle) Update(c *gin.Context) {
	var group map[string]interface{}
	if err := c.ShouldBindJSON(&group); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if err := cl.s.Update(group); err != nil {
		response.Error(c, constant.CODE_UPDATE_FAILED, constant.CODE_UPDATE_FAILED.Msg())
		return
	}

	response.Success(c)
}

// Find 查询Group
// @Tags     Group
// @Summary  查询Group
// @accept   application/json
// @Produce  application/json
// @Param    data query     domain.Group true "查询Group"
// @Success  200  {string} string            "{"code":200,"msg":"查询成功","data":{}"}"
// @Router   /group [get]
func (cl *GroupHandle) Find(c *gin.Context) {
	var group domain.Group
	if err := c.ShouldBindQuery(&group); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	res, err := cl.s.Find(group)

	if err != nil {
		response.Error(c, constant.CODE_FIND_FAILED, constant.CODE_FIND_FAILED.Msg())
		return
	}

	response.Success(c, res)
}

// List 查询Group列表
// @Tags     Group
// @Summary  查询Group列表
// @accept   application/json
// @Produce  application/json
// @Param    data query     domain.Group true "查询Group列表"
// @Success  200  {string} string            "{"code":200,"msg":"查询成功","data":{}"}"
// @Router   /group/list [get]
func (cl *GroupHandle) List(c *gin.Context) {
	var group domain.PageGroupSearch
	if err := c.ShouldBindQuery(&group); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	res, err := cl.s.List(group)

	if err != nil {
		response.Error(c, constant.CODE_FIND_FAILED, constant.CODE_FIND_FAILED.Msg())
		return
	}

	response.Success(c, res)
}