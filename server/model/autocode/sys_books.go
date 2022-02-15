// 自动生成模板SysBooks
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
      "time"
)
// SysBooks 结构体
// 如果含有time.Time 请自行import time包
type SysBooks struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:图书名称;size:100;"`
      Author  string `json:"author" form:"author" gorm:"column:author;comment:作者;size:50;"`
      Press  string `json:"press" form:"press" gorm:"column:press;comment:出版社;size:100;"`
      PubDate  *time.Time `json:"pubDate" form:"pubDate" gorm:"column:pub_date;comment:出版时间;"`
      Upc  string `json:"upc" form:"upc" gorm:"column:upc;comment:图书编码;size:50;"`
      BookcaseId  string `json:"bookcaseId" form:"bookcaseId" gorm:"column:bookcase_id;comment:书架编号;size:20;"`
      Price  *float64 `json:"price" form:"price" gorm:"column:price;comment:定价;size:10,2;"`
      TypeId  *int `json:"typeId" form:"typeId" gorm:"column:type_id;comment:类型id;size:11;"`
      Type   SysBookType   `json:"type" gorm:"foreignKey:TypeId;"`
      Stock  []SysStock `json:"stock" gorm:"foreignKey:BookId;"`
      Photo  string `json:"photo" form:"photo" gorm:"column:photo;comment:封面;size:200;"`
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:状态;size:11;"`
      Amount  *int `json:"amount" form:"amount" gorm:"column:amount;comment:数量;size:11;"`
}


// TableName SysBooks 表名
func (SysBooks) TableName() string {
  return "sys_books"
}

