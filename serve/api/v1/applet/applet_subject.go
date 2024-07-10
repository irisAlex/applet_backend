package applet

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/applet"
	appleteq "github.com/flipped-aurora/gin-vue-admin/server/model/applet/request"
	appletep "github.com/flipped-aurora/gin-vue-admin/server/model/applet/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type SubjectApi struct{}

func (s *SubjectApi) CreateSubjectApi(c *gin.Context) {
	var Subject applet.Subject
	err := c.ShouldBindJSON(&Subject)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(Subject, utils.ApiVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = SubjectService.CreateApi(Subject)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func (s *SubjectApi) GetSubjectApiList(c *gin.Context) {
	var pageInfo appleteq.SearchSubjectParams
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := SubjectService.GetAPIInfoList(pageInfo.Subject, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

func (s *SubjectApi) GetSubjectById(c *gin.Context) {
	var idInfo request.GetById
	err := c.ShouldBindJSON(&idInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(idInfo, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	api, err := SubjectService.GetSubjectById(idInfo.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(appletep.SubjectResponse{Applet: api}, "获取成功", c)
}

func (s *SubjectApi) GetSubjectByEId(c *gin.Context) {
	var idInfo request.GetById
	err := c.ShouldBindJSON(&idInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(idInfo, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	api, err := SubjectService.GetSubjectByEID(idInfo.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(appletep.SubjectResponse{Applet: api}, "获取成功", c)
}

func (s *SubjectApi) GetSubjectByEName(c *gin.Context) {
	var idInfo request.GetByName
	err := c.ShouldBindJSON(&idInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(idInfo, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	api, err := SubjectService.GetSubjectByEName(idInfo.Name)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(appletep.SubjectResponse{Applet: api}, "获取成功", c)
}

func (s *SubjectApi) DeleteSubject(c *gin.Context) {
	var api applet.Subject
	err := c.ShouldBindJSON(&api)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(api.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = SubjectService.DeleteSubject(api)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (s *SubjectApi) UpdateSubject(c *gin.Context) {
	var api applet.Subject
	err := c.ShouldBindJSON(&api)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(api, utils.ApiVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = SubjectService.UpdateSubject(api)
	if err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}
