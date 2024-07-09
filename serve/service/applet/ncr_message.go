package applet

import (
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/applet"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type MessageService struct{}

func (apiService *MessageService) CreateMessage(supplier applet.Message) (err error) {
	return global.GVA_DB.Create(&supplier).Error
}

func (apiService *MessageService) DeleteMessage(supplier applet.Message) (err error) {
	var entity applet.Message
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

func (apiService *MessageService) GetMessageInfoList(api applet.Message, info request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&applet.Message{})
	var apiList []applet.Message

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

func (apiService *MessageService) GetMessageByname(name string) (list interface{}, total int64, err error) {
	var apiList []applet.Message
	db := global.GVA_DB.Model(&applet.Message{})
	err = db.Where("message_receive_name = ? and message_is_active = ?", name, true).Find(&apiList).Error
	if err != nil {
		return apiList, total, err
	}

	err = db.Count(&total).Error
	if err != nil {
		return apiList, total, err
	}
	return apiList, total, err
}

func (apiService *MessageService) GetMessageByAppletID(nid int64, name string) (applet.Message, error) {
	var apiList applet.Message
	db := global.GVA_DB.Model(&applet.Message{})
	err := db.Where("applet_id = ? and message_receive_name = ? ", nid, name).First(&apiList).Error
	if err != nil {
		return applet.Message{}, err
	}
	return apiList, nil
}

func (apiService *MessageService) UpdateMessage(api applet.Message) (err error) {
	err = global.GVA_DB.Save(&api).Error
	if err != nil {
		return err
	}
	return nil
}

func (apiService *MessageService) SetMessageState(ID uint) (err error) {
	err = global.GVA_DB.Model(&applet.Message{}).Where("id = ?", ID).Update("message_is_active", false).Error
	return err
}
