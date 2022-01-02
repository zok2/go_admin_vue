import service from '@/utils/request'

// @Tags SysBooks
// @Summary 创建SysBooks
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysBooks true "创建SysBooks"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysBooks/createSysBooks [post]
export const createSysBooks = (data) => {
  return service({
    url: '/sysBooks/createSysBooks',
    method: 'post',
    data
  })
}

// @Tags SysBooks
// @Summary 删除SysBooks
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysBooks true "删除SysBooks"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysBooks/deleteSysBooks [delete]
export const deleteSysBooks = (data) => {
  return service({
    url: '/sysBooks/deleteSysBooks',
    method: 'delete',
    data
  })
}

// @Tags SysBooks
// @Summary 删除SysBooks
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SysBooks"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysBooks/deleteSysBooks [delete]
export const deleteSysBooksByIds = (data) => {
  return service({
    url: '/sysBooks/deleteSysBooksByIds',
    method: 'delete',
    data
  })
}

// @Tags SysBooks
// @Summary 更新SysBooks
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysBooks true "更新SysBooks"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysBooks/updateSysBooks [put]
export const updateSysBooks = (data) => {
  return service({
    url: '/sysBooks/updateSysBooks',
    method: 'put',
    data
  })
}

// @Tags SysBooks
// @Summary 用id查询SysBooks
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SysBooks true "用id查询SysBooks"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysBooks/findSysBooks [get]
export const findSysBooks = (params) => {
  return service({
    url: '/sysBooks/findSysBooks',
    method: 'get',
    params
  })
}

// @Tags SysBooks
// @Summary 分页获取SysBooks列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取SysBooks列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysBooks/getSysBooksList [get]
export const getSysBooksList = (params) => {
  return service({
    url: '/sysBooks/getSysBooksList',
    method: 'get',
    params
  })
}
