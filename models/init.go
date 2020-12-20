package models

import "github.com/astaxie/beego/orm"

func init() {
	// orm.RegisterDataBase("default", "mysql", "root:root@tcp(182.92.218.34:3306)/blog?charset=utf8")
	orm.RegisterModel(new(MenuModel))
}
