package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

type ExcelInfo struct {
	FileName string               `json:"fileName"` // 文件名
	InfoList []system.SysBaseMenu `json:"infoList"`
}
type ExcelBook struct {
	FileName string               `json:"fileName"` // 文件名
	BookInfoList []autocode.SysBooks `json:"bookInfoList"`
}