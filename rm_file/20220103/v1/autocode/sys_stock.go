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

type SysStockApi struct {
}

var sysStockService = service.ServiceGroupApp.AutoCodeServiceGroup.SysStockService


// CreateSysStock 创建SysStock
// @Tags SysStock
// @Summary 创建SysStock
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.SysStock true "创建SysStock"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysStock/createSysStock [post]
func (sysStockApi *SysStockApi) CreateSysStock(c *gin.Context) {
	var sysStock autocode.SysStock
	_ = c.ShouldBindJSON(&sysStock)
	if err := sysStockService.CreateSysStock(sysStock); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSysStock 删除SysStock
// @Tags SysStock
// @Summary 删除SysStock
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.SysStock true "删除SysStock"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysStock/deleteSysStock [delete]
func (sysStockApi *SysStockApi) DeleteSysStock(c *gin.Context) {
	var sysStock autocode.SysStock
	_ = c.ShouldBindJSON(&sysStock)
	if err := sysStockService.DeleteSysStock(sysStock); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSysStockByIds 批量删除SysStock
// @Tags SysStock
// @Summary 批量删除SysStock
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SysStock"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sysStock/deleteSysStockByIds [delete]
func (sysStockApi *SysStockApi) DeleteSysStockByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := sysStockService.DeleteSysStockByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSysStock 更新SysStock
// @Tags SysStock
// @Summary 更新SysStock
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.SysStock true "更新SysStock"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysStock/updateSysStock [put]
func (sysStockApi *SysStockApi) UpdateSysStock(c *gin.Context) {
	var sysStock autocode.SysStock
	_ = c.ShouldBindJSON(&sysStock)
	if err := sysStockService.UpdateSysStock(sysStock); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSysStock 用id查询SysStock
// @Tags SysStock
// @Summary 用id查询SysStock
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.SysStock true "用id查询SysStock"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysStock/findSysStock [get]
func (sysStockApi *SysStockApi) FindSysStock(c *gin.Context) {
	var sysStock autocode.SysStock
	_ = c.ShouldBindQuery(&sysStock)
	if err, resysStock := sysStockService.GetSysStock(sysStock.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resysStock": resysStock}, c)
	}
}

// GetSysStockList 分页获取SysStock列表
// @Tags SysStock
// @Summary 分页获取SysStock列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.SysStockSearch true "分页获取SysStock列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysStock/getSysStockList [get]
func (sysStockApi *SysStockApi) GetSysStockList(c *gin.Context) {
	var pageInfo autocodeReq.SysStockSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sysStockService.GetSysStockInfoList(pageInfo); err != nil {
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
