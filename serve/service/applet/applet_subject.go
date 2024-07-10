package applet

import (
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/applet"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateApi
//@description: 新增基础api
//@param: api model.SysApi
//@return: err error

type SubjectApiService struct{}

func (apiService *SubjectApiService) CreateApi(Subject applet.Subject) (err error) {
	return global.GVA_DB.Create(&Subject).Error
}

func (apiService *SubjectApiService) DeleteSubject(Subject applet.Subject) (err error) {
	var entity applet.Subject
	err = global.GVA_DB.Where("id = ?", Subject.ID).First(&entity).Error // 根据id查询api记录
	if errors.Is(err, gorm.ErrRecordNotFound) {                          // api记录不存在
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

func (apiService *SubjectApiService) GetAPIInfoList(api applet.Subject, info request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&applet.Subject{})
	var apiList []applet.Subject

	// if api.Addr != "" {
	// 	db = db.Where("addr LIKE ?", api.Addr)
	// }

	// if api.Contacts != "" {
	// 	db = db.Where("contacts LIKE ?", api.Contacts)
	// }

	// if api.Email != "" {
	// 	db = db.Where("email LIKE ?", api.Email)
	// }

	// if api.Phone != "" {
	// 	db = db.Where("phone LIKE ?", api.Phone)
	// }

	// if api.Product != "" {
	// 	db = db.Where("product LIKE ?", api.Product)
	// }

	// if api.Name != "" {
	// 	db = db.Where("name = ?", api.Name)
	// }

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
			orderMap["country"] = true
			orderMap["genre"] = true
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
			err = db.Order("id").Find(&apiList).Error
		}
	}
	return apiList, total, err
}

// //@author: [piexlmax](https://github.com/piexlmax)
// //@function: GetAllApis
// //@description: 获取所有的api
// //@return:  apis []model.SysApi, err error

// func (apiService *SubjectApiService) GetAllApis() (apis []system.SysApi, err error) {
// 	err = global.GVA_DB.Find(&apis).Error
// 	return
// }

// //@author: [piexlmax](https://github.com/piexlmax)
// //@function: GetApiById
// //@description: 根据id获取api
// //@param: id float64
// //@return: api model.SysApi, err error

func (apiService *SubjectApiService) GetSubjectById(id int) (api applet.Subject, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&api).Error
	return
}

// //@author: [piexlmax](https://github.com/piexlmax)
// //@function: UpdateApi
// //@description: 根据id更新api
// //@param: api model.SysApi
// //@return: err error

func (apiService *SubjectApiService) UpdateSubject(api applet.Subject) (err error) {
	err = global.GVA_DB.Save(&api).Error
	if err != nil {
		return err
	}
	return nil
}

func (apiService *SubjectApiService) GetSubjectByEID(id int) (api []applet.Subject, err error) {
	err = global.GVA_DB.Where("parent_id=? ", id).Find(&api).Error
	return
}

func (apiService *SubjectApiService) GetSubjectByEName(name string) (api []applet.Subject, err error) {
	err = global.GVA_DB.Where("parent_id=? and education_level_name=?", 0, name).Find(&api).Error
	return
}
