// 自动生成模板SysBookType
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SysBookType 结构体
// 如果含有time.Time 请自行import time包
type SysBookType struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:分类名称;size:50;"`
      Pid  *int `json:"pid" form:"pid" gorm:"column:pid;comment:父id;size:11;"`
}


// TableName SysBookType 表名
func (SysBookType) TableName() string {
  return "sys_book_type"
}

