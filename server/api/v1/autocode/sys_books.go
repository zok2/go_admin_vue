package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
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
	Status := 2
	BookId := int(sysBooks.ID)
	sysStock :=  autocode.SysStock{Status:&Status,BookId:&BookId}
	if err,total := sysStockService.GetSysStockInfoTotal(sysStock); err != nil || total != 0 {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败,有未归还书籍", c)
	}else {
		if err := sysBooksService.DeleteSysBooks(sysBooks); err != nil {
			global.GVA_LOG.Error("删除失败!", zap.Error(err))
			response.FailWithMessage("删除失败", c)
		} else {
			IDS := request.IdsReq{Ids: []int{1}}
			_ = sysStockService.DeleteSysStockByBookIds(IDS)
			response.OkWithMessage("删除成功", c)
		}
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
	for _,value :=range IDS.Ids{
		Status := 2
		BookId := value
		sysStock :=  autocode.SysStock{Status:&Status,BookId:&BookId}
		if err,total := sysStockService.GetSysStockInfoTotal(sysStock); err != nil || total!=0 {
			response.FailWithMessage("批量删除失败,部分书籍中有未归还书籍", c)
			return
		}
	}

   	if err := sysBooksService.DeleteSysBooksByIds(IDS); err != nil {
   		_ = sysStockService.DeleteSysStockByIds(IDS)
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
	var SysBookFindSearch autocodeReq.SysBookFindSearch
	_ = c.ShouldBindQuery(&SysBookFindSearch)
	if err, resysBooks := sysBooksService.GetSysBooks(SysBookFindSearch.ID,SysBookFindSearch.StockStatus); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resysBooks": resysBooks}, c)
	}
}

// BorrowedSysBooks 借书-还书
// @Tags SysBooks
// @Summary 用id查询SysBooks
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.SysBooks true "用id查询SysBooks"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysBooks/BorrowSysBooks [POST]
func (sysBooksApi *SysBooksApi) BorrowedSysBooks(c *gin.Context) {
		var sysBookRentLog autocode.SysBookRentLog
		_ = c.ShouldBindJSON(&sysBookRentLog)
		if err := sysStockService.UpdateSysStock(sysBooks); err != nil {
			global.GVA_LOG.Error("更新失败!", zap.Error(err))
			response.FailWithMessage("更新失败", c)
		} else {
			response.OkWithMessage("更新成功", c)
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
