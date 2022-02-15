package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type SysStockService struct {
}

// CreateSysStock 创建SysStock记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysStockService *SysStockService) CreateSysStock(sysStock autocode.SysStock) (err error) {
	err = global.GVA_DB.Create(&sysStock).Error
	return err
}

// DeleteSysStock 删除SysStock记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysStockService *SysStockService)DeleteSysStock(sysStock autocode.SysStock) (err error) {
	err = global.GVA_DB.Delete(&sysStock).Error
	return err
}

// DeleteSysStockByIds 批量删除SysStock记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysStockService *SysStockService)DeleteSysStockByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.SysStock{},"id in ?",ids.Ids).Error
	return err
}

// UpdateSysStock 更新SysStock记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysStockService *SysStockService)UpdateSysStock(sysStock autocode.SysStock) (err error) {
	err = global.GVA_DB.Save(&sysStock).Error
	return err
}

// GetSysStock 根据id获取SysStock记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysStockService *SysStockService)GetSysStock(id uint) (err error, sysStock autocode.SysStock) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysStock).Error
	return
}

// GetSysStockInfoList 分页获取SysStock记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysStockService *SysStockService)GetSysStockInfoList(info autoCodeReq.SysStockSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.SysStock{})
    var sysStocks []autocode.SysStock
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StockNo != "" {
        db = db.Where("stock_no LIKE ?","%"+ info.StockNo+"%")
    }
    if info.BookId != nil {
        db = db.Where("book_id = ?",info.BookId)
    }
    if info.Status != nil {
        db = db.Where("status = ?",info.Status)
    }
    if info.UserId != nil {
        db = db.Where("user_id = ?",info.UserId)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&sysStocks).Error
	return err, sysStocks, total
}
