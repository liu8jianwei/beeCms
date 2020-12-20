package controllers

import (
	"cmsBeego/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	var info models.Page
	c.SetSession("ename", 345)
	uer := c.GetSession("ename")
	logs.Informational("user login jack")
	fmt.Print(uer)
	models.UpdatePage()
	info = models.GetPage()
	c.Data["Website"] = info.Website
	c.Data["Email"] = info.Email
	c.TplName = "index.tpl"
}
