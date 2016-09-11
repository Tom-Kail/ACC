package main

import (
	"runtime"

	_ "github.com/Tom-Kail/acc/go/beego_test/routers"
	"github.com/astaxie/beego"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	beego.Run()
}
