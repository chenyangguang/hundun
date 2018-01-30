package controllers

import (
	"github.com/astaxie/beego"
)

// UserController user control
type UserController struct {
	beego.Controller
}

// Get get user index
func (c *UserController) Get() {
	c.TplName = "user.tpl"
}

// Post user login
func (c *UserController) Post() {

}
