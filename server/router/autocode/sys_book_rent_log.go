package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysBookRentLogRouter struct {
}

// InitSysBookRentLogRouter 初始化 SysBookRentLog 路由信息
func (s *SysBookRentLogRouter) InitSysBookRentLogRouter(Router *gin.RouterGroup) {
	sysBookRentLogRouter := Router.Group("sysBookRentLog").Use(middleware.OperationRecord())
	sysBookRentLogRouterWithoutRecord := Router.Group("sysBookRentLog")
	var sysBookRentLogApi = v1.ApiGroupApp.AutoCodeApiGroup.SysBookRentLogApi
	{
		sysBookRentLogRouter.POST("createSysBookRentLog", sysBookRentLogApi.CreateSysBookRentLog)   // 新建SysBookRentLog
		sysBookRentLogRouter.DELETE("deleteSysBookRentLog", sysBookRentLogApi.DeleteSysBookRentLog) // 删除SysBookRentLog
		sysBookRentLogRouter.DELETE("deleteSysBookRentLogByIds", sysBookRentLogApi.DeleteSysBookRentLogByIds) // 批量删除SysBookRentLog
		sysBookRentLogRouter.PUT("updateSysBookRentLog", sysBookRentLogApi.UpdateSysBookRentLog)    // 更新SysBookRentLog
	}
	{
		sysBookRentLogRouterWithoutRecord.GET("findSysBookRentLog", sysBookRentLogApi.FindSysBookRentLog)        // 根据ID获取SysBookRentLog
		sysBookRentLogRouterWithoutRecord.GET("getSysBookRentLogList", sysBookRentLogApi.GetSysBookRentLogList)  // 获取SysBookRentLog列表
	}
}
