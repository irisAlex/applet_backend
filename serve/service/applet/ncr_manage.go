package applet

import (
	"errors"
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/applet"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type ManageService struct{}

func (apiService *ManageService) CreateManage(supplier applet.Manage) (err error) {
	return global.GVA_DB.Create(&supplier).Error
}

func (apiService *ManageService) DeleteManage(supplier applet.Manage) (err error) {
	var entity applet.Manage
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

func (apiService *ManageService) GetManageInfoList(api applet.Manage, info request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&applet.Manage{})
	var apiList []applet.Manage

	if api.Department != "" {
		db = db.Where("department = ?", api.Department)
	}
	if api.Mold != "" {
		db = db.Where("mold = ?", api.Mold)
	}

	if api.Category != "" {
		db = db.Where("category = ?", api.Category)
	}

	if api.Checkout_Name != "" {
		db = db.Where("checkout_name = ?", api.Checkout_Name)
	}
	if api.Checkout_Number != "" {
		db = db.Where("checkout_number = ?", api.Checkout_Number)
	}

	if api.Project != "" {
		db = db.Where("project = ?", api.Project)
	}

	if api.Process_Mode != "" {
		db = db.Where("process_mode = ?", api.Process_Mode)
	}

	if !api.CreatedAt.IsZero() {
		db = db.Where(" created_at = ? ", api.CreatedAt.Format("2006-01-02"))
	}

	if !api.Deferred_Date.IsZero() {
		db = db.Where(" deferred_date = ? ", api.Deferred_Date.Format("2006-01-02"))
	}

	if info.Keyword != "" {
		db = db.Where(" operation_type = ? ", info.Keyword)
		db = db.Where(" is_applet = ? ", true)
	}

	if api.Operation_Type != "" {
		db = db.Where(" operation_type = ? ", api.Operation_Type)
		db = db.Where(" is_applet = ? ", true)
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
			if orderMap[order] {
				if desc {
					OrderStr = order + " desc"
				} else {
					OrderStr = order + "desc"
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

func (apiService *ManageService) GetManageById(id int) (api applet.Manage, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&api).Error
	return
}

func (apiService *ManageService) UpdateManage(api applet.Manage) (err error) {
	err = global.GVA_DB.Save(&api).Error
	if err != nil {
		return err
	}
	return nil
}

func (apiService *ManageService) SetAppletStatus(ID uint, st string) (err error) {
	err = global.GVA_DB.Model(&applet.Manage{}).Where("id = ?", ID).Update("status", st).Error
	return err
}

func (apiService *ManageService) SetAppletPassDate(ID uint, op_type string) (err error) {
	err = global.GVA_DB.Model(&applet.Manage{}).Where("id = ?", ID).Updates(applet.Manage{Operation_Type: op_type, Pass_Date: time.Now(), Is_Applet: true}).Error
	return err
}

func (apiService *ManageService) CloseAllById(ID uint) (err error) {
	err = global.GVA_DB.Model(&applet.Manage{}).Where("id = ?", ID).Updates(applet.Manage{Operation_Type: "-1", Is_Applet: false}).Error
	return err
}

func (apiService *ManageService) UpdateParts(api applet.Manage) (err error) {
	err = global.GVA_DB.Model(&applet.Manage{}).Where("id = ?", api.ID).Updates(applet.Manage{Status: api.Status, Deferred_Date: api.Deferred_Date}).Error
	return err
}

func (apiService *ManageService) SetApplet(ID uint) (err error) {
	err = global.GVA_DB.Model(&applet.Manage{}).Where("id = ?", ID).Update("is_applet", true).Error
	return err
}
