package applet

import (
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/applet"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type PostService struct{}

func (apiService *PostService) CreatePost(supplier applet.Post) (err error) {
	return global.GVA_DB.Create(&supplier).Error
}

func (apiService *PostService) DeletePost(supplier applet.Post) (err error) {
	var entity applet.Post
	err = global.GVA_DB.Where("id = ?", supplier.ID).First(&entity).Error // 根据id查询api记录
	if errors.Is(err, gorm.ErrRecordNotFound) {                           // api记录不存在
		return err
	}
	err = global.GVA_DB.Delete(&entity).Error
	if err != nil {
		return err
	}
	//CasbinServiceApp.ClearCasbin(1, entity.Path, entity.Method)
	if err != nil {
		return err
	}
	return nil
}

func (apiService *PostService) GetPostInfoList(api applet.Post, info request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&applet.Post{})
	var apiList []applet.Post
	if api.Province_Name != "" {
		db = db.Where(" province_name = ? ", api.Province_Name)
	}
	if api.Source_Category != "" {
		db = db.Where(" source_category = ? ", api.Source_Category)
	}
	if api.Educational_Require != "" {
		db = db.Where(" educational_require = ? ", api.Educational_Require)
	}
	if api.Specialty != "" {
		db = db.Where(" specialty = ? ", api.Specialty)
	}

	if api.Province_Code != "" {
		db = db.Where(" province_code = ? ", api.Province_Code)
	}

	if api.City_Code != "" {
		db = db.Where(" city_code = ? ", api.City_Code)
	}
	if api.District_Code != "" {
		db = db.Where(" district_code = ? ", api.District_Code)
	}

	if api.Enter_Year != "" {
		db = db.Where(" enter_year = ? ", api.Enter_Year)
	}
	if api.Post_Name != "" {
		db = db.Where(" post_name = ? ", api.Post_Name)
	}

	err = db.Count(&total).Error

	if err != nil {
		return apiList, total, err
	} else {
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			var OrderStr string
			// 设置有效排序key 防止sql注入
			// 感谢 Tom4t0 提交漏洞信息
			orderMap := make(map[string]bool, 5)
			orderMap["id"] = true
			orderMap["name"] = true
			if orderMap[order] {
				if desc {
					OrderStr = order + " desc"
				} else {
					OrderStr = order
				}
			} else { // didn't match any order key in `orderMap`
				err = fmt.Errorf("非法的排序字段: %v", order)
				return apiList, total, err
			}

			err = db.Debug().Order(OrderStr).Find(&apiList).Error
		} else {
			err = db.Debug().Order("id").Find(&apiList).Error
		}
	}
	return apiList, total, err
}

func (apiService *PostService) GetPostById(id int) (api applet.Post, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&api).Error
	return
}

func (apiService *PostService) UpdatePost(api applet.Post) (err error) {
	err = global.GVA_DB.Save(&api).Error
	if err != nil {
		return err
	}
	return nil
}

func (apiService *PostService) Previous_Post() (pp []applet.PreviousPostYear, err error) {
	var results []applet.PreviousPostYear
	err = global.GVA_DB.Model(&applet.Post{}).Select("count(1) as post_number , enter_year as year, SUM(enter_number) AS enter_number , SUM(apply_number) as apply_number").
		Group("enter_year").
		Order("enter_year").
		Scan(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (apiService *PostService) ProvinceStatistics() (ps []applet.ProvinceStatistic, err error) {
	var results []applet.ProvinceStatistic
	err = global.GVA_DB.Model(&applet.Post{}).Select("province_name , count(1) as post_number ,  SUM(enter_number) AS enter_total").
		Group("province_name").
		Order("province_name").
		Scan(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (apiService *PostService) CompanyStatistics(pn string) (ps []applet.ProvinceStatistic, err error) {
	var results []applet.ProvinceStatistic
	err = global.GVA_DB.Model(&applet.Post{}).Select("organization_name as province_name , count(1) as post_number ,  SUM(enter_number) AS enter_total").Where("province_name=?", pn).
		Group("organization_name").
		Order("organization_name").
		Scan(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (apiService *PostService) GetPostByMajor(name string) (api []applet.Post, err error) {
	err = global.GVA_DB.Where("specialty = ?", name).Find(&api).Error
	return
}
