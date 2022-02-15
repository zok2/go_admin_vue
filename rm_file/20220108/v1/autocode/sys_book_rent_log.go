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

type SysBookRentLogApi struct {
}

var sysBookRentLogService = service.ServiceGroupApp.AutoCodeServiceGroup.SysBookRentLogService


// CreateSysBookRentLog 创建SysBookRentLog
// @Tags SysBookRentLog
// @Summary 创建SysBookRentLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.SysBookRentLog true "创建SysBookRentLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysBookRentLog/createSysBookRentLog [post]
func (sysBookRentLogApi *SysBookRentLogApi) CreateSysBookRentLog(c *gin.Context) {
	var sysBookRentLog autocode.SysBookRentLog
	_ = c.ShouldBindJSON(&sysBookRentLog)
	sysStockService.GetSysStock(sysBookRentLog.)
	if err := sysBookRentLogService.CreateSysBookRentLog(sysBookRentLog); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSysBookRentLog 删除SysBookRentLog
// @Tags SysBookRentLog
// @Summary 删除SysBookRentLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.SysBookRentLog true "删除SysBookRentLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysBookRentLog/deleteSysBookRentLog [delete]
func (sysBookRentLogApi *SysBookRentLogApi) DeleteSysBookRentLog(c *gin.Context) {
	var sysBookRentLog autocode.SysBookRentLog
	_ = c.ShouldBindJSON(&sysBookRentLog)
	if err := sysBookRentLogService.DeleteSysBookRentLog(sysBookRentLog); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSysBookRentLogByIds 批量删除SysBookRentLog
// @Tags SysBookRentLog
// @Summary 批量删除SysBookRentLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SysBookRentLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sysBookRentLog/deleteSysBookRentLogByIds [delete]
func (sysBookRentLogApi *SysBookRentLogApi) DeleteSysBookRentLogByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := sysBookRentLogService.DeleteSysBookRentLogByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSysBookRentLog 更新SysBookRentLog
// @Tags SysBookRentLog
// @Summary 更新SysBookRentLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.SysBookRentLog true "更新SysBookRentLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysBookRentLog/updateSysBookRentLog [put]
func (sysBookRentLogApi *SysBookRentLogApi) UpdateSysBookRentLog(c *gin.Context) {
	var sysBookRentLog autocode.SysBookRentLog
	_ = c.ShouldBindJSON(&sysBookRentLog)
	if err := sysBookRentLogService.UpdateSysBookRentLog(sysBookRentLog); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSysBookRentLog 用id查询SysBookRentLog
// @Tags SysBookRentLog
// @Summary 用id查询SysBookRentLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.SysBookRentLog true "用id查询SysBookRentLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysBookRentLog/findSysBookRentLog [get]
func (sysBookRentLogApi *SysBookRentLogApi) FindSysBookRentLog(c *gin.Context) {
	var sysBookRentLog autocode.SysBookRentLog
	_ = c.ShouldBindQuery(&sysBookRentLog)
	if err, resysBookRentLog := sysBookRentLogService.GetSysBookRentLog(sysBookRentLog.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resysBookRentLog": resysBookRentLog}, c)
	}
}

// GetSysBookRentLogList 分页获取SysBookRentLog列表
// @Tags SysBookRentLog
// @Summary 分页获取SysBookRentLog列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.SysBookRentLogSearch true "分页获取SysBookRentLog列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysBookRentLog/getSysBookRentLogList [get]
func (sysBookRentLogApi *SysBookRentLogApi) GetSysBookRentLogList(c *gin.Context) {
	var pageInfo autocodeReq.SysBookRentLogSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sysBookRentLogService.GetSysBookRentLogInfoList(pageInfo); err != nil {
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
