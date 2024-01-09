package routers

import (
	"github.com/Madou-Shinni/gin-quickstart/common"
	"github.com/gin-gonic/gin"
)

// 注册路由
func WSRouterRegister(r *gin.Engine) {
	wsGroup := r.Group("ws")
	{
		wsGroup.GET("conn", common.WsHandler)
	}
}
