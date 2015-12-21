package model

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

// 下拉框基础Struct
type Combase struct {
	Name string `orm:"column(org_industry)"` // 显示名称
	Val  string `orm:"column(org_industry)"` // 实际值
}

// 测试行业下拉框
func ShowTestIndus() []*Combase {
	// combo 结果集
	var combos []*Combase
	o := orm.NewOrm()
	num, err := o.Raw("SELECT org_industry FROM tax_record_ref WHERE org_industry!='' GROUP BY org_industry").QueryRows(&combos)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Numbers:", num)
	}
	return combos
}
