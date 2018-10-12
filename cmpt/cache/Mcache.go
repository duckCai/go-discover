package mcache

import (
	"fmt"
	"time"
	"github.com/astaxie/beego/cache"
)

var AppCache cache.Cache

func InitCache() {
	fmt.Println("init cache cmpt");
	var err error
	AppCache, err = cache.NewCache("memory", `{"interval":-1}`)
	AppCache.Put("AppList", "GO-Discover", 1000000*time.Hour)
	fmt.Println("value----", err, AppCache.Get("test"))
}

