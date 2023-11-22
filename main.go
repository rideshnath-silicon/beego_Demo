package main

import (
	_ "CrudDemo/routers"

	"github.com/astaxie/beego"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}
