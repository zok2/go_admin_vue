import service from '@/utils/request'

// @Tags SysStock
// @Summary 创建SysStock
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysStock true "创建SysStock"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysStock/createSysStock [post]
export const createSysStock = (data) => {
  return service({
    url: '/sysStock/createSysStock',
    method: 'post',
    data
  })
}

// @Tags SysStock
// @Summary 删除SysStock
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysStock true "删除SysStock"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysStock/deleteSysStock [delete]
export const deleteSysStock = (data) => {
  return service({
    url: '/sysStock/deleteSysStock',
    method: 'delete',
    data
  })
}

// @Tags SysStock
// @Summary 删除SysStock
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SysStock"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysStock/deleteSysStock [delete]
export const deleteSysStockByIds = (data) => {
  return service({
    url: '/sysStock/deleteSysStockByIds',
    method: 'delete',
    data
  })
}

// @Tags SysStock
// @Summary 更新SysStock
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysStock true "更新SysStock"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysStock/updateSysStock [put]
export const updateSysStock = (data) => {
  return service({
    url: '/sysStock/updateSysStock',
    method: 'put',
    data
  })
}

// @Tags SysStock
// @Summary 用id查询SysStock
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SysStock true "用id查询SysStock"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysStock/findSysStock [get]
export const findSysStock = (params) => {
  return service({
    url: '/sysStock/findSysStock',
    method: 'get',
    params
  })
}

// @Tags SysStock
// @Summary 分页获取SysStock列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取SysStock列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysStock/getSysStockList [get]
export const getSysStockList = (params) => {
  return service({
    url: '/sysStock/getSysStockList',
    method: 'get',
    params
  })
}
