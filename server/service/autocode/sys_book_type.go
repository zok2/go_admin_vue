package autocode

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type SysBookTypeService struct {
}

// CreateSysBookType 创建SysBookType记录

func (sysBookTypeService *SysBookTypeService) CreateSysBookType(sysBookType autocode.SysBookType) (err error) {
	err = global.GVA_DB.Create(&sysBookType).Error
	return err
}

// DeleteSysBookType 删除SysBookType记录

func (sysBookTypeService *SysBookTypeService)DeleteSysBookType(sysBookType autocode.SysBookType) (err error) {
	if !errors.Is(global.GVA_DB.Where("type_id = ?", sysBookType.ID).First(&autocode.SysBooks{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if !errors.Is(global.GVA_DB.Where("parent_id = ?", sysBookType.ID).First(&autocode.SysBookType{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此分类有子分类正在使用")
	}
	err = global.GVA_DB.Delete(&sysBookType).Error
	return err
}

// DeleteSysBookTypeByIds 批量删除SysBookType记录

func (sysBookTypeService *SysBookTypeService)DeleteSysBookTypeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.SysBookType{},"id in ?",ids.Ids).Error
	return err
}

// UpdateSysBookType 更新SysBookType记录

func (sysBookTypeService *SysBookTypeService)UpdateSysBookType(sysBookType autocode.SysBookType) (err error) {
	err = global.GVA_DB.Save(&sysBookType).Error
	return err
}

// GetSysBookType 根据id获取SysBookType记录

func (sysBookTypeService *SysBookTypeService)GetSysBookType(id uint) (err error, sysBookType autocode.SysBookType) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysBookType).Error
	return
}

// GetSysBookTypeInfoList 分页获取SysBookType记录

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
    db = db.Where("pid = ?",0)

	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&sysBookTypes).Error
	if len(sysBookTypes) > 0 {
		for k := range sysBookTypes {
			err = sysBookTypeService.findChildrenSysBookType(&sysBookTypes[k])
		}
	}
	return err, sysBookTypes, total
}

func (sysBookTypeService *SysBookTypeService) findChildrenSysBookType(sysBookType *autocode.SysBookType) (err error) {
	err = global.GVA_DB.Where("pid = ?", sysBookType.ID).Find(&sysBookType.Children).Error
	if len(sysBookType.Children) > 0 {
		for k := range sysBookType.Children {
			err = sysBookTypeService.findChildrenSysBookType(&sysBookType.Children[k])
		}
	}
	return err
}
