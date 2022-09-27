package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysBooksRouter struct {
}

// InitSysBooksRouter 初始化 SysBooks 路由信息
func (s *SysBooksRouter) InitSysBooksRouter(Router *gin.RouterGroup) {
	sysBooksRouter := Router.Group("sysBooks").Use(middleware.OperationRecord())
	sysBooksRouterWithoutRecord := Router.Group("sysBooks")
	var sysBooksApi = v1.ApiGroupApp.AutoCodeApiGroup.SysBooksApi
	{
		sysBooksRouter.POST("createSysBooks", sysBooksApi.CreateSysBooks)             // 新建SysBooks
		sysBooksRouter.POST("importExcel", sysBooksApi.ImportExcel)                // 新建SysBooks
		sysBooksRouter.POST("borrowedSysBooks", sysBooksApi.BorrowedSysBooks)         // 新建SysBooks
		sysBooksRouter.DELETE("deleteSysBooks", sysBooksApi.DeleteSysBooks)           // 删除SysBooks
		sysBooksRouter.DELETE("deleteSysBooksByIds", sysBooksApi.DeleteSysBooksByIds) // 批量删除SysBooks
		sysBooksRouter.PUT("updateSysBooks", sysBooksApi.UpdateSysBooks)              // 更新SysBooks
	}
	{
		sysBooksRouterWithoutRecord.GET("findSysBooks", sysBooksApi.FindSysBooks)        // 根据ID获取SysBooks
		sysBooksRouterWithoutRecord.GET("getSysBooksList", sysBooksApi.GetSysBooksList)  // 获取SysBooks列表
	}
}
