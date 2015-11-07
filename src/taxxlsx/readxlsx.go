package taxxlsx

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tealeg/xlsx"
	"math"
	"model"
	_ "time"
)

/*对应字段的索引常量值*/
const (
	Id = iota
	OrgSerialNum
	OrgName
	OrgLegal
	OrgRegT
	OrgAddr
	OrgRegCap
	OrgTaxOffice
	OrgIndus
	TaxIncome1
	TaxVat1
	TaxIncome2
	TaxVat2
	TaxIncome3
	TaxVat3
	TaxIncomeVal1
	TaxIncomeTax1
	TaxIncomeVal2
	TaxIncomeTax2
	TaxIncomeVal3
	TaxIncomeTax3
	StatTaxSum
	StatTaxAvg
	StatCheckT
)

func init() {
	// 设置默认数据库
	orm.RegisterDataBase("default", "sqlite3", "../data/taxrec.db", 30)
	// 同步db(没有对应表时会建立相应的表)
	// orm.RunSyncdb("default", false, true)
}

func ReadAndSave() {
	// 文件路径
	// file := "/Users/Laughing/Downloads/选案资料.xlsx"
	file := "F:/workspace/go/taxrecrd/file/选案资料.xlsx"
	xlFile, err := xlsx.OpenFile(file)
	if err != nil {
		fmt.Println(err)
	}
	// 新建一条记录的对象
	taxrec := new(model.TaxRecordRef)
	// 循环读取表格
	for _, sheet := range xlFile.Sheets {
		for rnum, row := range sheet.Rows {
			// 忽略表头
			if rnum <= 1 {
				continue
			}
			// if rnum > 3 {

			// 	break
			// }
			for index, cell := range row.Cells {
				// fmt.Print(cell.Value)
				switch index {
				case Id:
					// 如果是id列写入税务记录
					taxrec.Id, _ = cell.Int()
				case OrgSerialNum:
					taxrec.OrgSerialNum = cell.Value
				case OrgName:
					taxrec.OrgName = cell.Value
				case OrgLegal:
					taxrec.OrgLegal = cell.Value
				case OrgRegT:
					excelTime, _ := cell.Float()
					taxrec.OrgRegT = conv2Unix(excelTime)
				case OrgAddr:
					taxrec.OrgAddr = cell.Value
				case OrgRegCap:
					taxrec.OrgRegCap, _ = cell.Int()
				case OrgTaxOffice:
					taxrec.OrgTaxOffice = cell.Value
				case OrgIndus:
					taxrec.OrgIndus = cell.Value
				case TaxIncome1:
					taxrec.TaxIncome1, _ = cell.Float()
				case TaxVat1:
					taxrec.TaxVat1, _ = cell.Float()
				case TaxIncome2:
					taxrec.TaxIncome2, _ = cell.Float()
				case TaxVat2:
					taxrec.TaxVat2, _ = cell.Float()
				case TaxIncome3:
					taxrec.TaxIncome3, _ = cell.Float()
				case TaxVat3:
					taxrec.TaxVat3, _ = cell.Float()
				case TaxIncomeVal1:
					taxrec.TaxIncomeVal1, _ = cell.Float()
				case TaxIncomeTax1:
					taxrec.TaxIncomeTax1, _ = cell.Float()
				case TaxIncomeVal2:
					taxrec.TaxIncomeVal2, _ = cell.Float()
				case TaxIncomeTax2:
					taxrec.TaxIncomeTax2, _ = cell.Float()
				case TaxIncomeVal3:
					taxrec.TaxIncomeVal3, _ = cell.Float()
				case TaxIncomeTax3:
					taxrec.TaxIncomeTax3, _ = cell.Float()
				case StatTaxSum:
					taxrec.StatTaxSum, _ = cell.Float()
				case StatTaxAvg:
					taxrec.StatTaxAvg, _ = cell.Float()
				case StatCheckT:
					// 拿到单元格时间
					excelTime, _ := cell.Float()
					// 默认设置为0
					taxrec.StatCheckT = conv2Unix(excelTime)
				default:
					fmt.Print(cell.Value)
				}
			}
			if taxrec.Id != -1 {
				// fmt.Print(taxrec)
				// 插入到数据库
				saveTaxRec(taxrec)
			}
		}
	}
}

/*写到数据库*/
func saveTaxRec(taxrcf *model.TaxRecordRef) {
	// 使用orm接口
	o := orm.NewOrm()
	// 写入
	fmt.Println(o.Insert(taxrcf))
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
