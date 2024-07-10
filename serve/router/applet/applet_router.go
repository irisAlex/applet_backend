package applet

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

// 供应商 router
type SubjectRouter struct{}

func (s *SubjectRouter) InitSubjectRouter(Router *gin.RouterGroup) {
	SubjectApiRouter := Router.Group("Subject").Use(middleware.OperationRecord())
	SubjectApiRouterWithoutRecord := Router.Group("Subject")

	SubjectRouterApi := v1.ApiGroupApp.AppletApiGroup.SubjectApi
	{
		SubjectApiRouter.POST("createSubjectApi", SubjectRouterApi.CreateSubjectApi) // 创建Api
		SubjectApiRouter.POST("deleteSubject", SubjectRouterApi.DeleteSubject)       // 删除Api
		SubjectApiRouter.POST("getSubjectById", SubjectRouterApi.GetSubjectById)     // 获取单条Api消息
		SubjectApiRouter.POST("updateSubject", SubjectRouterApi.UpdateSubject)       // 更新api
		// SubjectApiRouter.DELETE("deleteApisByIds", SubjectRouterApi.DeleteApisByIds) // 删除选中api
	}
	{
		SubjectApiRouterWithoutRecord.POST("getAllSubjectList", SubjectRouterApi.GetSubjectApiList) // 获取所有api
		//SubjectApiRouterWithoutRecord.POST("getSubjectList", SubjectRouterApi.GetSubjectList) // 获取Api列表
	}
}

//项目

type PostRouter struct{}

func (s *PostRouter) InitPostRouter(Router *gin.RouterGroup) {
	typeApiRouter := Router.Group("Post").Use(middleware.OperationRecord())
	typeApiRouterWithoutRecord := Router.Group("Post")

	typeRouterApi := v1.ApiGroupApp.AppletApiGroup.PostApi
	{
		typeApiRouter.POST("createPost", typeRouterApi.CreatePost)   // 创建Api
		typeApiRouter.POST("deletePost", typeRouterApi.DeletePost)   // 删除Api
		typeApiRouter.POST("getPostById", typeRouterApi.GetPostById) // 获取单条Api消息
		typeApiRouter.POST("updatePost", typeRouterApi.UpdatePost)   // 更新api
		// SubjectApiRouter.DELETE("deleteApisByIds", SubjectRouterApi.DeleteApisByIds) // 删除选中api
	}
	{
		typeApiRouterWithoutRecord.POST("getAllPostList", typeRouterApi.GetPostList) // 获取所有api
	}
}
