package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/applet"

type SubjectResponse struct {
	Applet applet.Subject `json:"subject"`
}

type PostResponse struct {
	Applet applet.Post `json:"post"`
}
