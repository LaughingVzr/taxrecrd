package taxxlsx

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tealeg/xlsx"
	"math"
	"model"
	"time"
)

/*对应字段的索引常量值*/
const (
	Id = iota
	OrgName
	OrgSerialNum
	OrgIndus
	OrgBusScope
	OrgLegal
	OrgRegT
	OrgAddr
	OrgRegCap
	OrgTaxOffice
	OrgIsExport
	TaxIncome1
	TaxExIncome1
	TaxVat1
	TaxIncomeTax1
	TaxSum1
	TaxIncome2
	TaxExIncome2
	TaxVat2
	TaxIncomeTax2
	TaxSum2
	TaxIncome3
	TaxExIncome3
	TaxVat3
	TaxIncomeTax3
	TaxSum3
	StatTaxSum
	StatCheckT
	StatYear
	IsImportant
)

func init() {
	// 设置默认数据库
	orm.RegisterDataBase("default", "mysql", "root:root@/taxrecd?charset=utf8", 30)
	// 调试
	// orm.Debug = true
	// 同步db(没有对应表时会建立相应的表)
	// orm.RunSyncdb("default", false, true)
}

func ReadAndSave(filePath string) bool {
	// 文件路径
	// file := "/Users/Laughing/Downloads/选案资料.xlsx"
	// file := "F:/workspace/go/taxrecrd/file/选案资料.xlsx"
	xlFile, err := xlsx.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	// 切片数量
	sls := 0
	// 结果
	result := false
	// 循环读取表格
	for sheetnum, sheet := range xlFile.Sheets {
		if sheetnum >= 1 {
			// return true
			break
		}
		// 按行数计算切片循环数
		fmt.Println("Total Rows:", sheet.MaxRow)
		fmt.Println("Remainder:", (sheet.MaxRow)%100)
		sls = (sheet.MaxRow) / 100
		if (sheet.MaxRow)%100 > 0 {
			sls += 1
		}
		// 创建channel用于接收每个并发的完成情况
		c := make(chan bool)
		fmt.Println("Total Slice:", sls)
		// 循环拆分切片
		for i := 1; i <= sls; i++ {
			start := (i-1)*100 + 1
			end := start + 100
			if start < 0 {
				start = 0
			}
			// 如果超出边界则回到最大边界值
			if end > sheet.MaxRow {
				end = sheet.MaxRow
			}
			// fmt.Println("Start:", start, "  End:", end)
			tempSlice := sheet.Rows[start:end]
			// 去执行插入数据
			go readRows(tempSlice, c, int64(sls))
		}
		for {
			time.Sleep(1000)
			if <-c {
				result = true
				break
			}
		}
	}
	fmt.Println("Result:", result)
	return result
}

/*写到数据库*/
func saveTaxRec(taxrcf model.TaxRecordRef) int64 {
	// 使用orm接口
	o := orm.NewOrm()
	// 写入
	sucrow, err := o.Insert(&taxrcf)
	if err != nil {
		fmt.Println("数据库错误", err)
	}
	return sucrow
}

/*convert time to unix time (int64)*/
func conv2Unix(t float64) int64 {
	var unixTime int64 = 0
	// 判断是否为NaN的类型(是IEEE754标准中"not-a-numbner"的描述)
	if !math.IsNaN(t) {
		// 转化为golang中的时间UTC
		utcTime := xlsx.TimeFromExcelTime(t, false)
		// 从UTC时间转换到时间戳时间
		unixTime = utcTime.Unix()
	}
	return unixTime
}

