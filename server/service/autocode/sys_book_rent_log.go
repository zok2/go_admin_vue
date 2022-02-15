package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type SysBookRentLogService struct {
}

// CreateSysBookRentLog 创建SysBookRentLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysBookRentLogService *SysBookRentLogService) CreateSysBookRentLog(sysBookRentLog autocode.SysBookRentLog) (err error) {
	err = global.GVA_DB.Create(&sysBookRentLog).Error
	return err
}

// DeleteSysBookRentLog 删除SysBookRentLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysBookRentLogService *SysBookRentLogService)DeleteSysBookRentLog(sysBookRentLog autocode.SysBookRentLog) (err error) {
	err = global.GVA_DB.Delete(&sysBookRentLog).Error
	return err
}

// DeleteSysBookRentLogByIds 批量删除SysBookRentLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysBookRentLogService *SysBookRentLogService)DeleteSysBookRentLogByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.SysBookRentLog{},"id in ?",ids.Ids).Error
	return err
}

// UpdateSysBookRentLog 更新SysBookRentLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysBookRentLogService *SysBookRentLogService)UpdateSysBookRentLog(sysBookRentLog autocode.SysBookRentLog) (err error) {
	err = global.GVA_DB.Save(&sysBookRentLog).Error
	return err
}

// GetSysBookRentLog 根据id获取SysBookRentLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysBookRentLogService *SysBookRentLogService)GetSysBookRentLog(id uint) (err error, sysBookRentLog autocode.SysBookRentLog) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysBookRentLog).Error
	return
}

// GetSysBookRentLogInfoList 分页获取SysBookRentLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysBookRentLogService *SysBookRentLogService)GetSysBookRentLogInfoList(info autoCodeReq.SysBookRentLogSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.SysBookRentLog{})
    var sysBookRentLogs []autocode.SysBookRentLog
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Type != nil {
        db = db.Where("type = ?",info.Type)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&sysBookRentLogs).Error
	return err, sysBookRentLogs, total
}
