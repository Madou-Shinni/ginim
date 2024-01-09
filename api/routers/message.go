package routers

import (
	"github.com/Madou-Shinni/gin-quickstart/api/handle"
	"github.com/gin-gonic/gin"
)

// 注册路由
func MessageRouterRegister(r *gin.Engine) {
	messageGroup := r.Group("message")
	messageHandle := handle.NewMessageHandle()
	{
		messageGroup.POST("", messageHandle.Add)
		messageGroup.DELETE("", messageHandle.Delete)
		messageGroup.DELETE("/delete-batch", messageHandle.DeleteByIds)
		messageGroup.GET("", messageHandle.Find)
		messageGroup.GET("/list", messageHandle.List)
		messageGroup.PUT("", messageHandle.Update)
	}
}