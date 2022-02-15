package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysStockRouter struct {
}

// InitSysStockRouter 初始化 SysStock 路由信息
func (s *SysStockRouter) InitSysStockRouter(Router *gin.RouterGroup) {
	sysStockRouter := Router.Group("sysStock").Use(middleware.OperationRecord())
	sysStockRouterWithoutRecord := Router.Group("sysStock")
	var sysStockApi = v1.ApiGroupApp.AutoCodeApiGroup.SysStockApi
	{
		sysStockRouter.POST("createSysStock", sysStockApi.CreateSysStock)   // 新建SysStock
		sysStockRouter.DELETE("deleteSysStock", sysStockApi.DeleteSysStock) // 删除SysStock
		sysStockRouter.DELETE("deleteSysStockByIds", sysStockApi.DeleteSysStockByIds) // 批量删除SysStock
		sysStockRouter.PUT("updateSysStock", sysStockApi.UpdateSysStock)    // 更新SysStock
	}
	{
		sysStockRouterWithoutRecord.GET("findSysStock", sysStockApi.FindSysStock)        // 根据ID获取SysStock
		sysStockRouterWithoutRecord.GET("getSysStockList", sysStockApi.GetSysStockList)  // 获取SysStock列表
	}
}
