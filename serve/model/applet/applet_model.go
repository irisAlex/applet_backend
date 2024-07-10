package applet

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type Post struct {
	global.GVA_MODEL
	Enter_Number        int64  `json:"enter_number" gorm:"comment:类别"`
	Enter_Year          int64  `json:"enter_year" gorm:"comment:类别"`
	Enter_Subject       string `json:"enter_subject" gorm:"comment:类别"`
	Enter_Source        string `json:"enter_source" gorm:"comment:类别"`
	Enter_Ratio         string `json:"enter_ratio" gorm:"comment:类别"`
	Subject             string `json:"subject" gorm:"comment:类别"`
	Post_Name           string `json:"post_name" gorm:"comment:类别"`
	Post_Code           string `json:"post_code" gorm:"comment:类别"`
	Workplace           string `json:"workplace" gorm:"comment:类别"`
	Post_Category       string `json:"post_category" gorm:"comment:类别"`
	Perform_Work        string `json:"perform_work" gorm:"comment:类别"`
	Organization_Name   string `json:"organization_name" gorm:"comment:类别"`
	Organization_Code   string `json:"organization_code" gorm:"comment:类别"`
	Source_Category     string `json:"source_category" gorm:"comment:类别"`
	Educational_Require string `json:"educational_require" gorm:"comment:类别"`
	Degree_Require      string `json:"gedegree_requirenre" gorm:"comment:类别"`
	Career              string `json:"career" gorm:"comment:类别"`
	Title_Require       string `json:"title_require" gorm:"comment:类别"`
	Qualification       string `json:"qualification" gorm:"comment:类别"`
	Other               string `json:"other" gorm:"comment:类别"`
	Specialty           string `json:"specialty" gorm:"comment:类别"`
	Fractional_Line     string `json:"fractional_line" gorm:"comment:类别"`
}

func (Post) TableName() string {
	return "applet_post"
}

type Subject struct {
	global.GVA_MODEL
	Name                 string `json:"name" gorm:"comment:类别"`
	Parent_Id            int64  `json:"parent_id" gorm:"comment:类别"`
	Level_Name           string `json:"level_name" gorm:"comment:类别"`
	Education_Level_Name string `json:"education_level_name" gorm:"comment:类别"`
}

func (Subject) TableName() string {
	return "applet_Subjects"
}
