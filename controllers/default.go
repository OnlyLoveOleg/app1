package controllers

import (
	"github.com/TalLannder/app1/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	u := models.User{}
	u.Name = "Sponge Bob"
	u.Age = 100

	sess := c.GetSession("session")

	if sess == nil {
		c.SetSession("session", u)
	}

	c.Data["Session"] = sess
	c.TplName = "index.tpl"
}
