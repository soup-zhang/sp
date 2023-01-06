package excels

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
)

type ReqExcel struct {
	ExcelPath    string      //路径
	ExcelName    string      //文件名 book.xlsx
	ExcelData    []DataSheet //sheet 数据源
	DefaultSheet string      //默认选中的工作表
}

type DataSheet struct {
	SheetTitle []string        //表头 列标题
	SheetData  [][]interface{} //数据源
	SheetName  string          //SheetName 工作表名称
}

func CreatExcel(reqExcel ReqExcel) (err error) {
	f := excelize.NewFile()

	for _, v := range reqExcel.ExcelData {
		makeSheet(v, f, reqExcel.DefaultSheet)
	}

	f.DeleteSheet("Sheet1") //删除默认工作表名称

	fileName := reqExcel.ExcelPath + reqExcel.ExcelName
	if err := f.SaveAs(fileName); err != nil {
		return err
	}
	return err
}

func makeSheet(reqSheet DataSheet, f *excelize.File, defaultSheet string) {
	// Create a new sheet.
	index := f.NewSheet(reqSheet.SheetName)

	//表头
	for columnNum, v := range reqSheet.SheetTitle {
		sheetPosition := NumToLetter(columnNum+1) + "1" //确认表格位置A1,A2..
		f.SetCellValue(reqSheet.SheetName, sheetPosition, v)
	}

	//数据
	for lineNum, v := range reqSheet.SheetData {
		// Set value of a cell.
		columnNum := 0
		for _, vv := range v {
			columnNum++
			sheetPosition := NumToLetter(columnNum) + strconv.Itoa(lineNum+2) //确认表格位置B1,B2..
			//包里已对vv变量类型做了处理
			f.SetCellValue(reqSheet.SheetName, sheetPosition, vv)
		}
	}

	if reqSheet.SheetName == defaultSheet {
		// Set active sheet of the workbook.
		f.SetActiveSheet(index)
	}

}

//数字转字母
func NumToLetter(Num int) string {
	var (
		Str  string = ""
		k    int
		temp []int //保存转化后每一位数据的值，然后通过索引的方式匹配A-Z
	)
	//用来匹配的字符A-Z
	Slice := []string{"", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O",
		"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	if Num > 26 { //数据大于26需要进行拆分
		for {
			k = Num % 26 //从个位开始拆分，如果求余为0，说明末尾为26，也就是Z，如果是转化为26进制数，则末尾是可以为0的，这里必须为A-Z中的一个
			if k == 0 {
				temp = append(temp, 26)
				k = 26
			} else {
				temp = append(temp, k)
			}
			Num = (Num - k) / 26 //减去Num最后一位数的值，因为已经记录在temp中
			if Num <= 26 {       //小于等于26直接进行匹配，不需要进行数据拆分
				temp = append(temp, Num)
				break
			}
		}
	} else {
		return Slice[Num]
	}
	for _, value := range temp {
		Str = Slice[value] + Str //因为数据切分后存储顺序是反的，所以Str要放在后面
	}
	return Str
}

