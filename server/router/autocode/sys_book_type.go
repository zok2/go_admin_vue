package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysBookTypeRouter struct {
}

// InitSysBookTypeRouter 初始化 SysBookType 路由信息
func (s *SysBookTypeRouter) InitSysBookTypeRouter(Router *gin.RouterGroup) {
	sysBookTypeRouter := Router.Group("sysBookType").Use(middleware.OperationRecord())
	sysBookTypeRouterWithoutRecord := Router.Group("sysBookType")
	var sysBookTypeApi = v1.ApiGroupApp.AutoCodeApiGroup.SysBookTypeApi
	{
		sysBookTypeRouter.POST("createSysBookType", sysBookTypeApi.CreateSysBookType)   // 新建SysBookType
		sysBookTypeRouter.DELETE("deleteSysBookType", sysBookTypeApi.DeleteSysBookType) // 删除SysBookType
		sysBookTypeRouter.DELETE("deleteSysBookTypeByIds", sysBookTypeApi.DeleteSysBookTypeByIds) // 批量删除SysBookType
		sysBookTypeRouter.PUT("updateSysBookType", sysBookTypeApi.UpdateSysBookType)    // 更新SysBookType
	}
	{
		sysBookTypeRouterWithoutRecord.GET("findSysBookType", sysBookTypeApi.FindSysBookType)        // 根据ID获取SysBookType
		sysBookTypeRouterWithoutRecord.GET("getSysBookTypeList", sysBookTypeApi.GetSysBookTypeList)  // 获取SysBookType列表
	}
}
