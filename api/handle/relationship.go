package handle

import (
	"github.com/Madou-Shinni/gin-quickstart/common"
	"github.com/Madou-Shinni/gin-quickstart/internal/domain"
	"github.com/Madou-Shinni/gin-quickstart/internal/service"
	"github.com/Madou-Shinni/gin-quickstart/pkg/constant"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
	"github.com/Madou-Shinni/gin-quickstart/pkg/response"
	"github.com/gin-gonic/gin"
)

type RelationshipHandle struct {
	s *service.RelationshipService
}

func NewRelationshipHandle() *RelationshipHandle {
	return &RelationshipHandle{s: service.NewRelationshipService()}
}

// Add 创建Relationship
// @Tags     Relationship
// @Summary  创建Relationship
// @accept   application/json
// @Produce  application/json
// @Param    data body     domain.Relationship true "创建Relationship"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /relationship [post]
func (cl *RelationshipHandle) Add(c *gin.Context) {
	var relationship domain.Relationship
	if err := c.ShouldBindJSON(&relationship); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if err := cl.s.Add(relationship); err != nil {
		response.Error(c, constant.CODE_ADD_FAILED, constant.CODE_ADD_FAILED.Msg())
		return
	}

	response.Success(c)
}

// Delete 删除Relationship
// @Tags     Relationship
// @Summary  删除Relationship
// @accept   application/json
// @Produce  application/json
// @Param    data body     domain.Relationship true "删除Relationship"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /relationship [delete]
func (cl *RelationshipHandle) Delete(c *gin.Context) {
	var relationship domain.Relationship
	if err := c.ShouldBindJSON(&relationship); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if err := cl.s.Delete(relationship); err != nil {
		response.Error(c, constant.CODE_DELETE_FAILED, constant.CODE_DELETE_FAILED.Msg())
		return
	}

	response.Success(c)
}

// DeleteByIds 批量删除Relationship
// @Tags     Relationship
// @Summary  批量删除Relationship
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.Ids true "批量删除Relationship"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /relationship/delete-batch [delete]
func (cl *RelationshipHandle) DeleteByIds(c *gin.Context) {
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

// Update 修改Relationship
// @Tags     Relationship
// @Summary  修改Relationship
// @accept   application/json
// @Produce  application/json
// @Param    data body     domain.Relationship true "修改Relationship"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /relationship [put]
func (cl *RelationshipHandle) Update(c *gin.Context) {
	var relationship map[string]interface{}
	if err := c.ShouldBindJSON(&relationship); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if err := cl.s.Update(relationship); err != nil {
		response.Error(c, constant.CODE_UPDATE_FAILED, constant.CODE_UPDATE_FAILED.Msg())
		return
	}

	response.Success(c)
}

// Find 查询Relationship
// @Tags     Relationship
// @Summary  查询Relationship
// @accept   application/json
// @Produce  application/json
// @Param    data query     domain.Relationship true "查询Relationship"
// @Success  200  {string} string            "{"code":200,"msg":"查询成功","data":{}"}"
// @Router   /relationship [get]
func (cl *RelationshipHandle) Find(c *gin.Context) {
	var relationship domain.Relationship
	if err := c.ShouldBindQuery(&relationship); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	res, err := cl.s.Find(relationship)

	if err != nil {
		response.Error(c, constant.CODE_FIND_FAILED, constant.CODE_FIND_FAILED.Msg())
		return
	}

	response.Success(c, res)
}

// List 查询Relationship列表
// @Tags     Relationship
// @Summary  查询Relationship列表
// @accept   application/json
// @Produce  application/json
// @Param    data query     domain.Relationship true "查询Relationship列表"
// @Success  200  {string} string            "{"code":200,"msg":"查询成功","data":{}"}"
// @Router   /relationship/list [get]
func (cl *RelationshipHandle) List(c *gin.Context) {
	var relationship domain.PageRelationshipSearch
	if err := c.ShouldBindQuery(&relationship); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	userId, err := common.GetUserIdFromCtx(c)
	if err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, err.Error())
		return
	}

	relationship.Owner = userId

	res, err := cl.s.List(relationship)

	if err != nil {
		response.Error(c, constant.CODE_FIND_FAILED, constant.CODE_FIND_FAILED.Msg())
		return
	}

	response.Success(c, res)
}
