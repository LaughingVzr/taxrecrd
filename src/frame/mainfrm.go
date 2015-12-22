package frame

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
	modelp "model"
	xls "taxxlsx"
	"time"
)

const (
	TITLE = "稽查核查筛选工具" // 标题
)

/*主窗口Model*/
type TaxMainWindow struct {
	*walk.MainWindow
	tabWidget    *walk.TabWidget
	prevFilePath string
}

/*税务记录列表model*/
type TaxRecModel struct {
	walk.TableModelBase
	items []*modelp.TaxRecordRef
}

/*定向搜索条件*/
type StatFilter struct {
	OrgIndus        string  // 行业
	OrgBusScope     string  // 经营范围
	StatTaxSumStart float64 // 税收下限
	StatTaxSumEnd   float64 // 税收上限
	OrgIsExport     bool    // 是否是出口企业
	IsImportant     bool    // 是否是重点税源企业
}

// 列表对象
var model *TaxRecModel

// 增值纳税表
var vatTable string = "tax_record_ref"

// 行业数据
var combs []*modelp.Combase

// 行业列表
var indusComb *walk.ComboBox

/*初始化主窗口*/
func StartMF() {

	// var mainFrame *walk.MainWindow
	mw := new(TaxMainWindow)
	// 空集合
	var taxs []*modelp.TaxRecordRef
	// 如果库里有数据会优先展示库里的数据
	taxs = defaDisplayData()
	model = NewTaxRecModel(taxs)

	// 筛选参数绑定对象
	var binder *walk.DataBinder
	param := new(StatFilter)
	// 主菜单
	// var mainMenu *walk.Menu
	if _, err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    TITLE,
		MinSize:  Size{800, 600}, // 最小大小
		Layout:   VBox{},
		DataBinder: DataBinder{
			AssignTo:   &binder,
			DataSource: param,
		},
		Children: []Widget{
			// 表格部分
			TableView{
				Columns: []TableViewColumn{
					{Title: "序号"},
					{Title: "纳税人名称"},
					{Title: "纳税人识别号"},
					{Title: "行业"},
					{Title: "经营范围"},
					{Title: "法定代表人"},
					{Title: "开业日期", Format: "2006-01-02"},
					{Title: "生产经营地址"},
					{Title: "注册资本", Format: "%.2f", Alignment: AlignFar},
					{Title: "主管税务机关"},
					{Title: "是否出口企业"},
					{Title: "第一年应税收入", Format: "%.2f", Alignment: AlignFar},
					{Title: "第一年出口收入", Format: "%.2f", Alignment: AlignFar},
					{Title: "第一年增值税", Format: "%.2f", Alignment: AlignFar},
					{Title: "第一年所得税", Format: "%.2f", Alignment: AlignFar},
					{Title: "第一年税额合计", Format: "%.2f", Alignment: AlignFar},
					{Title: "第二年应税收入", Format: "%.2f", Alignment: AlignFar},
					{Title: "第二年出口收入", Format: "%.2f", Alignment: AlignFar},
					{Title: "第二年增值税", Format: "%.2f", Alignment: AlignFar},
					{Title: "第二年所得税", Format: "%.2f", Alignment: AlignFar},
					{Title: "第二年税额合计", Format: "%.2f", Alignment: AlignFar},
					{Title: "第三年应税收入", Format: "%.2f", Alignment: AlignFar},
					{Title: "第三年出口收入", Format: "%.2f", Alignment: AlignFar},
					{Title: "第三年增值税", Format: "%.2f", Alignment: AlignFar},
					{Title: "第三年所得税", Format: "%.2f", Alignment: AlignFar},
					{Title: "第三年税额合计", Format: "%.2f", Alignment: AlignFar},
					{Title: "三年缴纳税款总计", Format: "%.2f", Alignment: AlignFar},
					{Title: "稽查时间", Format: "2006-01-02"},
					{Title: "稽查年度"},
					{Title: "是否重点税源"},
				},
				Model: model,
			},
			GroupBox{
				Title:  "筛选数据",
				Layout: VBox{},
				Children: []Widget{
					Composite{
						Layout: HBox{},
						Children: []Widget{
							Label{
								Text: "行业：",
							},
							ComboBox{
								AssignTo:      &indusComb,
								Value:         Bind("OrgIndus"),
								BindingMember: "Val",
								DisplayMember: "Name",
								Model:         modelp.ShowTestIndus(),
							},
							Label{
								Text: "经营范围：",
							},
							LineEdit{
								Text: Bind("OrgBusScope"),
							},
						},
					},
					Composite{
						Layout: HBox{},
						Children: []Widget{
							Label{
								Text: "税收收入大于：",
							},
							NumberEdit{
								Value:    Bind("StatTaxSumStart", Range{0.00, 99999999999999999.99}),
								Prefix:   "￥ ",
								Decimals: 2,
							},
							Label{
								Text: "税收收入小于：",
							},
							NumberEdit{
								Value:    Bind("StatTaxSumEnd", Range{0.00, 99999999999999999.99}),
								Prefix:   "￥ ",
								Decimals: 2,
							},
						},
					},
					Composite{
						Layout: HBox{},
						Children: []Widget{
							CheckBox{
								Checked: Bind("IsImportant"),
							},
							Label{
								Text: "是否是重点税源企业",
							},
							CheckBox{
								Checked: Bind("OrgIsExport"),
							},
							Label{
								Text: "是否是出口企业",
							},
							PushButton{
								Text: "抽查数据",
								OnClicked: func() {
									// 提交数据到绑定器
									if err := binder.Submit(); err != nil {
										log.Print(err)
										return
									}
									// 执行查询并更新列表
									fmt.Println("Params:", param)
									model.RestRows(queryData(param))
								},
							},
							PushButton{
								Text: "导入数据",
								OnClicked: func() {
									mw.openFileAction()
								},
							},
							PushButton{
								Text: "刷新数据",
								OnClicked: func() {
									// 刷新行业列表
									getCombs()
									// 清空参数
									err := binder.Reset()
									if err != nil {
										fmt.Println(err)
									}
									// 刷新数据
									model.RestRows(defaDisplayData())
								},
							},
							PushButton{
								Text: "清空数据",
								OnClicked: func() {
									// 清空数据然后查询展示空列表
									truncateData(vatTable)
									model.RestRows(defaDisplayData())
									walk.MsgBox(mw, "成功", "清空成功", walk.MsgBoxIconInformation)
								},
							},
						},
					},
				},
			},
		},
		// 这里直接Run,否则碰到了莫名奇妙的二次才能响应的问题
	}.Run()); err != nil {
		fmt.Println(err)
	}
	// mw.MainWindow.Run()
}

