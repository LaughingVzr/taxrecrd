package de

/*一张表搞定源文件字段*/
type TaxRecordRef struct {
	Id            int     `orm:"auto"` // 自增ID
	OrgSerialNum  string  // 纳税机构识别码
	OrgName       string  // 纳税机构名称
	OrgLegal      string  // 纳税机构法人
	OrgRegT       int64   `orm:"column(org_reg_time)"` // 纳税机构注册时间
	OrgAddr       string  `orm:"column(org_address)"`  // 机构地址
	OrgRegCap     int     // 机构注册资金
	OrgTaxOffice  string  // 机构主管税务机关
	OrgIndus      string  `orm:"column(org_industry)"` // 机构所属行业
	TaxIncome1    float64 // 应纳税收入
	TaxVat1       float64 // 增值税
	TaxIncome2    float64 // 应纳税收入
	TaxVat2       float64 // 增值税
	TaxIncome3    float64 // 应纳税收入
	TaxVat3       float64 // 增值税
	TaxIncomeVal1 float64 // 应缴税所得额
	TaxIncomeTax1 float64 // 应缴纳所得税
	TaxIncomeVal2 float64 // 应缴税所得额
	TaxIncomeTax2 float64 // 应缴纳所得税
	TaxIncomeVal3 float64 // 应缴税所得额
	TaxIncomeTax3 float64 // 应缴纳所得税
	StatTaxSum    float64 // 应纳税总额
	StatTaxAvg    float64 // 平均应纳税额
	StatCheckT    int64   `orm:"column(stat_check_time)"` // 检查时间

}

// open a search dialog
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
							} else {
								dlg.Accept()
							}
							// fmt.Printf("%+v", &param)
						},
					},
				},
			},
		},
	}.Run(owner)
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
					{Title: "机构注册时间", Format: "2006-01-02", Width: 100},
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
// 菜单点击事件
	var openFileAction *walk.Action
MenuItems: []MenuItem{
			Menu{
				Text: "&数据源",
				Items: []MenuItem{
					Action{
						AssignTo:    &openFileAction,
						Text:        "&导入数据",
						OnTriggered: mw.openFileAction,
					},
				},
			},
		},
