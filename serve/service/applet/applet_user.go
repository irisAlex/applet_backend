package applet

import (
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/applet"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type UserService struct{}

func (apiService *UserService) CreateUser(fav applet.User) (err error) {
	return global.GVA_DB.Create(&fav).Error
}

func (apiService *UserService) DeleteUser(fav applet.User) (err error) {
	var entity applet.User
	err = global.GVA_DB.Where("id = ?", fav.ID).First(&entity).Error // 根据id查询api记录
	if errors.Is(err, gorm.ErrRecordNotFound) {                      // api记录不存在
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

func (apiService *UserService) GetUserInfoList(api applet.User, info request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&applet.User{})
	var apiList []applet.User
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

			err = db.Order(OrderStr).Find(&apiList).Error
		} else {
			err = db.Debug().Order("id").Find(&apiList).Error
		}
	}
	return apiList, total, err
}

func (apiService *UserService) GetUserById(id int) (api applet.User, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&api).Error
	return
}

func (apiService *UserService) UpdateUser(api applet.User) (err error) {
	err = global.GVA_DB.Save(&api).Error
	if err != nil {
		return err
	}
	return nil
}

func (apiService *UserService) IsUser(id int) (api applet.User, err error) {
	err = global.GVA_DB.Where("pid = ? and is_fav=?", id, true).First(&api).Error
	return
}

func (apiService *UserService) GetUserPost(uid int) (res []applet.Post, err error) {
	err = global.GVA_DB.Raw("SELECT applet_post.* FROM applet_fav LEFT JOIN applet_post ON applet_post.id = applet_fav.pid Where applet_fav.uid=?", uid).Scan(&res).Error
	return
}
