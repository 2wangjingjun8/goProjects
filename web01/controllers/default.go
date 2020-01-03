package controllers

import (
	"github.com/astaxie/beego"
)

// MainController is a struct
type MainController struct {
	beego.Controller
}

// Get is a method
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

// GetLogin is a method
func (c *MainController) GetLogin() {
	c.Data["Website"] = "GetLogin.me"
	c.Data["Email"] = "GetLogin@gmail.com"
	c.TplName = "index.tpl"
}

// HandleLogin is a method
func (c *MainController) HandleLogin() {
	c.Data["Website"] = "HandleLogin.me"
	c.Data["Email"] = "HandleLogin@gmail.com"
	c.TplName = "index.tpl"
}
