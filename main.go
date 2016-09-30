package main

import (
	"encoding/gob"
	"github.com/TalLannder/app1/models"
	_ "github.com/TalLannder/app1/routers"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
)

func main() {
	gob.Register(models.User{})
	beego.Run()
}
