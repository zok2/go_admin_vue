package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type SysBookTypeService struct {
}

// CreateSysBookType 创建SysBookType记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysBookTypeService *SysBookTypeService) CreateSysBookType(sysBookType autocode.SysBookType) (err error) {
	err = global.GVA_DB.Create(&sysBookType).Error
	return err
}

// DeleteSysBookType 删除SysBookType记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysBookTypeService *SysBookTypeService)DeleteSysBookType(sysBookType autocode.SysBookType) (err error) {
	err = global.GVA_DB.Delete(&sysBookType).Error
	return err
}

// DeleteSysBookTypeByIds 批量删除SysBookType记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysBookTypeService *SysBookTypeService)DeleteSysBookTypeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.SysBookType{},"id in ?",ids.Ids).Error
	return err
}

// UpdateSysBookType 更新SysBookType记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysBookTypeService *SysBookTypeService)UpdateSysBookType(sysBookType autocode.SysBookType) (err error) {
	err = global.GVA_DB.Save(&sysBookType).Error
	return err
}

// GetSysBookType 根据id获取SysBookType记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysBookTypeService *SysBookTypeService)GetSysBookType(id uint) (err error, sysBookType autocode.SysBookType) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysBookType).Error
	return
}

// GetSysBookTypeInfoList 分页获取SysBookType记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysBookTypeService *SysBookTypeService)GetSysBookTypeInfoList(info autoCodeReq.SysBookTypeSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.SysBookType{})
    var sysBookTypes []autocode.SysBookType
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.Pid != nil {
        db = db.Where("pid = ?",info.Pid)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&sysBookTypes).Error
	return err, sysBookTypes, total
}
