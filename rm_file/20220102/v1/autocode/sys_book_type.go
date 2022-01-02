package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type SysBookTypeApi struct {
}

var sysBookTypeService = service.ServiceGroupApp.AutoCodeServiceGroup.SysBookTypeService


// CreateSysBookType 创建SysBookType
// @Tags SysBookType
// @Summary 创建SysBookType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.SysBookType true "创建SysBookType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysBookType/createSysBookType [post]
func (sysBookTypeApi *SysBookTypeApi) CreateSysBookType(c *gin.Context) {
	var sysBookType autocode.SysBookType
	_ = c.ShouldBindJSON(&sysBookType)
	if err := sysBookTypeService.CreateSysBookType(sysBookType); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSysBookType 删除SysBookType
// @Tags SysBookType
// @Summary 删除SysBookType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.SysBookType true "删除SysBookType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysBookType/deleteSysBookType [delete]
func (sysBookTypeApi *SysBookTypeApi) DeleteSysBookType(c *gin.Context) {
	var sysBookType autocode.SysBookType
	_ = c.ShouldBindJSON(&sysBookType)
	if err := sysBookTypeService.DeleteSysBookType(sysBookType); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSysBookTypeByIds 批量删除SysBookType
// @Tags SysBookType
// @Summary 批量删除SysBookType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SysBookType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sysBookType/deleteSysBookTypeByIds [delete]
func (sysBookTypeApi *SysBookTypeApi) DeleteSysBookTypeByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := sysBookTypeService.DeleteSysBookTypeByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSysBookType 更新SysBookType
// @Tags SysBookType
// @Summary 更新SysBookType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.SysBookType true "更新SysBookType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysBookType/updateSysBookType [put]
func (sysBookTypeApi *SysBookTypeApi) UpdateSysBookType(c *gin.Context) {
	var sysBookType autocode.SysBookType
	_ = c.ShouldBindJSON(&sysBookType)
	if err := sysBookTypeService.UpdateSysBookType(sysBookType); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSysBookType 用id查询SysBookType
// @Tags SysBookType
// @Summary 用id查询SysBookType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.SysBookType true "用id查询SysBookType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysBookType/findSysBookType [get]
func (sysBookTypeApi *SysBookTypeApi) FindSysBookType(c *gin.Context) {
	var sysBookType autocode.SysBookType
	_ = c.ShouldBindQuery(&sysBookType)
	if err, resysBookType := sysBookTypeService.GetSysBookType(sysBookType.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resysBookType": resysBookType}, c)
	}
}

// GetSysBookTypeList 分页获取SysBookType列表
// @Tags SysBookType
// @Summary 分页获取SysBookType列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.SysBookTypeSearch true "分页获取SysBookType列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysBookType/getSysBookTypeList [get]
func (sysBookTypeApi *SysBookTypeApi) GetSysBookTypeList(c *gin.Context) {
	var pageInfo autocodeReq.SysBookTypeSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sysBookTypeService.GetSysBookTypeInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
