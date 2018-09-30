package main

import (
	_ "go-discover/routers"
	"github.com/astaxie/beego"
	"fmt"
	"time"
	"github.com/astaxie/beego/cache"
)

var cacheUtils *Cache

func init() {
	fmt.Println("init cache cmpt");
	bm, err := cache.NewCache("memory", `{"interval":-1}`)
	bm.Put("test", 1, 1000000*time.Hour)
	fmt.Println("value----", err, bm.Get("test"))
}

func main() {
	beego.Run()
}

