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

type SysBooksApi struct {
}

var sysBooksService = service.ServiceGroupApp.AutoCodeServiceGroup.SysBooksService


// CreateSysBooks 创建SysBooks
// @Tags SysBooks
// @Summary 创建SysBooks
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.SysBooks true "创建SysBooks"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysBooks/createSysBooks [post]
func (sysBooksApi *SysBooksApi) CreateSysBooks(c *gin.Context) {
	var sysBooks autocode.SysBooks
	_ = c.ShouldBindJSON(&sysBooks)
	if err := sysBooksService.CreateSysBooks(sysBooks); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSysBooks 删除SysBooks
// @Tags SysBooks
// @Summary 删除SysBooks
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.SysBooks true "删除SysBooks"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysBooks/deleteSysBooks [delete]
func (sysBooksApi *SysBooksApi) DeleteSysBooks(c *gin.Context) {
	var sysBooks autocode.SysBooks
	_ = c.ShouldBindJSON(&sysBooks)
	if err := sysBooksService.DeleteSysBooks(sysBooks); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSysBooksByIds 批量删除SysBooks
// @Tags SysBooks
// @Summary 批量删除SysBooks
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SysBooks"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sysBooks/deleteSysBooksByIds [delete]
func (sysBooksApi *SysBooksApi) DeleteSysBooksByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := sysBooksService.DeleteSysBooksByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSysBooks 更新SysBooks
// @Tags SysBooks
// @Summary 更新SysBooks
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.SysBooks true "更新SysBooks"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysBooks/updateSysBooks [put]
func (sysBooksApi *SysBooksApi) UpdateSysBooks(c *gin.Context) {
	var sysBooks autocode.SysBooks
	_ = c.ShouldBindJSON(&sysBooks)
	if err := sysBooksService.UpdateSysBooks(sysBooks); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSysBooks 用id查询SysBooks
// @Tags SysBooks
// @Summary 用id查询SysBooks
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.SysBooks true "用id查询SysBooks"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysBooks/findSysBooks [get]
func (sysBooksApi *SysBooksApi) FindSysBooks(c *gin.Context) {
	var sysBooks autocode.SysBooks
	_ = c.ShouldBindQuery(&sysBooks)
	if err, resysBooks := sysBooksService.GetSysBooks(sysBooks.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resysBooks": resysBooks}, c)
	}
}

// GetSysBooksList 分页获取SysBooks列表
// @Tags SysBooks
// @Summary 分页获取SysBooks列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.SysBooksSearch true "分页获取SysBooks列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysBooks/getSysBooksList [get]
func (sysBooksApi *SysBooksApi) GetSysBooksList(c *gin.Context) {
	var pageInfo autocodeReq.SysBooksSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sysBooksService.GetSysBooksInfoList(pageInfo); err != nil {
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
