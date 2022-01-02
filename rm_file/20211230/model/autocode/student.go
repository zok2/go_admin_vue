// 自动生成模板Student
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Student 结构体
// 如果含有time.Time 请自行import time包
type Student struct {
      global.GVA_MODEL
      Sn  string `json:"sn" form:"sn" gorm:"column:sn;comment:;size:20;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;size:20;"`
      Sex  *int `json:"sex" form:"sex" gorm:"column:sex;comment:;size:4;"`
      Email  string `json:"email" form:"email" gorm:"column:email;comment:;size:50;"`
}


// TableName Student 表名
func (Student) TableName() string {
  return "student"
}

