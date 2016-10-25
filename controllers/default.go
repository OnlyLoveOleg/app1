package controllers

import (
  "fmt"
  "encoding/json"
	"github.com/TalLannder/app1/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
  user := new(models.User)
	user.Name = "Sponge Bob"
	user.Age = 100

	sess := c.GetSession("session")

	if sess == nil {
    b, err := json.Marshal(user)
    if err != nil {
        fmt.Println("error:", err)
    }
    newSession := make(map[string]interface{})
    newSession["user"] = b
    fmt.Printf("%#v\n", newSession)
		c.SetSession("session", newSession)
	} else {
	  c.Data["Session"] = sess
  }
	c.TplName = "index.tpl"
}


type Cat struct {
    Name string
    Age  int
}

func (c *MainController) GetCat() {
  fmt.Println("In GetUserSession - start")

	sess := c.GetSession("session")

  cat := new(Cat)

  if v, ok := sess.(map[string]interface{})["user"].([]byte); ok {
    err := json.Unmarshal(v, &cat)
    if err != nil {
        fmt.Println("error:", err)
    }
  }

  fmt.Printf("%#v\n", cat)
	c.Data["Cat"] = cat
	c.TplName = "GetCat.tpl"
}
