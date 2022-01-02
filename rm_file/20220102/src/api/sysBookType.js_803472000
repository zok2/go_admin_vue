import service from '@/utils/request'

// @Tags SysBookType
// @Summary 创建SysBookType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysBookType true "创建SysBookType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysBookType/createSysBookType [post]
export const createSysBookType = (data) => {
  return service({
    url: '/sysBookType/createSysBookType',
    method: 'post',
    data
  })
}

// @Tags SysBookType
// @Summary 删除SysBookType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysBookType true "删除SysBookType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysBookType/deleteSysBookType [delete]
export const deleteSysBookType = (data) => {
  return service({
    url: '/sysBookType/deleteSysBookType',
    method: 'delete',
    data
  })
}

// @Tags SysBookType
// @Summary 删除SysBookType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SysBookType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysBookType/deleteSysBookType [delete]
export const deleteSysBookTypeByIds = (data) => {
  return service({
    url: '/sysBookType/deleteSysBookTypeByIds',
    method: 'delete',
    data
  })
}

// @Tags SysBookType
// @Summary 更新SysBookType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysBookType true "更新SysBookType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysBookType/updateSysBookType [put]
export const updateSysBookType = (data) => {
  return service({
    url: '/sysBookType/updateSysBookType',
    method: 'put',
    data
  })
}

// @Tags SysBookType
// @Summary 用id查询SysBookType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SysBookType true "用id查询SysBookType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysBookType/findSysBookType [get]
export const findSysBookType = (params) => {
  return service({
    url: '/sysBookType/findSysBookType',
    method: 'get',
    params
  })
}

// @Tags SysBookType
// @Summary 分页获取SysBookType列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取SysBookType列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysBookType/getSysBookTypeList [get]
export const getSysBookTypeList = (params) => {
  return service({
    url: '/sysBookType/getSysBookTypeList',
    method: 'get',
    params
  })
}
