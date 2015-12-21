package main

import (
	"frame"
)

func init() {
	// 设置默认数据库
	// orm.RegisterDataBase("default", "sqlite3", "../data/taxrec.db", 30)
	// 同步db(没有对应表时会建立相应的表)
	// orm.RunSyncdb("default", false, true)
}

func main() {
	// 使用orm接口
	// o := orm.NewOrm()

	// taxrec := new(model.TaxRecordRef)
	// taxrec.OrgName = "华龙方便面"

	// fmt.Println(o.Insert(taxrec))
	// xls.ReadAndSave()
	frame.StartMF()
}
