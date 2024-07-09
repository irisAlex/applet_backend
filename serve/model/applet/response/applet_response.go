package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/applet"

type SupplierResponse struct {
	Applet applet.Supplier `json:"supplier"`
}

type TypeResponse struct {
	Applet applet.TypeM `json:"typem"`
}

type ProjectResponse struct {
	Applet applet.Project `json:"project"`
}

type ManageResponse struct {
	Applet applet.Manage `json:"manage"`
}
type ComplaintResponse struct {
	Applet applet.Complaint `json:"compliant"`
}

type MessageResponse struct {
	Applet applet.Message `json:"message"`
}

type ProductResponse struct {
	Applet applet.Product `json:"product"`
}
