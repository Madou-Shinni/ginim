package routers

import (
	"github.com/Madou-Shinni/gin-quickstart/api/handle"
	"github.com/gin-gonic/gin"
)

// 注册路由
func ConversationRouterRegister(r *gin.RouterGroup) {
	conversationGroup := r.Group("conversation")
	conversationHandle := handle.NewConversationHandle()
	{
		conversationGroup.POST("", conversationHandle.Add)
		conversationGroup.DELETE("", conversationHandle.Delete)
		conversationGroup.DELETE("/delete-batch", conversationHandle.DeleteByIds)
		conversationGroup.GET("", conversationHandle.Find)
		conversationGroup.GET("/list", conversationHandle.List)
		conversationGroup.PUT("", conversationHandle.Update)
	}
}
