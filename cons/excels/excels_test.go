/**
 * @Time: 2023/1/6 14:32
 * @Author: soupzhb@gmail.com
 * @File: excels_test.go
 * @Software: GoLand
 */

package excels

import (
	"testing"
	"time"
)

type ReqList struct {
	Name string
	Sex  string
	Age  int
}

func TestCreatExcel(t *testing.T) {
	//数据
	data := []ReqList{
		{"张三","男", 28},
		{"李四","男", 28},
	}
	dataSheet := GetSheet(data)
	// excel data 封装
	var excelData ReqExcel
	excelData.DefaultSheet = "收件人列表"
	excelData.ExcelName = time.Now().Format("2006-01-02_15:04:05") + ".xlsx"
	excelData.ExcelPath = "./" //os.Getenv("BASE_DIR") + "/tmp/"
	excelData.ExcelData = append(excelData.ExcelData, dataSheet)

	err := CreatExcel(excelData)
	t.Log(err)
}

// GetSheet
func GetSheet(data []ReqList) (dataSheet DataSheet) {
	dataSheet.SheetName = "测试组"
	dataSheet.SheetTitle = []string{
		"测试人姓名",
		"测试人性别",
		"测试人年龄",
	}

	for _, v := range data {
		var dd []interface{}

		dd = append(dd,
			v.Name,
			v.Sex,
			v.Age,
		)
		dataSheet.SheetData = append(dataSheet.SheetData, dd)
	}
	return
}
