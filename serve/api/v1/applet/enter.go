package applet

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	SupplierApi
	TypeApi
	ProjectApi
	ManageApi
	ComplaintApi
	MessageApi
	ProductApi
}

var (
	SupplierService  = service.ServiceGroupApp.AppletServiceGroup.SupplierApiService
	TypeService      = service.ServiceGroupApp.AppletServiceGroup.TypeApiService
	ProjectService   = service.ServiceGroupApp.AppletServiceGroup.ProjectService
	ManageService    = service.ServiceGroupApp.AppletServiceGroup.ManageService
	ComplaintService = service.ServiceGroupApp.AppletServiceGroup.ComplaintService
	MessageService   = service.ServiceGroupApp.AppletServiceGroup.MessageService
	ProductService   = service.ServiceGroupApp.AppletServiceGroup.ProductApiService
)
