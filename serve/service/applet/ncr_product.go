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

type ProductApiService struct{}

func (apiService *ProductApiService) CreateApi(supplier applet.Product) (err error) {
	return global.GVA_DB.Create(&supplier).Error
}

func (apiService *ProductApiService) DeleteProduct(supplier applet.Product) (err error) {
	var entity applet.Product
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

func (apiService *ProductApiService) GetAPIInfoList(api applet.Product, info request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&applet.Product{})
	var apiList []applet.Product

	if api.Name != "" {
		db = db.Where("name LIKE ?", api.Name)
	}

	if api.Serialnumber != "" {
		db = db.Where("serialnumber LIKE ?", api.Serialnumber)
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
			orderMap["serialnumber"] = true
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

// func (apiService *ProductApiService) GetAllApis() (apis []system.SysApi, err error) {
// 	err = global.GVA_DB.Find(&apis).Error
// 	return
// }

// //@author: [piexlmax](https://github.com/piexlmax)
// //@function: GetApiById
// //@description: 根据id获取api
// //@param: id float64
// //@return: api model.SysApi, err error

func (apiService *ProductApiService) GetProductById(id int) (api applet.Product, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&api).Error
	return
}

// //@author: [piexlmax](https://github.com/piexlmax)
// //@function: UpdateApi
// //@description: 根据id更新api
// //@param: api model.SysApi
// //@return: err error

func (apiService *ProductApiService) UpdateProduct(api applet.Product) (err error) {
	err = global.GVA_DB.Save(&api).Error
	if err != nil {
		return err
	}
	return nil
}

// //@author: [piexlmax](https://github.com/piexlmax)
// //@function: DeleteApis
// //@description: 删除选中API
// //@param: apis []model.SysApi
// //@return: err error

// func (apiService *ProductApiService) DeleteApisByIds(ids request.IdsReq) (err error) {
// 	var apis []system.SysApi
// 	err = global.GVA_DB.Find(&apis, "id in ?", ids.Ids).Delete(&apis).Error
// 	if err != nil {
// 		return err
// 	} else {
// 		for _, sysApi := range apis {
// 			CasbinServiceApp.ClearCasbin(1, sysApi.Path, sysApi.Method)
// 		}
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return err
// }
