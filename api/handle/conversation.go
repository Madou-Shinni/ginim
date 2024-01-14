package handle

import (
	"github.com/Madou-Shinni/gin-quickstart/internal/domain"
	"github.com/Madou-Shinni/gin-quickstart/internal/service"
	"github.com/Madou-Shinni/gin-quickstart/pkg/constant"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
	"github.com/Madou-Shinni/gin-quickstart/pkg/response"
	"github.com/gin-gonic/gin"
)

type ConversationHandle struct {
	s *service.ConversationService
}

func NewConversationHandle() *ConversationHandle {
	return &ConversationHandle{s: service.NewConversationService()}
}

// Add 创建Conversation
// @Tags     Conversation
// @Summary  创建Conversation
// @accept   application/json
// @Produce  application/json
// @Param    data body     domain.Conversation true "创建Conversation"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /conversation [post]
func (cl *ConversationHandle) Add(c *gin.Context) {
	var conversation domain.Conversation
	if err := c.ShouldBindJSON(&conversation); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if err := cl.s.Add(conversation); err != nil {
		response.Error(c, constant.CODE_ADD_FAILED, constant.CODE_ADD_FAILED.Msg())
		return
	}

	response.Success(c)
}

// Delete 删除Conversation
// @Tags     Conversation
// @Summary  删除Conversation
// @accept   application/json
// @Produce  application/json
// @Param    data body     domain.Conversation true "删除Conversation"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /conversation [delete]
func (cl *ConversationHandle) Delete(c *gin.Context) {
	var conversation domain.Conversation
	if err := c.ShouldBindJSON(&conversation); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if err := cl.s.Delete(conversation); err != nil {
		response.Error(c, constant.CODE_DELETE_FAILED, constant.CODE_DELETE_FAILED.Msg())
		return
	}

	response.Success(c)
}

// DeleteByIds 批量删除Conversation
// @Tags     Conversation
// @Summary  批量删除Conversation
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.Ids true "批量删除Conversation"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /conversation/delete-batch [delete]
func (cl *ConversationHandle) DeleteByIds(c *gin.Context) {
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

// Update 修改Conversation
// @Tags     Conversation
// @Summary  修改Conversation
// @accept   application/json
// @Produce  application/json
// @Param    data body     domain.Conversation true "修改Conversation"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /conversation [put]
func (cl *ConversationHandle) Update(c *gin.Context) {
	var conversation map[string]interface{}
	if err := c.ShouldBindJSON(&conversation); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if err := cl.s.Update(conversation); err != nil {
		response.Error(c, constant.CODE_UPDATE_FAILED, constant.CODE_UPDATE_FAILED.Msg())
		return
	}

	response.Success(c)
}

// Find 查询Conversation
// @Tags     Conversation
// @Summary  查询Conversation
// @accept   application/json
// @Produce  application/json
// @Param    data query     domain.Conversation true "查询Conversation"
// @Success  200  {string} string            "{"code":200,"msg":"查询成功","data":{}"}"
// @Router   /conversation [get]
func (cl *ConversationHandle) Find(c *gin.Context) {
	var conversation domain.Conversation
	if err := c.ShouldBindQuery(&conversation); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	res, err := cl.s.Find(conversation)

	if err != nil {
		response.Error(c, constant.CODE_FIND_FAILED, constant.CODE_FIND_FAILED.Msg())
		return
	}

	response.Success(c, res)
}

// List 查询Conversation列表
// @Tags     Conversation
// @Summary  查询Conversation列表
// @accept   application/json
// @Produce  application/json
// @Param    data query     domain.Conversation true "查询Conversation列表"
// @Success  200  {string} string            "{"code":200,"msg":"查询成功","data":{}"}"
// @Router   /conversation/list [get]
func (cl *ConversationHandle) List(c *gin.Context) {
	var conversation domain.PageConversationSearch
	if err := c.ShouldBindQuery(&conversation); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	res, err := cl.s.List(conversation)

	if err != nil {
		response.Error(c, constant.CODE_FIND_FAILED, constant.CODE_FIND_FAILED.Msg())
		return
	}

	response.Success(c, res)
}
