import service from '@/utils/request'

// @Tags SysBookRentLog
// @Summary 创建SysBookRentLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysBookRentLog true "创建SysBookRentLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysBookRentLog/createSysBookRentLog [post]
export const createSysBookRentLog = (data) => {
  return service({
    url: '/sysBookRentLog/createSysBookRentLog',
    method: 'post',
    data
  })
}

// @Tags SysBookRentLog
// @Summary 删除SysBookRentLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysBookRentLog true "删除SysBookRentLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysBookRentLog/deleteSysBookRentLog [delete]
export const deleteSysBookRentLog = (data) => {
  return service({
    url: '/sysBookRentLog/deleteSysBookRentLog',
    method: 'delete',
    data
  })
}

// @Tags SysBookRentLog
// @Summary 删除SysBookRentLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SysBookRentLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysBookRentLog/deleteSysBookRentLog [delete]
export const deleteSysBookRentLogByIds = (data) => {
  return service({
    url: '/sysBookRentLog/deleteSysBookRentLogByIds',
    method: 'delete',
    data
  })
}

// @Tags SysBookRentLog
// @Summary 更新SysBookRentLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysBookRentLog true "更新SysBookRentLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysBookRentLog/updateSysBookRentLog [put]
export const updateSysBookRentLog = (data) => {
  return service({
    url: '/sysBookRentLog/updateSysBookRentLog',
    method: 'put',
    data
  })
}

// @Tags SysBookRentLog
// @Summary 用id查询SysBookRentLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SysBookRentLog true "用id查询SysBookRentLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysBookRentLog/findSysBookRentLog [get]
export const findSysBookRentLog = (params) => {
  return service({
    url: '/sysBookRentLog/findSysBookRentLog',
    method: 'get',
    params
  })
}

// @Tags SysBookRentLog
// @Summary 分页获取SysBookRentLog列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取SysBookRentLog列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysBookRentLog/getSysBookRentLogList [get]
export const getSysBookRentLogList = (params) => {
  return service({
    url: '/sysBookRentLog/getSysBookRentLogList',
    method: 'get',
    params
  })
}