// 并发执行解析插入数据库的工作(从上方的循环读取中抽出)
func readRows(rows []*xlsx.Row, c chan bool, max_row_count int64) bool {
	// 新建一条记录的对象
	var taxrec model.TaxRecordRef
	for _, row := range rows {
		for index, cell := range row.Cells {
			// fmt.Print(cell.Value)
			switch index {
			case Id:
				// 如果是id列写入税务记录
				taxrec.Id, _ = cell.Int()
			case OrgName:
				taxrec.OrgName = cell.Value
			case OrgSerialNum:
				taxrec.OrgSerialNum = cell.Value
			case OrgIndus:
				taxrec.OrgIndus = cell.Value
			case OrgBusScope:
				taxrec.OrgBusScope = cell.Value
			case OrgLegal:
				taxrec.OrgLegal = cell.Value
			case OrgRegT:
				excelTime, _ := cell.Float()
				taxrec.OrgRegT = conv2Unix(excelTime)
			case OrgAddr:
				taxrec.OrgAddr = cell.Value
			case OrgRegCap:
				temVal, _ := cell.Float()
				taxrec.OrgRegCap = Round(temVal, 3)
			case OrgTaxOffice:
				taxrec.OrgTaxOffice = cell.Value
			case OrgIsExport:
				taxrec.OrgIsExport = cell.Value
			case TaxIncome1:
				temVal, _ := cell.Float()
				taxrec.TaxIncome1 = Round(temVal, 3)
			case TaxExIncome1:
				temVal, _ := cell.Float()
				taxrec.TaxExIncome1 = Round(temVal, 3)
			case TaxVat1:
				temVal, _ := cell.Float()
				taxrec.TaxVat1 = Round(temVal, 3)
			case TaxIncomeTax1:
				temVal, _ := cell.Float()
				taxrec.TaxIncomeTax1 = Round(temVal, 3)
			case TaxSum1:
				temVal, _ := cell.Float()
				taxrec.TaxSum1 = Round(temVal, 3)
			case TaxIncome2:
				temVal, _ := cell.Float()
				taxrec.TaxIncome2 = Round(temVal, 3)
			case TaxExIncome2:
				temVal, _ := cell.Float()
				taxrec.TaxExIncome2 = Round(temVal, 3)
			case TaxVat2:
				temVal, _ := cell.Float()
				taxrec.TaxVat2 = Round(temVal, 3)
			case TaxIncomeTax2:
				temVal, _ := cell.Float()
				taxrec.TaxIncomeTax2 = Round(temVal, 3)
			case TaxSum2:
				temVal, _ := cell.Float()
				taxrec.TaxSum2 = Round(temVal, 3)
			case TaxIncome3:
				temVal, _ := cell.Float()
				taxrec.TaxIncome3 = Round(temVal, 3)
			case TaxExIncome3:
				temVal, _ := cell.Float()
				taxrec.TaxExIncome3 = Round(temVal, 3)
			case TaxVat3:
				temVal, _ := cell.Float()
				taxrec.TaxVat3 = Round(temVal, 3)
			case TaxIncomeTax3:
				temVal, _ := cell.Float()
				taxrec.TaxIncomeTax3 = Round(temVal, 3)
			case TaxSum3:
				temVal, _ := cell.Float()
				taxrec.TaxSum3 = Round(temVal, 3)
			case StatTaxSum:
				temVal, _ := cell.Float()
				taxrec.StatTaxSum = Round(temVal, 3)
			case StatCheckT:
				// 拿到单元格时间
				excelTime, _ := cell.Float()
				// 默认设置为0
				taxrec.StatCheckT = conv2Unix(excelTime)
			case StatYear:
				taxrec.StatYear = cell.Value
			case IsImportant:
				taxrec.IsImportant = cell.Value
			default:
				fmt.Println(cell.Value)
			}
		}
		// fmt.Println("Object:", taxrec)
		// 插入数据库
		if sucrow := saveTaxRec(taxrec); sucrow == max_row_count {
			c <- true
		}
	}
	return false
}

// 截断float64后的小数点,可自定义位数
func Round(f float64, n int) float64 {
	// 处理非数字的问题
	if math.IsNaN(f) {
		return 0
	}
	pow10_n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
}
