package applet

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	SubjectApi
	PostApi
	FavApi
}

var (
	SubjectService = service.ServiceGroupApp.AppletServiceGroup.SubjectApiService
	PostService    = service.ServiceGroupApp.AppletServiceGroup.PostService
	FavService     = service.ServiceGroupApp.AppletServiceGroup.FavService
)
