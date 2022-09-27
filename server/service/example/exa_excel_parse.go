package example

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/xuri/excelize/v2"
	"strconv"
	"time"
)

type ExcelService struct{}

func (exa *ExcelService) ParseInfoList2Excel(infoList []system.SysBaseMenu, filePath string) error {
	excel := excelize.NewFile()
	excel.SetSheetRow("Sheet1", "A1", &[]string{"ID", "路由Name", "路由Path", "是否隐藏", "父节点", "排序", "文件名称"})
	for i, menu := range infoList {
		axis := fmt.Sprintf("A%d", i+2)
		excel.SetSheetRow("Sheet1", axis, &[]interface{}{
			menu.ID,
			menu.Name,
			menu.Path,
			menu.Hidden,
			menu.ParentId,
			menu.Sort,
			menu.Component,
		})
	}
	err := excel.SaveAs(filePath)
	return err
}

func (exa *ExcelService) ParseBookList2Excel(filePath string) ([]autocode.SysBooks,error) {
	skipHeader := true
	fixedHeader := []string{"图书名","作者","出版社","图书编码","书架编号","出版时间","定价","类型id","封面","状态","数量"}
	file, err := excelize.OpenFile(global.GVA_CONFIG.Excel.Dir +filePath)
	if err != nil {
		return nil, err
	}
	books := make([]autocode.SysBooks,0)
	rows, err := file.Rows("Sheet1")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		row, err := rows.Columns()
		if err != nil {
			return nil, err
		}
		if skipHeader {
			if exa.compareStrSlice(row, fixedHeader) {
				skipHeader = false
				continue
			} else {
				return nil, errors.New("Excel格式错误")
			}
		}
		if len(row) != len(fixedHeader) {
			continue
		}
		PubDate,_ := time.Parse("2006/1/2",row[5])
		Price,_ := strconv.ParseFloat(row[6],2)
		TypeId,_ := strconv.Atoi(row[7])
		Status,_ := strconv.Atoi(row[9])
		Amount,_ := strconv.Atoi(row[10])
		var (
			book = autocode.SysBooks{
				Name:       row[0],
				Author:     row[1],
				Press:      row[2],
				Upc:        row[3],
				BookcaseId: row[4],
				PubDate:    &PubDate,
				Price:      &Price,
				TypeId:     &TypeId,
				Photo:		"1",
				Status:     &Status,
				Amount:     &Amount,
			}
		)
		books = append(books, book)
	}
	return books, nil

}

func (exa *ExcelService) ParseExcel2InfoList() ([]system.SysBaseMenu, error) {
	skipHeader := true
	fixedHeader := []string{"ID", "路由Name", "路由Path", "是否隐藏", "父节点", "排序", "文件名称"}
	file, err := excelize.OpenFile(global.GVA_CONFIG.Excel.Dir + "ExcelImport.xlsx")
	if err != nil {
		return nil, err
	}
	menus := make([]system.SysBaseMenu, 0)
	rows, err := file.Rows("Sheet1")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		row, err := rows.Columns()
		if err != nil {
			return nil, err
		}
		if skipHeader {
			if exa.compareStrSlice(row, fixedHeader) {
				skipHeader = false
				continue
			} else {
				return nil, errors.New("Excel格式错误")
			}
		}
		if len(row) != len(fixedHeader) {
			continue
		}
		id, _ := strconv.Atoi(row[0])
		hidden, _ := strconv.ParseBool(row[3])
		sort, _ := strconv.Atoi(row[5])
		menu := system.SysBaseMenu{
			GVA_MODEL: global.GVA_MODEL{
				ID: uint(id),
			},
			Name:      row[1],
			Path:      row[2],
			Hidden:    hidden,
			ParentId:  row[4],
			Sort:      sort,
			Component: row[6],
		}
		menus = append(menus, menu)
	}
	return menus, nil
}

func (exa *ExcelService) compareStrSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	if (b == nil) != (a == nil) {
		return false
	}
	for key, value := range a {
		if value != b[key] {
			return false
		}
	}
	return true
}