// new tax model
func NewTaxRecModel(taxs []*modelp.TaxRecordRef) *TaxRecModel {
	trm := new(TaxRecModel)
	trm.RestRows(taxs)
	// fmt.Println(trm)
	return trm
}

// 重置表内数据
func (trm *TaxRecModel) RestRows(taxs []*modelp.TaxRecordRef) {
	trm.items = taxs
	trm.PublishRowsReset()
}

// 根据条件返回查询结果
func getSearchResult(param *StatFilter) []*modelp.TaxRecordRef {
	return queryData(param)
}

// 刷新行业数据
func getCombs() {
	indusComb.SetModel(modelp.ShowTestIndus())
}

// open file action
func (tmw *TaxMainWindow) openFileAction() {
	if err := tmw.openFileDia(); err != nil {
		log.Fatal(err)
	}
}

// real open file dialog
func (tmw *TaxMainWindow) openFileDia() error {
	dlg := new(walk.FileDialog)
	dlg.FilePath = tmw.prevFilePath
	dlg.Filter = "Excel Files(*.xls;*.xlsx)|*.xls;*.xlsx"
	dlg.Title = "请选择Excel文件"

	if ok, err := dlg.ShowOpen(tmw); err != nil {
		return err
	} else if !ok {
		return nil
	}

	// 拿到选择的文件的路径放入tmw struct
	tmw.prevFilePath = dlg.FilePath

	fmt.Println(tmw.prevFilePath)
	// pre clear data
	truncateData(vatTable)
	// invoke taxxlsx read and save 2 db
	if iSuc := xls.ReadAndSave(tmw.prevFilePath); iSuc {
		walk.MsgBox(tmw, "成功", "导入成功", walk.MsgBoxIconInformation)
		// 显示默认列表数据
		model.RestRows(defaDisplayData())
	}
	return nil
}

