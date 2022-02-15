package autocode

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"strconv"
)

type SysBooksService struct {
}

// CreateSysBooks 创建SysBooks记录

func (sysBooksService *SysBooksService) CreateSysBooks(sysBooks autocode.SysBooks) (err error) {
	num :=  sysBooks.Amount
	err = global.GVA_DB.Create(&sysBooks).Error

	var stock []autocode.SysStock
	Status := 3
	CreatorId := 1
	BookId := int(sysBooks.ID)
	for i := 0; i < *num; i++ {
		no := 10000+i
		stock = append(stock, autocode.SysStock{
			StockNo: fmt.Sprint(sysBooks.Upc,strconv.Itoa(no)),
			BookId: &BookId,
			Status: &Status,
			CreatorId: &CreatorId,
		})
	}
	_ = global.GVA_DB.Create(&stock).Error
	return err
}

// DeleteSysBooks 删除SysBooks记录

func (sysBooksService *SysBooksService)DeleteSysBooks(sysBooks autocode.SysBooks) (err error) {
	err = global.GVA_DB.Delete(&sysBooks).Error
	return err
}

// DeleteSysBooksByIds 批量删除SysBooks记录

func (sysBooksService *SysBooksService)DeleteSysBooksByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.SysBooks{},"id in ?",ids.Ids).Error
	if err == nil {
		_ = global.GVA_DB.Delete(&[]autocode.SysStock{}, "book_id in ?", ids.Ids).Error
	}
	return err
}

// UpdateSysBooks 更新SysBooks记录

func (sysBooksService *SysBooksService)UpdateSysBooks(sysBooks autocode.SysBooks) (err error) {
	err = global.GVA_DB.Save(&sysBooks).Error
	return err
}

// GetSysBooks 根据id获取SysBooks记录

func (sysBooksService *SysBooksService)GetSysBooks(id uint,stockStatus int) (err error, sysBooks autocode.SysBooks) {
	db := global.GVA_DB.Model(&autocode.SysBooks{})
	if stockStatus != 999 {
		db.Preload("Stock","status= ?",stockStatus)
	}
	err = db.Where("id = ?", id).Preload("Type").First(&sysBooks).Error
	return
}

// GetSysBooksInfoList 分页获取SysBooks记录

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
    if info.Upc != "" {
        db = db.Where("upc = ?",info.Upc)
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
	err = db.Limit(limit).Preload("Type").Preload("Stock","Status=3").Offset(offset).Find(&sysBookss).Error
	return err, sysBookss, total
}
