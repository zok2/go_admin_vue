package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type SysBooksService struct {
}

// CreateSysBooks 创建SysBooks记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysBooksService *SysBooksService) CreateSysBooks(sysBooks autocode.SysBooks) (err error) {
	err = global.GVA_DB.Create(&sysBooks).Error
	return err
}

// DeleteSysBooks 删除SysBooks记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysBooksService *SysBooksService)DeleteSysBooks(sysBooks autocode.SysBooks) (err error) {
	err = global.GVA_DB.Delete(&sysBooks).Error
	return err
}

// DeleteSysBooksByIds 批量删除SysBooks记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysBooksService *SysBooksService)DeleteSysBooksByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.SysBooks{},"id in ?",ids.Ids).Error
	return err
}

// UpdateSysBooks 更新SysBooks记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysBooksService *SysBooksService)UpdateSysBooks(sysBooks autocode.SysBooks) (err error) {
	err = global.GVA_DB.Save(&sysBooks).Error
	return err
}

// GetSysBooks 根据id获取SysBooks记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysBooksService *SysBooksService)GetSysBooks(id uint) (err error, sysBooks autocode.SysBooks) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysBooks).Error
	return
}

// GetSysBooksInfoList 分页获取SysBooks记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysBooksService *SysBooksService)GetSysBooksInfoList(info autoCodeReq.SysBooksSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.SysBooks{})
    var sysBookss []autocode.SysBooks
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.Isbn != "" {
        db = db.Where("isbn = ?",info.Isbn)
    }
    if info.TypeId != nil {
        db = db.Where("type_id = ?",info.TypeId)
    }
    if info.Status != nil {
        db = db.Where("status = ?",info.Status)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&sysBookss).Error
	return err, sysBookss, total
}