// 根据条件查询数据
func queryData(param *StatFilter) []*modelp.TaxRecordRef {
	// 结果集
	var taxs []*modelp.TaxRecordRef
	// orm.Debug = true

	o := orm.NewOrm()
	// 设置要查询的表
	qs := o.QueryTable("tax_record_ref")
	// 自定义条件
	cond := orm.NewCondition()
	// 判断参数
	if param.OrgIndus != "" {
		cond = cond.And("org_industry__contains", param.OrgIndus)
	}

	if param.OrgBusScope != "" {
		cond = cond.And("org_bus_scope__contains", param.OrgBusScope)
	}

	if param.StatTaxSumStart >= 0 {
		cond = cond.And("stat_tax_sum__gte", param.StatTaxSumStart)
	}

	if param.StatTaxSumEnd > 0 {
		cond = cond.And("stat_tax_sum__lte", param.StatTaxSumEnd)
	}

	if param.OrgIsExport {
		cond = cond.And("org_is_export__contains", "是")
	}

	if param.IsImportant {
		cond = cond.And("is_important__contains", "Y")
	}

	// 设置有效记录
	cond = cond.And("status", 0)

	// if param.RowLimit > 0 {
	// 	qs = qs.Limit(param.RowLimit)
	// }

	qs = qs.SetCond(cond)
	// 查询
	qs.All(&taxs)
	// fmt.Printf("%+v", taxs)
	go setRowsInvalid(taxs)
	// 返回
	return taxs
}

// 默认展示全部数据查询
func defaDisplayData() []*modelp.TaxRecordRef {
	// 稍等数据大量插入加载
	time.Sleep(1500)
	// 结果集
	var taxs []*modelp.TaxRecordRef
	o := orm.NewOrm()
	num, err := o.Raw("SELECT * FROM tax_record_ref WHERE org_name!='' AND status=0").QueryRows(&taxs)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Numbers:", num)
	}
	return taxs
}

// 清空某表,重置索引
func truncateData(table string) {
	o := orm.NewOrm()
	_, err1 := o.Raw("TRUNCATE TABLE " + table).Exec()
	if err1 != nil {
		fmt.Println(err1)
	}
	// 重置索引
	_, err2 := o.Raw("ALTER TABLE " + table + " AUTO_INCREMENT=1").Exec()
	if err2 != nil {
		fmt.Println(err2)
	}
	// 等候一会儿清理数据
	time.Sleep(2000)
}

// 设置失效数据
func setRowsInvalid(rows []*modelp.TaxRecordRef) {
	o := orm.NewOrm()
	for _, tax := range rows {
		// fmt.Println(tax)
		_, err := o.Raw("UPDATE tax_record_ref SET status=1 WHERE id=?", tax.Id).Exec()
		if err != nil {
			fmt.Println(err)
		}
	}
}

/*--------------Important----------*/
// Called by the TableView from SetModel and every time the model publishes a
// RowsReset event.

func (m *TaxRecModel) RowCount() int {
	return len(m.items)
}

// Called by the TableView when it needs the text to display for a given cell.
func (m *TaxRecModel) Value(row, col int) interface{} {
	item := m.items[row]

	switch col {
	case 0:
		return item.Id
	case 1:
		return item.OrgName
	case 2:
		return item.OrgSerialNum
	case 3:
		return item.OrgIndus
	case 4:
		return item.OrgBusScope
	case 5:
		return item.OrgLegal
	case 6:
		regT := item.OrgRegT
		if regT == 0 {
			return ""
		}
		return time.Unix(regT, 0)
	case 7:
		return item.OrgAddr
	case 8:
		return item.OrgRegCap
	case 9:
		return item.OrgTaxOffice
	case 10:
		return item.OrgIsExport
	case 11:
		return item.TaxIncome1
	case 12:
		return item.TaxExIncome1
	case 13:
		return item.TaxVat1
	case 14:
		return item.TaxIncomeTax1
	case 15:
		return item.TaxSum1
	case 16:
		return item.TaxIncome2
	case 17:
		return item.TaxExIncome2
	case 18:
		return item.TaxVat2
	case 19:
		return item.TaxIncomeTax2
	case 20:
		return item.TaxSum2
	case 21:
		return item.TaxIncome3
	case 22:
		return item.TaxExIncome3
	case 23:
		return item.TaxVat3
	case 24:
		return item.TaxIncomeTax3
	case 25:
		return item.TaxSum3
	case 26:
		return item.StatTaxSum
	case 27:
		chkT := item.StatCheckT
		if chkT == 0 {
			return ""
		}
		return time.Unix(chkT, 0)
	case 28:
		return item.StatYear
	case 29:
		return item.IsImportant
	}

	panic("unexpected col")
}
