package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type SysStockSearch struct {
	autocode.SysStock
	request.PageInfo
}

type SysStockBorrowedStruct struct {
	StockId int   `json:"stockId"`
	Status  int    `json:"status"`
	BookId  int   `json:"bookId"`
	UserId  int   `json:"userId"`
	Type    int    `json:"type"`
	Remark  string `json:"remark"`
}
