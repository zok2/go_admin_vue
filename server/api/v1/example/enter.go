package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
)

type ApiGroup struct {
	autocode.SysBooksApi
	ExcelApi
	CustomerApi
	FileUploadAndDownloadApi
}

var (
	excelService                 = service.ServiceGroupApp.ExampleServiceGroup.ExcelService
	excelBookService             = service.ServiceGroupApp.AutoCodeServiceGroup.SysBooksService
	customerService              = service.ServiceGroupApp.ExampleServiceGroup.CustomerService
	fileUploadAndDownloadService = service.ServiceGroupApp.ExampleServiceGroup.FileUploadAndDownloadService
)
