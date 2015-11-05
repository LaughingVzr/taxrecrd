package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id       int
	UserName string
	Passwd   string
}

func init() {
	// 设置默认数据库
	orm.RegisterDataBase("default", "sqlite3", "../data/taxrec.db", 30)
	// 注册Model模型
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()
	user := User{Id: 2, UserName: "Lwz", Passwd: "sdlfkjdf"}
	id, err := o.Insert(&user)
	fmt.Println("ID:%d,ERR:%v\n", id, err)
}
