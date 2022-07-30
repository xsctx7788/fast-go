package main

import (
	_ "fast-go/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}
