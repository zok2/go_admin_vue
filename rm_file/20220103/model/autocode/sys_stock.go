// 自动生成模板SysStock
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SysStock 结构体
// 如果含有time.Time 请自行import time包
type SysStock struct {
      global.GVA_MODEL
      StockNo  string `json:"stockNo" form:"stockNo" gorm:"column:stock_no;comment:编号;size:50;"`
      BookId  *int `json:"bookId" form:"bookId" gorm:"column:book_id;comment:图书id;size:11;"`
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:图书状态;size:11;"`
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:借阅者id;size:11;"`
      CreatorId  *int `json:"creatorId" form:"creatorId" gorm:"column:creator_id;comment:创建人id;size:11;"`
}


// TableName SysStock 表名
func (SysStock) TableName() string {
  return "sys_stock"
}

