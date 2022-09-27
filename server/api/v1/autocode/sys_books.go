package autocode

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"time"
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


func (sysBooksApi *SysBooksApi) ImportExcel(c *gin.Context) {
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	var excelInfo example.ExcelBook
	_ = c.ShouldBindJSON(&excelInfo)

	now := time.Now()            //获取当前时间
	timestamp1 := strconv.FormatInt(now.Unix(),10)     //时间戳

	_ = c.SaveUploadedFile(header, global.GVA_CONFIG.Excel.Dir+timestamp1+".xlsx")

	BookInfoList ,err := excelService.ParseBookList2Excel(timestamp1+".xlsx")
	for _,book := range BookInfoList{
		_ = sysBooksService.CreateSysBooks(book)
	}
	if err != nil {
		global.GVA_LOG.Error("转换Excel失败!", zap.Error(err))
		response.FailWithMessage("转换Excel失败", c)
		return
	}

	c.Writer.Header().Add("success", "true")
	response.OkWithMessage("导入成功", c)
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
	sysStock := autocode.SysStock{Status: &Status, BookId: &BookId}
	if err, total := sysStockService.GetSysStockInfoTotal(sysStock); err != nil || total != 0 {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败,有未归还书籍", c)
	} else {
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
	for _, value := range IDS.Ids {
		Status := 2
		BookId := value
		sysStock := autocode.SysStock{Status: &Status, BookId: &BookId}
		if err, total := sysStockService.GetSysStockInfoTotal(sysStock); err != nil || total != 0 {
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
	if err, resysBooks := sysBooksService.GetSysBooks(SysBookFindSearch.ID, SysBookFindSearch.StockStatus); err != nil {
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
	var sysStockBorrowed autocodeReq.SysStockBorrowedStruct
	_ = c.ShouldBindJSON(&sysStockBorrowed)

	Remark := sysStockBorrowed.Remark
	Return_at := ""
	message := "更新成功"
	nowTime := time.Now()
	day := sysStockBorrowed.ReturnAt
	if sysStockBorrowed.Type == 1 {
		beforeTime := nowTime.AddDate(0, 0, day)
		beforeTimeS := beforeTime.Unix() // 秒时间戳
		Return_at = time.Unix(beforeTimeS, 0).Format("2006-01-02") // 固定格式的日期时间戳
	}else {
		_,Stock := sysStockService.GetSysStock(uint(sysStockBorrowed.StockId))
		beforeTime, _ := time.Parse("2006-01-02",Stock.ReturnAt)
		beforeTimeS := beforeTime.Unix() // 秒时间戳
		if beforeTimeS > nowTime.Unix() {
			yun := Stock.Day * 0.5
			message = fmt.Sprintf("未超出借阅时间，正常收费%f元",yun)
		}else {
			yun := Stock.Day * 0.75
			message = fmt.Sprintf("超出借阅时间，收费%f元",yun)
		}
	}

	if err := sysStockService.ChangeStatus(sysStockBorrowed.StockId, sysStockBorrowed.Status,Remark,Return_at,day);err != nil {
			global.GVA_LOG.Error("更新失败!", zap.Error(err))
			response.FailWithMessage("更新失败", c)

	} else {
		SysUserID := int(utils.GetUserID(c))
		log := autocode.SysBookRentLog{BookId: &sysStockBorrowed.BookId, StockId: &sysStockBorrowed.StockId,Type: &sysStockBorrowed.Type, Remark: sysStockBorrowed.Remark, UserId: &sysStockBorrowed.UserId, CreatorId: &SysUserID}
		_ = sysBookRentLogService.CreateSysBookRentLog(log)
		response.OkWithMessage(message, c)
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
