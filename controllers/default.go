package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/cache"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	bm, err := cache.NewCache("memory", `{"interval":-1}`)
	fmt.Println("value----", err, bm.Get("test"))
	c.Data["AppName"] = "Go-discover"
	c.TplName = "index.tpl"
}
