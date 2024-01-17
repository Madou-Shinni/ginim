package routers

import (
	"github.com/Madou-Shinni/gin-quickstart/api/handle"
	"github.com/gin-gonic/gin"
)

// 注册路由
func RelationshipRouterRegister(r *gin.RouterGroup) {
	relationshipGroup := r.Group("relationship")
	relationshipHandle := handle.NewRelationshipHandle()
	{
		relationshipGroup.POST("", relationshipHandle.Add)
		relationshipGroup.DELETE("", relationshipHandle.Delete)
		relationshipGroup.DELETE("/delete-batch", relationshipHandle.DeleteByIds)
		relationshipGroup.GET("", relationshipHandle.Find)
		relationshipGroup.GET("/list", relationshipHandle.List)
		relationshipGroup.PUT("", relationshipHandle.Update)
	}
}
