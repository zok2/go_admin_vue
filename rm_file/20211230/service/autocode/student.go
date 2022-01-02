package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type StudentService struct {
}

// CreateStudent 创建Student记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentService *StudentService) CreateStudent(student autocode.Student) (err error) {
	err = global.GVA_DB.Create(&student).Error
	return err
}

// DeleteStudent 删除Student记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentService *StudentService)DeleteStudent(student autocode.Student) (err error) {
	err = global.GVA_DB.Delete(&student).Error
	return err
}

// DeleteStudentByIds 批量删除Student记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentService *StudentService)DeleteStudentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Student{},"id in ?",ids.Ids).Error
	return err
}

// UpdateStudent 更新Student记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentService *StudentService)UpdateStudent(student autocode.Student) (err error) {
	err = global.GVA_DB.Save(&student).Error
	return err
}

// GetStudent 根据id获取Student记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentService *StudentService)GetStudent(id uint) (err error, student autocode.Student) {
	err = global.GVA_DB.Where("id = ?", id).First(&student).Error
	return
}

// GetStudentInfoList 分页获取Student记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentService *StudentService)GetStudentInfoList(info autoCodeReq.StudentSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.Student{})
    var students []autocode.Student
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Sn != "" {
        db = db.Where("sn = ?",info.Sn)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.Sex != nil {
        db = db.Where("sex = ?",info.Sex)
    }
    if info.Email != "" {
        db = db.Where("email = ?",info.Email)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&students).Error
	return err, students, total
}
