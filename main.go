package main

import (
	_ "go-discover/routers"
	"github.com/astaxie/beego"
	"go-discover/cmpt/cache"
)


func init() {
	mcache.InitCache()
}

func main() {
	beego.Run()
}

