// 自动生成模板SysBookRentLog
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
      "github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// SysBookRentLog 结构体
// 如果含有time.Time 请自行import time包
type SysBookRentLog struct {
      global.GVA_MODEL
      BookId  *int `json:"bookId" form:"bookId" gorm:"column:book_id;comment:图书id;size:11;"`
      Stock  SysStock `json:"stock" gorm:"foreignKey:stockId;"`
      Book SysBooks  `json:"book" gorm:"foreignKey:BookId;"`
      Borrower system.SysUser  `json:"borrower" gorm:"foreignKey:userId;"`
      CreatorId  *int `json:"creatorId" form:"creatorId" gorm:"column:creator_id;comment:创建人id;size:11;"`
      Creator  system.SysUser  `json:"creator" gorm:"foreignKey:CreatorId;"`
      UserId  *int `json:"userId" form:"userId" gorm:"column:userId;comment:借阅者id;size:11;"`
      Type  *int `json:"type" form:"type" gorm:"column:type;comment:借还书类型;size:11;"`
      Remark  string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:200;"`
      StockId  *int `json:"stockId" form:"stockId" gorm:"column:stockId;comment:;size:11;"`
}


// TableName SysBookRentLog 表名
func (SysBookRentLog) TableName() string {
  return "sys_book_rent_log"
}

