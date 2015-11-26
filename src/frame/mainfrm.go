package frame

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	_ "github.com/lxn/win"
	_ "github.com/mattn/go-sqlite3"
	"log"
	_ "math/rand"
	model "model"
	_ "strings"
	xls "taxxlsx"
	_ "time"
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
	items []*model.TaxRecordRef
}

/*定向搜索条件*/
type StatFilter struct {
	OrgIndus     string  // 行业
	OrgTaxOffice string  // 所属税局
	OrgRegCap    int     // 经营规模
	StatTaxSum   float64 // 纳税数额(总)
	RowLimit     int     // 查询最大条数
}

/*初始化主窗口*/
func StartMF() {
	// 主窗口
	// var mainFrame *walk.MainWindow
	mw := new(TaxMainWindow)
	// 定向抽查参数
	param := new(StatFilter)
	// 主菜单
	// var mainMenu *walk.Menu
	if _, err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    TITLE,
		MinSize:  Size{800, 600}, // 最小大小
		Layout:   HBox{},
		Children: []Widget{
			PushButton{
				Text: "导入数据源",
				OnClicked: func() {
					mw.openFileAction()
				},
			},
			PushButton{
				Text: "定向抽查",
				OnClicked: func() {
					// mw.Close()
					if cmd, err := openDSearchDialog(mw, param); err != nil {
						log.Print(err)
					} else if cmd == walk.DlgCmdOK {
						// fmt.Printf("%+v", param)
						openResultTable(mw, param)
					}
				},
			},
		},
		// 这里直接Run,否则碰到了莫名奇妙的二次才能响应的问题
	}.Run()); err != nil {
		fmt.Println("error")
	}

	// mw.MainWindow.Run()
}

// new tax model
func NewTaxRecModel(taxs []*model.TaxRecordRef) *TaxRecModel {
	trm := new(TaxRecModel)
	trm.RestRows(taxs)
	// fmt.Println(trm)
	return trm
}

// 重置表内数据
func (trm *TaxRecModel) RestRows(taxs []*model.TaxRecordRef) {
	trm.items = taxs
	trm.PublishRowsReset()
}

func openResultTable(owner walk.Form, param *StatFilter) (int, error) {
	var dlg *walk.Dialog
	// var tw *walk.TableView
	// 查询数据
	model := NewTaxRecModel(queryData(param))

	return Dialog{
		AssignTo: &dlg,
		Title:    "定向抽查结果",
		MinSize:  Size{800, 600},
		Layout:   VBox{},
		Children: []Widget{
			TableView{
				// AssignTo: &tw,
				Columns: []TableViewColumn{
					{Title: "序号"},
					{Title: "机构码"},
					{Title: "机构名称"},
					{Title: "机构法人"},
					{Title: "机构注册时间", Format: "2006-01-02 15:04:05", Width: 100},
					{Title: "机构地址"},
					{Title: "机构注册资金"},
					{Title: "所属税务机关"},
					{Title: "所属行业"},
					{Title: "第一年应纳税收入", Format: "%.2f", Alignment: AlignFar},
					{Title: "增值税", Format: "%.2f", Alignment: AlignFar},
					{Title: "第二年应纳税收入", Format: "%.2f", Alignment: AlignFar},
					{Title: "增值税", Format: "%.2f", Alignment: AlignFar},
					{Title: "第三年应纳税收入", Format: "%.2f", Alignment: AlignFar},
					{Title: "增值税", Format: "%.2f", Alignment: AlignFar},
					{Title: "第一年应缴税所得额", Format: "%.2f", Alignment: AlignFar},
					{Title: "应缴纳所得税", Format: "%.2f", Alignment: AlignFar},
					{Title: "第二年应缴税所得额", Format: "%.2f", Alignment: AlignFar},
					{Title: "应缴纳所得税", Format: "%.2f", Alignment: AlignFar},
					{Title: "第三年应缴税所得额", Format: "%.2f", Alignment: AlignFar},
					{Title: "应缴纳所得税", Format: "%.2f", Alignment: AlignFar},
					{Title: "应纳税总额", Format: "%.2f", Alignment: AlignFar},
					{Title: "平均应纳税额", Format: "%.2f", Alignment: AlignFar},
					{Title: "检查时间", Format: "2006-01-02 15:04:05", Width: 150},
				},
				Model: model,
			},
		},
	}.Run(owner)
}

