package routers

import (
	"github.com/Madou-Shinni/gin-quickstart/api/handle"
	"github.com/gin-gonic/gin"
)

// 注册路由
func GroupRouterRegister(r *gin.Engine) {
	groupGroup := r.Group("group")
	groupHandle := handle.NewGroupHandle()
	{
		groupGroup.POST("", groupHandle.Add)
		groupGroup.DELETE("", groupHandle.Delete)
		groupGroup.DELETE("/delete-batch", groupHandle.DeleteByIds)
		groupGroup.GET("", groupHandle.Find)
		groupGroup.GET("/list", groupHandle.List)
		groupGroup.PUT("", groupHandle.Update)
	}
}