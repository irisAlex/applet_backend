package applet

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

// 供应商 router
type SubjectRouter struct{}

func (s *SubjectRouter) InitSubjectRouter(Router *gin.RouterGroup) {
	SubjectApiRouter := Router.Group("Subject").Use()
	SubjectApiRouterWithoutRecord := Router.Group("Subject")

	SubjectRouterApi := v1.ApiGroupApp.AppletApiGroup.SubjectApi
	{
		SubjectApiRouter.POST("createSubjectApi", SubjectRouterApi.CreateSubjectApi)   // 创建Api
		SubjectApiRouter.POST("deleteSubject", SubjectRouterApi.DeleteSubject)         // 删除Api
		SubjectApiRouter.POST("getSubjectById", SubjectRouterApi.GetSubjectById)       // 获取单条Api消息
		SubjectApiRouter.POST("updateSubject", SubjectRouterApi.UpdateSubject)         // 更新api
		SubjectApiRouter.POST("getSubjectByEID", SubjectRouterApi.GetSubjectByEId)     // 更新api
		SubjectApiRouter.POST("getSubjectByEName", SubjectRouterApi.GetSubjectByEName) // 更新api

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
	typeApiRouter := Router.Group("Post").Use()
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
		typeApiRouterWithoutRecord.POST("getAllPostList", typeRouterApi.GetPostList)               // 获取所有api
		typeApiRouterWithoutRecord.POST("getPreviousPost", typeRouterApi.Previous_Post)            // 获取所有api
		typeApiRouterWithoutRecord.POST("getProvinceStatistics", typeRouterApi.ProvinceStatistics) // 获取所有api
		typeApiRouterWithoutRecord.POST("getCompanyStatistics", typeRouterApi.CompanyStatistics)   // 获取所有api
		typeApiRouterWithoutRecord.POST("getPostByMajor", typeRouterApi.GetPostByMajor)            // 获取所有api
	}
}

type FavRouter struct{}

func (f *FavRouter) InitFavRouter(Router *gin.RouterGroup) {
	typeApiRouter := Router.Group("Fav").Use()
	typeApiRouterWithoutRecord := Router.Group("Fav")

	typeRouterApi := v1.ApiGroupApp.AppletApiGroup.FavApi
	{
		typeApiRouter.POST("createFav", typeRouterApi.CreateFav)   // 创建Api
		typeApiRouter.POST("deleteFav", typeRouterApi.DeleteFav)   // 删除Api
		typeApiRouter.POST("getFavById", typeRouterApi.GetFavById) // 获取单条Api消息
		typeApiRouter.POST("update", typeRouterApi.UpdateFav)      // 更新api
		typeApiRouter.POST("is_fav", typeRouterApi.IsFav)          // 更新api
		typeApiRouter.POST("getFavPost", typeRouterApi.GetFavPost) // 更新api

		// SubjectApiRouter.DELETE("deleteApisByIds", SubjectRouterApi.DeleteApisByIds) // 删除选中api
	}
	{
		typeApiRouterWithoutRecord.POST("getAllFavList", typeRouterApi.GetFavList) // 获取所有api
		// 获取所有api
	}
}

type UserRouter struct{}

func (f *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	typeApiRouter := Router.Group("User").Use()
	typeApiRouterWithoutRecord := Router.Group("User")

	typeRouterApi := v1.ApiGroupApp.AppletApiGroup.UserApi
	{
		typeApiRouter.POST("createUser", typeRouterApi.CreateUser)   // 创建Api
		typeApiRouter.POST("deleteUser", typeRouterApi.DeleteUser)   // 删除Api
		typeApiRouter.POST("getUserById", typeRouterApi.GetUserById) // 获取单条Api消息
		typeApiRouter.POST("update", typeRouterApi.UpdateUser)       // 更新api
		typeApiRouter.POST("is_fav", typeRouterApi.IsUser)           // 更新api

		// SubjectApiRouter.DELETE("deleteApisByIds", SubjectRouterApi.DeleteApisByIds) // 删除选中api
	}
	{
		typeApiRouterWithoutRecord.POST("getAllUserList", typeRouterApi.GetUserList) // 获取所有api
		// 获取所有api
	}
}
