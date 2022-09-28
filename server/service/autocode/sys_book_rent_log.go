package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
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
func (sysBookRentLogService *SysBookRentLogService)GetSysBookRentLog(stockId int) (err error, sysBookRentLog []autocode.SysBookRentLog) {
	err = global.GVA_DB.Where("stockId = ?", stockId).Preload("Book").Preload("Borrower").Preload("Creator").Find(&sysBookRentLog).Error
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
    if info.UserId != nil {
        db = db.Where("UserId = ?",info.UserId)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Preload("Book").Limit(limit).Preload("Borrower").Preload("Creator").Offset(offset).Order("updated_at desc").Find(&sysBookRentLogs).Error
	return err, sysBookRentLogs, total
}
