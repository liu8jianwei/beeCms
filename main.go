package main

import (
	_ "cmsBeego/routers"
	_ "cmsBeego/sysinit"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	logs.SetLevel(beego.LevelInformational)
	logs.SetLogger("file", `{"filename":"/tmp/go.log"}`)

	beego.Run()
}
