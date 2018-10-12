package controllers

import (
	"github.com/astaxie/beego"
	"go-discover/cmpt/cache"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["AppName"] = mcache.AppCache.Get("AppName")
	c.Data["RegApps"] = `{"interval":-1}`
	c.TplName = "index.tpl"
}

func (this *MainController) RegisterApp() {
	appName := this.GetString("appName")
	if appName != "" {
		var appList [50] string
		mcache.AppCache.Put("AppName", appName, 90000)
		this.Ctx.WriteString(appName + " register success!")
		return 
	} else {
		this.Ctx.WriteString("register failure! appName is empty")
		return
	}
}

func (this *MainController) ListApp() {
 	appList := mcache.AppCache.Get("AppList")
	this.Ctx.WriteString(appName + " register success!")
	return 
}