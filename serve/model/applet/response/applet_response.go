package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/applet"

type SubjectResponse struct {
	Applet interface{} `json:"subject"`
}

type PostResponse struct {
	Applet applet.Post `json:"post"`
}

type FavResponse struct {
	Applet applet.Fav `json:"fav"`
}

type FavArrResponse struct {
	Applet []applet.Post `json:"favp"`
}

type IsFavResponse struct {
	IsFav int `json:"is_fav"`
}

type PostMResponse struct {
	Applet []applet.Post `json:"post"`
}

type PPResponse struct {
	Applet []applet.PreviousPostYear `json:"pp"`
}

type PsResponse struct {
	Applet []applet.ProvinceStatistic `json:"ps"`
}
