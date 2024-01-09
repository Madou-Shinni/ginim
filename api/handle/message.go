package handle

import (
	"github.com/Madou-Shinni/gin-quickstart/internal/domain"
	"github.com/Madou-Shinni/gin-quickstart/internal/service"
	"github.com/Madou-Shinni/gin-quickstart/pkg/constant"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
	"github.com/Madou-Shinni/gin-quickstart/pkg/response"
	"github.com/gin-gonic/gin"
)

type MessageHandle struct {
	s *service.MessageService
}

func NewMessageHandle() *MessageHandle {
	return &MessageHandle{s: service.NewMessageService()}
}

// Add 创建Message
// @Tags     Message
// @Summary  创建Message
// @accept   application/json
// @Produce  application/json
// @Param    data body     domain.Message true "创建Message"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /message [post]
func (cl *MessageHandle) Add(c *gin.Context) {
	var message domain.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if err := cl.s.Add(message); err != nil {
		response.Error(c, constant.CODE_ADD_FAILED, constant.CODE_ADD_FAILED.Msg())
		return
	}

	response.Success(c)
}

// Delete 删除Message
// @Tags     Message
// @Summary  删除Message
// @accept   application/json
// @Produce  application/json
// @Param    data body     domain.Message true "删除Message"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /message [delete]
func (cl *MessageHandle) Delete(c *gin.Context) {
	var message domain.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if err := cl.s.Delete(message); err != nil {
		response.Error(c, constant.CODE_DELETE_FAILED, constant.CODE_DELETE_FAILED.Msg())
		return
	}

	response.Success(c)
}

// DeleteByIds 批量删除Message
// @Tags     Message
// @Summary  批量删除Message
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.Ids true "批量删除Message"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /message/delete-batch [delete]
func (cl *MessageHandle) DeleteByIds(c *gin.Context) {
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

// Update 修改Message
// @Tags     Message
// @Summary  修改Message
// @accept   application/json
// @Produce  application/json
// @Param    data body     domain.Message true "修改Message"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /message [put]
func (cl *MessageHandle) Update(c *gin.Context) {
	var message map[string]interface{}
	if err := c.ShouldBindJSON(&message); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if err := cl.s.Update(message); err != nil {
		response.Error(c, constant.CODE_UPDATE_FAILED, constant.CODE_UPDATE_FAILED.Msg())
		return
	}

	response.Success(c)
}

// Find 查询Message
// @Tags     Message
// @Summary  查询Message
// @accept   application/json
// @Produce  application/json
// @Param    data query     domain.Message true "查询Message"
// @Success  200  {string} string            "{"code":200,"msg":"查询成功","data":{}"}"
// @Router   /message [get]
func (cl *MessageHandle) Find(c *gin.Context) {
	var message domain.Message
	if err := c.ShouldBindQuery(&message); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	res, err := cl.s.Find(message)

	if err != nil {
		response.Error(c, constant.CODE_FIND_FAILED, constant.CODE_FIND_FAILED.Msg())
		return
	}

	response.Success(c, res)
}

// List 查询Message列表
// @Tags     Message
// @Summary  查询Message列表
// @accept   application/json
// @Produce  application/json
// @Param    data query     domain.Message true "查询Message列表"
// @Success  200  {string} string            "{"code":200,"msg":"查询成功","data":{}"}"
// @Router   /message/list [get]
func (cl *MessageHandle) List(c *gin.Context) {
	var message domain.PageMessageSearch
	if err := c.ShouldBindQuery(&message); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	res, err := cl.s.List(message)

	if err != nil {
		response.Error(c, constant.CODE_FIND_FAILED, constant.CODE_FIND_FAILED.Msg())
		return
	}

	response.Success(c, res)
}