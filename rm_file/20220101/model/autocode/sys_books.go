// 自动生成模板SysBooks
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SysBooks 结构体
// 如果含有time.Time 请自行import time包
type SysBooks struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:图书名称;size:100;"`
      Author  string `json:"author" form:"author" gorm:"column:author;comment:作者;size:50;"`
      Isbn  string `json:"isbn" form:"isbn" gorm:"column:isbn;comment:国际标准图书编码;size:50;"`
      BookcaseId  *int `json:"bookcaseId" form:"bookcaseId" gorm:"column:bookcase_id;comment:书架编号;size:11;"`
      Price  *int `json:"price" form:"price" gorm:"column:price;comment:定价（分）;size:11;"`
      TypeId  *int `json:"typeId" form:"typeId" gorm:"column:type_id;comment:类型id;size:11;"`
      Photo  string `json:"photo" form:"photo" gorm:"column:photo;comment:封面;size:200;"`
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:状态;size:11;"`
}


// TableName SysBooks 表名
func (SysBooks) TableName() string {
  return "sys_books"
}

