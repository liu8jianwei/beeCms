package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(182.92.218.34:3306)/blog?charset=utf8")
	orm.RegisterModel(new(Page))
}

type Page struct {
	Id      int
	Website string
	Email   string
}

func GetPage() Page {
	// rln := Page{Website: "baidu", Email: "163.com"}
	// return rln
	orm := orm.NewOrm()
	p := Page{Id: 1}
	err := orm.Read(&p)
	if err != nil {
		fmt.Println(err)
	}
	return p
}

func UpdatePage() {
	o := orm.NewOrm()
	p := Page{Id: 1, Website: "http://tencent.com"}
	num, err := o.Update(&p, "Website")
	fmt.Println(num)

	if err != nil {
		fmt.Println(err)
	}
}