// open a table dialog
func openDSearchDialog(owner walk.Form, param *StatFilter) (int, error) {
	var dlg *walk.Dialog
	var db *walk.DataBinder
	var acceptPB, cancelPB *walk.PushButton

	return Dialog{
		AssignTo:      &dlg,
		Title:         "定向抽查",
		MinSize:       Size{300, 100}, // 最小大小
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		DataBinder: DataBinder{
			AssignTo:   &db,
			DataSource: param,
		},
		Layout: VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					Label{
						Text: "行业:",
					},
					LineEdit{
						Text: Bind("OrgIndus"),
					},
					Label{
						Text: "所属税局:",
					},
					LineEdit{
						Text: Bind("OrgTaxOffice"),
					},
					Label{
						Text: "经营规模(注资):",
					},
					NumberEdit{
						Value:  Bind("OrgRegCap", Range{1, 1000000000000}),
						Prefix: "￥",
					},
					Label{
						Text: "纳税数额(总):",
					},
					NumberEdit{
						Value:    Bind("StatTaxSum", Range{0.01, 1000000000000.00}),
						Prefix:   "￥",
						Decimals: 2,
					},
					Label{
						Text: "查询条数:",
					},
					NumberEdit{
						Value: Bind("RowLimit", Range{1, 1000000000000}),
					},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					PushButton{
						AssignTo: &acceptPB,
						Text:     "搜索",
						OnClicked: func() {
							if err := db.Submit(); err != nil {
								log.Print(err)
								return
							}
							// fmt.Printf("%+v", &param)
							dlg.Accept()
						},
					},
				},
			},
		},
	}.Run(owner)
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
	// invoke taxxlsx read and save 2 db
	fmt.Println(tmw.prevFilePath)
	if iSuc := xls.ReadAndSave(tmw.prevFilePath); iSuc {
		walk.MsgBox(tmw, "成功", "导入成功", walk.MsgBoxIconInformation)
		// 关闭对话框
		// tmw.closeDialog()
		// showTable()
	}
	return nil
}

// 根据条件查询数据
func queryData(param *StatFilter) []*model.TaxRecordRef {
	// 结果集
	var taxs []*model.TaxRecordRef
	orm.Debug = true

	o := orm.NewOrm()
	// 设置要查询的表
	qs := o.QueryTable("tax_record_ref")
	// 自定义条件
	cond := orm.NewCondition()
	// 判断参数
	if param.OrgIndus != "" {
		cond = cond.And("org_industry__contains", param.OrgIndus)
	}

	if param.OrgRegCap >= 0 {
		cond = cond.And("org_reg_cap__gte", param.OrgRegCap)
	}

	if param.OrgTaxOffice != "" {
		cond = cond.And("org_tax_office__contains", param.OrgTaxOffice)
	}

	if param.StatTaxSum >= 0 {
		cond = cond.And("stat_tax_sum__gte", param.StatTaxSum)
	}
	if param.RowLimit > 0 {
		qs = qs.Limit(param.RowLimit)
	}
	// cond.And("org_reg_cap__gte", param.OrgRegCap).And("stat_tax_sum__gte", param.StatTaxSum)

	qs = qs.SetCond(cond)
	// 查询
	qs.All(&taxs)
	// fmt.Printf("%+v", taxs)
	// 返回
	return taxs

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
		return item.OrgSerialNum
	case 2:
		return item.OrgName
	case 3:
		return item.OrgLegal
	case 4:
		return item.OrgRegT
	case 5:
		return item.OrgAddr
	case 6:
		return item.OrgRegCap
	case 7:
		return item.OrgTaxOffice
	case 8:
		return item.OrgIndus
	case 9:
		return item.TaxIncome1
	case 10:
		return item.TaxVat1
	case 11:
		return item.TaxIncome2
	case 12:
		return item.TaxVat2
	case 13:
		return item.TaxIncome3
	case 14:
		return item.TaxVat3
	case 15:
		return item.TaxIncomeVal1
	case 16:
		return item.TaxIncomeTax1
	case 17:
		return item.TaxIncomeVal2
	case 18:
		return item.TaxIncomeTax2
	case 19:
		return item.TaxIncomeVal3
	case 20:
		return item.TaxIncomeTax3
	case 21:
		return item.StatTaxSum
	case 22:
		return item.StatTaxAvg
	case 23:
		return item.StatCheckT
	}

	panic("unexpected col")
}
