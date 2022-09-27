package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type SysStockService struct {
}

// CreateSysStock 创建SysStock记录

func (sysStockService *SysStockService) CreateSysStock(sysStock autocode.SysStock) (err error) {
	err = global.GVA_DB.Create(&sysStock).Error
	return err
}

// DeleteSysStock 删除SysStock记录

func (sysStockService *SysStockService)DeleteSysStock(sysStock autocode.SysStock) (err error) {
	err = global.GVA_DB.Delete(&sysStock).Error
	return err
}

// DeleteSysStockByIds 批量删除SysStock记录

func (sysStockService *SysStockService)DeleteSysStockByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.SysStock{},"id in ?",ids.Ids).Error
	return err
}

// DeleteSysStockByBookIds 批量删除SysStock记录

func (sysStockService *SysStockService)DeleteSysStockByBookIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.SysStock{},"book_id in ?",ids.Ids).Error
	return err
}

// UpdateSysStock 更新SysStock记录

func (sysStockService *SysStockService)UpdateSysStock(sysStock autocode.SysStock) (err error) {
	err = global.GVA_DB.Save(&sysStock).Error
	return err
}

func (sysStockService *SysStockService)ChangeStatus(Id int, status int,remark string, returnAt string,day int) (err error)  {
	var stock autocode.SysStock

	err = global.GVA_DB.Where("id = ? ", Id ).First(&stock).Update("status",status).
		Update("remark",remark).
		Update("day",day).
		Update("return_at", returnAt).Error
	return err
}

// GetSysStock 根据id获取SysStock记录

func (sysStockService *SysStockService)GetSysStock(id uint) (err error, sysStock autocode.SysStock) {
	err = global.GVA_DB.Preload("Book").Preload("Log").Preload("Borrower").Preload("Creator").Where("id = ?", id).First(&sysStock).Error
	return
}

// GetSysStockInfoList 分页获取SysStock记录

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

	err = db.Limit(limit).Preload("Book").Preload("Borrower").Preload("Creator").Offset(offset).Find(&sysStocks).Error
	return err, sysStocks, total
}

func (sysStockService *SysStockService)GetSysStockInfoTotal(info autocode.SysStock) (err error, total int64) {
	// 创建db
	db := global.GVA_DB.Model(&autocode.SysStock{})
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.BookId != nil {
		db = db.Where("book_id = ?",info.BookId)
	}
	if info.Status != nil {
		db = db.Where("status = ?",info.Status)
	}
	err = db.Count(&total).Error
	return err, total
}