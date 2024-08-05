package applet

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	SubjectApi
	PostApi
	FavApi
	UserApi
}

var (
	SubjectService = service.ServiceGroupApp.AppletServiceGroup.SubjectApiService
	PostService    = service.ServiceGroupApp.AppletServiceGroup.PostService
	FavService     = service.ServiceGroupApp.AppletServiceGroup.FavService
	UserService    = service.ServiceGroupApp.AppletServiceGroup.UserService
)
