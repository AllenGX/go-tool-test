package controllers

import (
	"github.com/astaxie/beego"
)

type DainController struct {
	beego.Controller
}

func (c *DainController) Get() {
	par := make(map[string]interface{})
	par["key"] = "value"
	c.Data["Email"] = "bbb"
	c.Data["json"] = &par
	c.ServeJSON()

}

func (c *DainController) Post() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
