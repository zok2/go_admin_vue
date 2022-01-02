// 自动生成模板UserInfo
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// UserInfo 结构体
// 如果含有time.Time 请自行import time包
type UserInfo struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;size:1;"`
      Sex  *int `json:"sex" form:"sex" gorm:"column:sex;comment:;size:1;"`
      Age  *int `json:"age" form:"age" gorm:"column:age;comment:;size:1;"`
      Address  string `json:"address" form:"address" gorm:"column:address;comment:;size:4;"`
      Email  string `json:"email" form:"email" gorm:"column:email;comment:;size:1;"`
      Sn  *int `json:"sn" form:"sn" gorm:"column:sn;comment:;size:3;"`
}


// TableName UserInfo 表名
func (UserInfo) TableName() string {
  return "user_info"
}

