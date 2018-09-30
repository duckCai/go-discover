package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["AppName"] = "Go-discover"
	c.TplName = "index.tpl"
}
