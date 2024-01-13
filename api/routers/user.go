package routers

import (
	"github.com/Madou-Shinni/gin-quickstart/api/handle"
	"github.com/gin-gonic/gin"
)

// 注册路由
func UserRouterRegister(r *gin.Engine) {
	userGroup := r.Group("user")
	userHandle := handle.NewUserHandle()
	{
		userGroup.POST("", userHandle.Add)
		userGroup.DELETE("", userHandle.Delete)
		userGroup.DELETE("/delete-batch", userHandle.DeleteByIds)
		userGroup.GET("", userHandle.Find)
		userGroup.GET("/list", userHandle.List)
		userGroup.PUT("", userHandle.Update)
		userGroup.GET("oauth2/github/callback", userHandle.LoginByGithub) // github第三方登录
	}
}
