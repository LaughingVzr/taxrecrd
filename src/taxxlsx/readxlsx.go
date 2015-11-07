package taxxlsx

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func Read() {
	// 文件路径
	file := "/Users/Laughing/Downloads/选案资料.xlsx"
	xlFile, err := xlsx.OpenFile(file)
	if err != nil {
		fmt.Println(err)
	}
	// 循环读取表格
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				fmt.Println(cell.Type())
			}
			fmt.Println()
		}
	}
}
